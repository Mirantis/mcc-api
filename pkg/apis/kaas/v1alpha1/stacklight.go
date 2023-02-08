package v1alpha1

// +gocode:public-api=true
type StackLightStatus struct {
	Prometheus      StackLightComponentStatus `json:"prometheus,omitempty"`
	Alerta          StackLightComponentStatus `json:"alerta,omitempty"`
	AlertManager    StackLightComponentStatus `json:"alertmanager,omitempty"`
	Grafana         StackLightComponentStatus `json:"grafana,omitempty"`
	Kibana          StackLightComponentStatus `json:"kibana,omitempty"`
	TelemeterServer StackLightComponentStatus `json:"telemeterServer,omitempty"`
}

// +gocode:public-api=true
type StackLightComponentStatus struct {
	ComponentStatus `json:",inline"`
}
