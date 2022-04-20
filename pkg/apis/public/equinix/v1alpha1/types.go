/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// EquinixMetalResourceStatus describes the status of a EquinixMetal resource.
type EquinixMetalResourceStatus string

var (
	// EquinixMetalResourceStatusNew represents an EquinixMetal resource requested.
	// The EquinixMetal infrastructure uses a queue to avoid any abuse. So a resource
	// does not get created straight away but it can wait for a bit in a queue.
	EquinixMetalResourceStatusNew = EquinixMetalResourceStatus("new")
	// EquinixMetalResourceStatusQueued represents a device waiting for his turn to be provisioned.
	// Time in queue depends on how many creation requests you already issued, or
	// from how many resources waiting to be deleted we have for you.
	EquinixMetalResourceStatusQueued = EquinixMetalResourceStatus("queued")
	// EquinixMetalResourceStatusProvisioning represents a resource that got dequeued
	// and it is actively processed by a worker.
	EquinixMetalResourceStatusProvisioning = EquinixMetalResourceStatus("provisioning")
	// EquinixMetalResourceStatusRunning represents an EquinixMetal resource already provisioned and in a active state.
	EquinixMetalResourceStatusRunning = EquinixMetalResourceStatus("active")
	// EquinixMetalResourceStatusErrored represents an EquinixMetal resource in a errored state.
	EquinixMetalResourceStatusErrored = EquinixMetalResourceStatus("errored")
	// EquinixMetalResourceStatusOff represents an EquinixMetal resource in off state.
	EquinixMetalResourceStatusOff = EquinixMetalResourceStatus("off")
)

// the purpose of Elastic IP (block) resource, i.e., how it will be used in a cluster.
type ElasticIPPurpose string

const ElasticIPPurposeTagPrefix = "kaas.mirantis.com/purpose:"

const (
	// for k8s API and UCP API LB
	ElasticIPPurposeHostLB = ElasticIPPurpose("host-lb")
	// for k8s services LBs
	ElasticIPPurposeServiceLB = ElasticIPPurpose("service-lb")
	// for needs of MOS cluster
	ElasticIPPurposeMOS = ElasticIPPurpose("mos")
)

// Tags defines a slice of tags.
type Tags []string
