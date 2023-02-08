package errors

import (
	"fmt"
	commonerrors "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/common"
)

// A more descriptive kind of error that represents an error condition that
// should be set in the Cluster.Status. The "Reason" field is meant for short,
// enum-style constants meant to be interpreted by clusters. The "Message"
// field is meant to be read by humans.
// +gocode:public-api=true
type ClusterError struct {
	Reason  commonerrors.ClusterStatusError
	Message string
}

func (e *ClusterError) Error() string {
	return e.Message
}

// Some error builders for ease of use. They set the appropriate "Reason"
// value, and all arguments are Printf-style varargs fed into Sprintf to
// construct the Message.
// +gocode:public-api=true
func InvalidClusterConfiguration(format string, args ...interface{}) *ClusterError {
	return &ClusterError{
		Reason:  commonerrors.InvalidConfigurationClusterError,
		Message: fmt.Sprintf(format, args...),
	}
}

// +gocode:public-api=true
func CreateCluster(format string, args ...interface{}) *ClusterError {
	return &ClusterError{
		Reason:  commonerrors.CreateClusterError,
		Message: fmt.Sprintf(format, args...),
	}
}

// +gocode:public-api=true
func DeleteCluster(format string, args ...interface{}) *ClusterError {
	return &ClusterError{
		Reason:  commonerrors.DeleteClusterError,
		Message: fmt.Sprintf(format, args...),
	}
}
