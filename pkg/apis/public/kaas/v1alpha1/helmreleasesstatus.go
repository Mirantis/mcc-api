package v1alpha1

import (
	"context"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Mirantis/mcc-api/v2/pkg/util/k8s"
)

type HelmReleasesStatus struct {
	StackLight      StackLightStatus         `json:"stacklight,omitempty"`
	IAM             IAMStatus                `json:"iam,omitempty"`
	DECC            DECCStatus               `json:"decc,omitempty"`
	ReleaseStatuses map[string]ReleaseStatus `json:"releaseStatuses,omitempty"`
}

type ReleaseStatus struct {
	Success bool `json:"success,omitempty"`
}

type ComponentStatus struct {
	URL string `json:"url,omitempty"`
}

func formatURL(scheme, host string, port int) *url.URL {
	if port > 0 {
		host = net.JoinHostPort(host, strconv.Itoa(port))
	}
	return &url.URL{Scheme: scheme, Host: host}
}

func (cs *ComponentStatus) UpdateURL(ctx context.Context, client crclient.Client, serviceNamespace, serviceName string, scheme string, port int) error {
	host, err := k8s.ServiceExternalAddress(ctx, client, serviceNamespace, serviceName)
	if err != nil {
		if k8serrors.IsNotFound(errors.Cause(err)) {
			cs.URL = ""
			return nil
		}
		return err
	}
	cs.URL = formatURL(scheme, host, port).String()
	return nil
}

type ComponentStatusGetter interface {
	GetStatus(client crclient.Client, serviceName string) ComponentStatus
}

func (s HelmReleasesStatus) Error() string {
	message := "not ready helm releases: "
	for release, status := range s.ReleaseStatuses {
		if !status.Success {
			message += release + ", "
		}
	}
	return strings.TrimSuffix(message, ", ")
}
