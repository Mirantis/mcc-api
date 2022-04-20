/*
Copyright 2022 The Kubernetes Authors.

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

package tags

import (
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// ApplyParams are function parameters used to apply tags on an aws resource.
type ApplyParams struct {
	BuildParams
	EC2Client ec2iface.EC2API
}

// BuildParams is used to build tags around an aws resource.
type BuildParams struct {
	// Lifecycle determines the resource lifecycle.
	Lifecycle ResourceLifecycle

	// ClusterName is the cluster name associated with the resource.
	ClusterName string

	// ClusterID is the cluster id associated with the resource.
	ClusterID string

	// ResourceID is the unique identifier of the resource to be tagged.
	ResourceID string

	// Name is the name of the resource, it's applied as the tag "Name" on AWS.
	// +optional
	Name *string

	// Role is the role associated to the resource.
	// +optional
	Role *string

	Owned *bool

	// Any additional tags to be added to the resource.
	// +optional
	Additional Map
}
