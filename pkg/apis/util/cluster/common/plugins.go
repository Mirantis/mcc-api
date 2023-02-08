package common

import (
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sync"
)

var (
	// +gocode:public-api=true
	providersMutex sync.Mutex
	// +gocode:public-api=true
	providers = make(map[string]interface{})
)

// RegisterClusterProvisioner registers a ClusterProvisioner by name.  This
// is expected to happen during app startup.
// +gocode:public-api=true
func RegisterClusterProvisioner(name string, provisioner interface{}) {
	providersMutex.Lock()
	defer providersMutex.Unlock()
	if _, found := providers[name]; found {
		klog.Fatalf("Cluster provisioner %q was registered twice", name)
	}
	klog.V(1).Infof("Registered cluster provisioner %q", name)
	providers[name] = provisioner
}

// +gocode:public-api=true
func ClusterProvisioner(name string) (interface{}, error) {
	providersMutex.Lock()
	defer providersMutex.Unlock()
	provisioner, found := providers[name]
	if !found {
		return nil, errors.Errorf("unable to find provisioner for %s", name)
	}
	return provisioner, nil
}
