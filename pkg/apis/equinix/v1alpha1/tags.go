package v1alpha1

const (
	// +gocode:public-api=true
	ControlPlaneTag = "kubernetes.io/role:master"
	// +gocode:public-api=true
	WorkerTag = "kubernetes.io/role:node"
)
