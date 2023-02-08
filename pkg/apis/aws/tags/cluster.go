package tags

import (
	"fmt"
)

// ClusterKey generates the key for resources associated with a cluster.
// +gocode:public-api=true
func ClusterKey(name string) string {
	return fmt.Sprintf("%s%s", NameKubernetesClusterPrefix, name)
}
