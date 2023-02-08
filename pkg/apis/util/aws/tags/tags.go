package tags

import (
	"github.com/Mirantis/mcc-api/v2/pkg/apis/aws/tags"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

// Apply tags a resource with tags including the cluster tag.
// +gocode:public-api=true
func Apply(params *tags.ApplyParams) error {
	tags := tags.Build(params.BuildParams)

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
func Ensure(current tags.Map, params *tags.ApplyParams) error {
	want := tags.Build(params.BuildParams)
	if !current.Equals(want) {
		return Apply(params)
	}
	return nil
}
