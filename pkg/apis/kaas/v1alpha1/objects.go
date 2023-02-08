package v1alpha1

import (
	"fmt"
	"strings"
)

type DaemonSet struct {
	Controller             `json:",inline"`
	UpdatedNumberScheduled int32 `json:"updatedNumberScheduled"`
	DesiredNumberScheduled int32 `json:"desiredNumberScheduled"`
}
type Service struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type Objects struct {
	Services     []Service    `json:"services,omitempty"`
	Deployments  []Controller `json:"deployments,omitempty"`
	StatefulSets []Controller `json:"statefulSets,omitempty"`
	DaemonSets   []DaemonSet  `json:"daemonSets,omitempty"`
}

func (o Objects) IsEmpty() bool {
	return len(o.Services) == 0 && len(o.Deployments) == 0 && len(o.StatefulSets) == 0 && len(o.DaemonSets) == 0
}
func (o Objects) Error() string {
	f := func(n, m int) string {
		if m == n {
			return "; "
		}

		return ", "
	}

	message := "not ready: "
	if len(o.Services) > 0 {
		message += "services: "
		for i := range o.Services {
			message += fmt.Sprintf("%s/%s%s", o.Services[i].Namespace, o.Services[i].Name, f(i, len(o.Services)-1))
		}
	}

	if len(o.Deployments) > 0 {
		message += "deployments: "
		for i := range o.Deployments {
			message += fmt.Sprintf("%s/%s got %d/%d replicas%s",
				o.Deployments[i].Namespace, o.Deployments[i].Name, o.Deployments[i].ReadyReplicas, o.Deployments[i].Replicas, f(i, len(o.Deployments)-1))
		}
	}

	if len(o.StatefulSets) > 0 {
		message += "statefulSets: "
		for i := range o.StatefulSets {
			message += fmt.Sprintf("%s/%s got %d/%d replicas%s",
				o.StatefulSets[i].Namespace, o.StatefulSets[i].Name, o.StatefulSets[i].ReadyReplicas, o.StatefulSets[i].Replicas, f(i, len(o.StatefulSets)-1))
		}
	}

	if len(o.DaemonSets) > 0 {
		message += "daemonSets: "
		for i := range o.DaemonSets {
			message += fmt.Sprintf("%s/%s ", o.DaemonSets[i].Namespace, o.DaemonSets[i].Name)
			if o.DaemonSets[i].UpdatedNumberScheduled < o.DaemonSets[i].DesiredNumberScheduled {
				message += fmt.Sprintf("scheduled %d/%d replicas%s",
					o.DaemonSets[i].UpdatedNumberScheduled, o.DaemonSets[i].DesiredNumberScheduled, f(i, len(o.DaemonSets)-1))
			}
			if o.DaemonSets[i].ReadyReplicas < o.DaemonSets[i].Replicas {
				message += fmt.Sprintf("ready %d/%d replicas%s",
					o.DaemonSets[i].ReadyReplicas, o.DaemonSets[i].Replicas, f(i, len(o.DaemonSets)-1))
			}
		}
	}

	return strings.TrimSuffix(message, "; ")
}
func (o Objects) StacklightObjects() Objects {
	res := Objects{}
	stacklightNamespace := "stacklight"
	for _, service := range o.Services {
		if service.Namespace == stacklightNamespace {
			res.Services = append(res.Services, service)
		}
	}
	for _, deployment := range o.Deployments {
		if deployment.Namespace == stacklightNamespace {
			res.Deployments = append(res.Deployments, deployment)
		}
	}
	for _, statefulset := range o.StatefulSets {
		if statefulset.Namespace == stacklightNamespace {
			res.StatefulSets = append(res.StatefulSets, statefulset)
		}
	}
	for _, daemonset := range o.DaemonSets {
		if daemonset.Namespace == stacklightNamespace {
			res.DaemonSets = append(res.DaemonSets, daemonset)
		}
	}
	return res
}

type Controller struct {
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	Replicas      int32  `json:"replicas"`
	ReadyReplicas int32  `json:"readyReplicas"`
}
