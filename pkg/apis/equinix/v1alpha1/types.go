package v1alpha1

var (
	// EquinixMetalResourceStatusNew represents an EquinixMetal resource requested.
	// The EquinixMetal infrastructure uses a queue to avoid any abuse. So a resource
	// does not get created straight away but it can wait for a bit in a queue.
	// +gocode:public-api=true
	EquinixMetalResourceStatusNew = EquinixMetalResourceStatus("new")
	// EquinixMetalResourceStatusQueued represents a device waiting for his turn to be provisioned.
	// Time in queue depends on how many creation requests you already issued, or
	// from how many resources waiting to be deleted we have for you.
	// +gocode:public-api=true
	EquinixMetalResourceStatusQueued = EquinixMetalResourceStatus("queued")
	// EquinixMetalResourceStatusProvisioning represents a resource that got dequeued
	// and it is actively processed by a worker.
	// +gocode:public-api=true
	EquinixMetalResourceStatusProvisioning = EquinixMetalResourceStatus("provisioning")
	// EquinixMetalResourceStatusRunning represents an EquinixMetal resource already provisioned and in a active state.
	// +gocode:public-api=true
	EquinixMetalResourceStatusRunning = EquinixMetalResourceStatus("active")
	// EquinixMetalResourceStatusErrored represents an EquinixMetal resource in a errored state.
	// +gocode:public-api=true
	EquinixMetalResourceStatusErrored = EquinixMetalResourceStatus("errored")
	// EquinixMetalResourceStatusOff represents an EquinixMetal resource in off state.
	// +gocode:public-api=true
	EquinixMetalResourceStatusOff = EquinixMetalResourceStatus("off")
)

// the purpose of Elastic IP (block) resource, i.e., how it will be used in a cluster.
// +gocode:public-api=true
type ElasticIPPurpose string

const (
	// for k8s API and UCP API LB
	// +gocode:public-api=true
	ElasticIPPurposeHostLB = ElasticIPPurpose("host-lb")
	// for k8s services LBs
	// +gocode:public-api=true
	ElasticIPPurposeServiceLB = ElasticIPPurpose("service-lb")
	// for needs of MOS cluster
	// +gocode:public-api=true
	ElasticIPPurposeMOS = ElasticIPPurpose("mos")
)

// Tags defines a slice of tags.
// +gocode:public-api=true
type Tags []string

// EquinixMetalResourceStatus describes the status of a EquinixMetal resource.
// +gocode:public-api=true
type EquinixMetalResourceStatus string

// +gocode:public-api=true
const ElasticIPPurposeTagPrefix = "kaas.mirantis.com/purpose:"
