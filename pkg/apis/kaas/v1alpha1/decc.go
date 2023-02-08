package v1alpha1

type CacheComponentStatus struct {
	ComponentStatus `json:",inline"`
}
type KaaSUIComponentStatus struct {
	ComponentStatus `json:",inline"`
}

// +gocode:public-api=true
type DECCStatus struct {
	UI    KaaSUIComponentStatus `json:"ui,omitempty"`
	Cache CacheComponentStatus  `json:"cache,omitempty"`
	Proxy ProxyComponentStatus  `json:"proxy,omitempty"`
}
type ProxyComponentStatus struct {
	ComponentStatus `json:",inline"`
}
