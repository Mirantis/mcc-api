package k8s

import (
	"context"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func ServiceExternalAddress(client crclient.Client, namespace, name string) (string, error) {
	var address string
	var svc corev1.Service
	err := client.Get(context.TODO(), crclient.ObjectKey{
		Namespace: namespace,
		Name:      name,
	}, &svc)
	if err != nil {
		return address, errors.Wrapf(err, "failed to get service %v/%v", namespace, name)
	}

	for _, in := range svc.Status.LoadBalancer.Ingress {
		if in.Hostname != "" {
			address = in.Hostname
			break
		}
		if in.IP != "" {
			address = in.IP
			break
		}
	}
	if address == "" {
		return address, errors.Errorf("no valid loadbalancer ingress points found for service %s/%s", namespace, name)
	}

	return address, nil
}
