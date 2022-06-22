package v1alpha1

import (
	"context"

	crclient "sigs.k8s.io/controller-runtime/pkg/client"

	perrors "github.com/Mirantis/mcc-api/pkg/errors"
)

type StackLightStatus struct {
	Prometheus      StackLightComponentStatus `json:"prometheus,omitempty"`
	Alerta          StackLightComponentStatus `json:"alerta,omitempty"`
	AlertManager    StackLightComponentStatus `json:"alertmanager,omitempty"`
	Grafana         StackLightComponentStatus `json:"grafana,omitempty"`
	Kibana          StackLightComponentStatus `json:"kibana,omitempty"`
	TelemeterServer StackLightComponentStatus `json:"telemeterServer,omitempty"`
}

type StackLightComponentStatus struct {
	ComponentStatus `json:",inline"`
}

func (st *StackLightStatus) UpdateStatus(ctx context.Context, client crclient.Client, iamScheme string) error {
	errs := perrors.NewErrorCollector("multiple errors when updating StackLight components status")

	err := st.Alerta.UpdateURL(ctx, client, "stacklight", "iam-proxy-alerta", iamScheme, 0)
	errs.Collect(err, "failed to update Alerta status")

	err = st.AlertManager.UpdateURL(ctx, client, "stacklight", "iam-proxy-alertmanager", iamScheme, 0)
	errs.Collect(err, "failed to update Alert Manager status")

	err = st.Grafana.UpdateURL(ctx, client, "stacklight", "iam-proxy-grafana", iamScheme, 0)
	errs.Collect(err, "failed to update Grafana status")

	err = st.Kibana.UpdateURL(ctx, client, "stacklight", "iam-proxy-kibana", iamScheme, 0)
	errs.Collect(err, "failed to update Kibana status")

	err = st.Prometheus.UpdateURL(ctx, client, "stacklight", "iam-proxy-prometheus", iamScheme, 0)
	errs.Collect(err, "failed to update Prometheus status")

	err = st.TelemeterServer.UpdateURL(ctx, client, "stacklight", "telemeter-server-external", "https", 0)
	errs.Collect(err, "failed to update TelemeterServer status")

	return errs.GetError()
}
