package tags

import (
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// ApplyParams are function parameters used to apply tags on an aws resource.
// +gocode:public-api=true
type ApplyParams struct {
	BuildParams
	EC2Client ec2iface.EC2API
}

// BuildParams is used to build tags around an aws resource.
// +gocode:public-api=true
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

// Apply tags a resource with tags including the cluster tag.
// +gocode:public-api=true
func Apply(params *ApplyParams) error {
	tags := Build(params.BuildParams)

	awsTags := make([]*ec2.Tag, 0, len(tags))
	for k, v := range tags {
		tag := &ec2.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}
		awsTags = append(awsTags, tag)
	}

	createTagsInput := &ec2.CreateTagsInput{
		Resources: aws.StringSlice([]string{params.ResourceID}),
		Tags:      awsTags,
	}

	_, err := params.EC2Client.CreateTags(createTagsInput)
	return errors.Wrapf(err, "failed to tag resource %q in cluster %q", params.ResourceID, params.ClusterID)
}

// Ensure applies the tags if the current tags differ from the params.
// +gocode:public-api=true
func Ensure(current Map, params *ApplyParams) error {
	want := Build(params.BuildParams)
	if !current.Equals(want) {
		return Apply(params)
	}
	return nil
}

// Build builds tags including the cluster tag and returns them in map form.
// +gocode:public-api=true
func Build(params BuildParams) Map {
	tags := make(Map)
	for k, v := range params.Additional {
		tags[k] = v
	}

	if params.Owned == nil || *params.Owned {
		tags[ClusterKey(params.ClusterID)] = string(params.Lifecycle)
	}

	if params.Lifecycle == ResourceLifecycleOwned {
		tags[NameAWSProviderManaged] = "true"
	}

	if params.Role != nil {
		tags[NameAWSClusterAPIRole] = *params.Role
	}

	if params.Name != nil {
		tags["Name"] = *params.Name
	}

	if params.ClusterName != "" {
		tags[NameKaaSCluster] = params.ClusterName
	}

	return tags
}
