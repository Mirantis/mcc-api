package v1alpha1

import (
	crclient "sigs.k8s.io/controller-runtime/pkg/client"

	perrors "github.com/Mirantis/mcc-api/pkg/errors"
)

type IAMStatus struct {
	Keycloak IAMComponentStatus `json:"keycloak,omitempty"`
	API      APIComponentStatus `json:"api,omitempty"`
}

type IAMComponentStatus struct {
	ComponentStatus `json:",inline"`
}

type APIComponentStatus struct {
	ComponentStatus `json:",inline"`
}

func (st *IAMStatus) UpdateStatus(client crclient.Client, spec *ClusterSpecMixin) error {
	var err error
	errs := perrors.NewErrorCollector("multiple errors during status updating")

	if spec.TLS.Keycloak == nil {
		err = st.Keycloak.UpdateURL(client, "kaas", "iam-keycloak-http", "https", 0)
		errs.Collect(err, "failed to update keycloak status")
	} else {
		st.Keycloak.URL = formatURL("https", spec.TLS.Keycloak.Hostname, 0).String()
	}

	err = st.API.UpdateURL(client, "kaas", "iam-api-http", "https", 0)
	errs.Collect(err, "failed to update IAM API status")

	return errs.GetError()
}
