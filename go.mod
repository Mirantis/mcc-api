module github.com/Mirantis/mcc-api

go 1.15

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
	github.com/Masterminds/semver v1.5.0
	github.com/aws/aws-sdk-go v1.37.19
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/elastic/go-sysinfo v1.7.1
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/google/uuid v1.2.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kube-object-storage/lib-bucket-provisioner v0.0.0-20220105185820-c1da9586e05b
	github.com/packethost/packngo v0.22.0
	github.com/pkg/errors v0.9.1
	github.com/rook/rook v1.8.8
	github.com/stretchr/testify v1.7.0
	github.com/thoas/go-funk v0.9.2
	go.uber.org/multierr v1.8.0
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	k8s.io/api v0.23.5
	k8s.io/apiextensions-apiserver v0.20.1
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/helm v2.17.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/kubectl v0.23.5
	sigs.k8s.io/controller-runtime v0.10.2
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.3.0
)
