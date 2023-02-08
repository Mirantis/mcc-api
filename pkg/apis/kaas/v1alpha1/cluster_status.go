package v1alpha1

import (
	"context"
	clusterv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/v1alpha1"
	lcmv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/lcm/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	deploymentutil "k8s.io/kubectl/pkg/util/deployment"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +gocode:public-api=true
var excludedReadinessCheckNamespaces = map[string]bool{"default": true}

// ClusterPollStatus describes cluster conditions that are updated as a result of polling
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +gocode:public-api=true
type ClusterPollStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	LoadBalancerStatus LoadBalancerStatus `json:"loadBalancerStatus,omitempty"`
	ConditionsSummary  `json:",inline"`
}

// mandatoryReadinessCheckNamespaces list of namespaces resources to check for notReadyKaasObjects
// +gocode:public-api=true
var mandatoryReadinessCheckNamespaces = map[string]bool{"kube-system": true, "kaas": true}

// ClusterPollStatusList contains a list of ClusterPollStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ClusterPollStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ClusterPollStatus `json:"items"`
}

// GetNotReadyObjects returns not readyobjects related to MCC
// +gocode:public-api=true
func GetNotReadyObjects(ctx context.Context, cl client.Client, nsReadinessList []string) (Objects, error) {
	var notReadyKaaSObjects Objects

	for _, ns := range nsReadinessList {
		opts := &client.ListOptions{
			Namespace: ns,
		}
		services := &corev1.ServiceList{}
		err := cl.List(ctx, services, opts)
		if err != nil {
			return Objects{}, err
		}
		for _, service := range services.Items {
			if service.Spec.Type != corev1.ServiceTypeLoadBalancer {
				continue
			}
			if service.Status.LoadBalancer.Ingress == nil || (service.Status.LoadBalancer.Ingress[0].IP == "" && service.Status.LoadBalancer.Ingress[0].Hostname == "") {
				notReadyKaaSObjects.Services = append(notReadyKaaSObjects.Services, Service{Name: service.Name, Namespace: service.Namespace})
			}
		}

		deployments := &appsv1.DeploymentList{}
		err = cl.List(ctx, deployments, opts)
		if err != nil {
			return Objects{}, err
		}
		for _, deployment := range deployments.Items {
			if !IsDeploymentReady(&deployment) {
				notReadyKaaSObjects.Deployments = append(notReadyKaaSObjects.Deployments, Controller{
					Name:          deployment.Name,
					Namespace:     deployment.Namespace,
					Replicas:      deployment.Status.Replicas,
					ReadyReplicas: deployment.Status.ReadyReplicas,
				})
			}
		}

		statefulsets := &appsv1.StatefulSetList{}
		err = cl.List(ctx, statefulsets, opts)
		if err != nil {
			return Objects{}, err
		}
		for _, statefulset := range statefulsets.Items {
			ready, ctrl := isStatefulSetReady(&statefulset)
			if !ready {
				notReadyKaaSObjects.StatefulSets = append(notReadyKaaSObjects.StatefulSets, ctrl)
			}
		}

		daemonsets := &appsv1.DaemonSetList{}
		err = cl.List(ctx, daemonsets, opts)
		if err != nil {
			return Objects{}, err
		}
		for _, daemonset := range daemonsets.Items {
			ready, ds := isDaemonSetReady(&daemonset)
			if !ready {
				notReadyKaaSObjects.DaemonSets = append(notReadyKaaSObjects.DaemonSets, ds)
			}
		}
	}
	return notReadyKaaSObjects, nil
}

// based on https://github.com/kubernetes/kubernetes/blob/f610eee1/staging/src/k8s.io/kubectl/pkg/polymorphichelpers/rollout_status.go#L59
// +gocode:public-api=true
func IsDeploymentReady(d *appsv1.Deployment) bool {
	if d.Generation > d.Status.ObservedGeneration {
		return false
	}
	cond := deploymentutil.GetDeploymentCondition(d.Status, appsv1.DeploymentProgressing)
	if cond != nil && cond.Reason == deploymentutil.TimedOutReason {
		return false
	}
	if d.Spec.Replicas != nil && d.Status.UpdatedReplicas < *d.Spec.Replicas {
		return false
	}
	if d.Status.Replicas > d.Status.UpdatedReplicas {
		return false
	}
	if d.Status.AvailableReplicas < d.Status.UpdatedReplicas {
		return false
	}

	return true
}

// +gocode:public-api=true
func isStatefulSetReady(sts *appsv1.StatefulSet) (bool, Controller) {
	ctrl := Controller{
		Name:          sts.Name,
		Namespace:     sts.Namespace,
		Replicas:      sts.Status.Replicas,
		ReadyReplicas: sts.Status.ReadyReplicas,
	}
	if sts.Generation > sts.Status.ObservedGeneration {
		return false, ctrl
	}
	if sts.Status.ReadyReplicas != sts.Status.Replicas {
		return false, ctrl
	}
	if sts.Spec.UpdateStrategy.Type != appsv1.OnDeleteStatefulSetStrategyType && sts.Status.UpdateRevision != sts.Status.CurrentRevision {
		return false, ctrl
	}
	return true, Controller{}
}

// +gocode:public-api=true
func isDaemonSetReady(ds *appsv1.DaemonSet) (bool, DaemonSet) {
	var err error
	var maxUnavailable int
	res := DaemonSet{
		Controller: Controller{
			Name:      ds.Name,
			Namespace: ds.Namespace,
		},
	}
	if ds.Spec.UpdateStrategy.Type == appsv1.RollingUpdateDaemonSetStrategyType {
		maxUnavailable, err = intstr.GetScaledValueFromIntOrPercent(ds.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable, int(ds.Status.DesiredNumberScheduled), true)
		if err != nil {
			maxUnavailable = int(ds.Status.DesiredNumberScheduled)
		}
		res.UpdatedNumberScheduled = ds.Status.UpdatedNumberScheduled
		res.DesiredNumberScheduled = ds.Status.DesiredNumberScheduled
	}
	expectedReady := int(ds.Status.DesiredNumberScheduled) - maxUnavailable
	res.Replicas = int32(expectedReady)
	res.ReadyReplicas = ds.Status.NumberReady

	if ds.Generation > ds.Status.ObservedGeneration {
		return false, res
	}
	if res.ReadyReplicas < res.Replicas || res.UpdatedNumberScheduled < res.DesiredNumberScheduled {
		return false, res
	}
	return true, DaemonSet{}
}

// GetNamespacesList returns list of namespaces from cluster related HelmBundle. The mandatory and excluded values are taken into account
// +gocode:public-api=true
func GetNamespacesList(ctx context.Context, cl client.Client, cluster *clusterv1.Cluster) ([]string, error) {
	namespaces := map[string]bool{}
	for ns := range mandatoryReadinessCheckNamespaces {
		namespaces[ns] = true
	}

	helmBundle := &lcmv1alpha1.HelmBundle{}
	err := cl.Get(ctx, client.ObjectKey{Namespace: cluster.Namespace, Name: cluster.Name}, helmBundle)
	if err != nil {
		if k8serrors.IsNotFound(errors.Cause(err)) {
			return getNamespaces(namespaces), nil
		}
		return nil, errors.Wrap(err, "failed to get HelmBundle object")
	}
	for _, release := range helmBundle.Spec.Releases {
		if !excludedReadinessCheckNamespaces[release.Namespace] {
			namespaces[release.Namespace] = true
		}
	}
	return getNamespaces(namespaces), nil
}

// +gocode:public-api=true
func getNamespaces(nsMap map[string]bool) []string {
	var list []string
	for ns := range nsMap {
		list = append(list, ns)
	}
	return list
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&ClusterPollStatus{}, &ClusterPollStatusList{})
}
