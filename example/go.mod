module github.com/Mirantis/mcc-api/example

go 1.16

// Kubernetes libraries
replace (
	github.com/go-logr/logr => github.com/go-logr/logr v0.4.0
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.2
	// libs pinned until https://github.com/rook/rook/pull/7913 resolved.
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.2
	k8s.io/client-go => k8s.io/client-go v0.20.2
	k8s.io/code-generator => k8s.io/code-generator v0.20.2
	k8s.io/component-base => k8s.io/component-base v0.20.2
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.10.0
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20201113171705-d219536bb9fd
	k8s.io/kubectl => k8s.io/kubectl v0.20.2
	k8s.io/kubernetes => k8s.io/kubernetes v0.21.1
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.7.0
	sigs.k8s.io/kustomize => sigs.k8s.io/kustomize v2.0.3+incompatible
)

// MiraCeph controller libraries
replace (
	// pinned to avoid OpenStack module failure after libs upgrade
	github.com/gophercloud/gophercloud => github.com/gophercloud/gophercloud v0.21.0
	github.com/kubernetes-incubator/external-storage => github.com/libopenstorage/external-storage v0.20.4-openstorage-rc3
	github.com/portworx/sched-ops => github.com/portworx/sched-ops v0.20.4-openstorage-rc3
)

require (
	github.com/spf13/cobra v1.4.0
	k8s.io/component-base v0.23.6
)

require (
	github.com/Mirantis/mcc-api v0.0.0-20220420210440-795b4750e3ba
	github.com/Nerzal/gocloak/v7 v7.11.0
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/gophercloud/gophercloud v0.20.0
	github.com/gophercloud/utils v0.0.0-20220307143606-8e7800759d16
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29
	golang.org/x/net v0.0.0-20211209124913-491a49abca63
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.6
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/klog/v2 v2.30.0
)
