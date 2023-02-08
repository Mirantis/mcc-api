module github.com/Mirantis/mcc-api/v2

go 1.15

replace (
	github.com/go-logr/logr => github.com/go-logr/logr v0.4.0
	github.com/go-logr/zapr => github.com/go-logr/zapr v0.4.0
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.21.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.2
	k8s.io/client-go => k8s.io/client-go v0.20.2
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.10.0
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
	k8s.io/kubectl => k8s.io/kubectl v0.20.2
	k8s.io/kubelet => k8s.io/kubelet v0.20.2
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.8.3
)

require (
	github.com/Masterminds/semver v1.5.0
	github.com/aws/aws-sdk-go v1.44.113
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/elastic/go-sysinfo v1.8.1
	github.com/go-kit/log v0.2.1
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/google/uuid v1.3.0
	github.com/metal3-io/baremetal-operator/pkg/hardwareutils v0.2.0
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.21.1 // indirect
	github.com/packethost/packngo v0.28.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/thoas/go-funk v0.9.2
	go.uber.org/multierr v1.8.0
	golang.org/x/net v0.0.0-20221004154528-8021a29435af // indirect
	k8s.io/api v0.25.0
	k8s.io/apiextensions-apiserver v0.20.1
	k8s.io/apimachinery v0.25.0
	k8s.io/client-go v0.21.3
	k8s.io/helm v2.17.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1 // indirect
	k8s.io/kubectl v0.0.0-00010101000000-000000000000
	k8s.io/utils v0.0.0-20220922133306-665eaaec4324
	sigs.k8s.io/controller-runtime v0.0.0-00010101000000-000000000000
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0
)
