package v1alpha1

import (
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type DECCStatus struct {
	UI    KaaSUIComponentStatus `json:"ui,omitempty"`
	Cache CacheComponentStatus  `json:"cache,omitempty"`
	Proxy ProxyComponentStatus  `json:"proxy,omitempty"`
}

type KaaSUIComponentStatus struct {
	ComponentStatus `json:",inline"`
}

func (st *DECCStatus) UpdateUIStatus(client crclient.Client, spec *ClusterSpecMixin, scheme string) error {
	if spec.TLS.UI == nil {
		return st.UI.UpdateURL(client, "kaas", "kaas-kaas-ui", scheme, 0)
	}
	st.UI.URL = formatURL(scheme, spec.TLS.UI.Hostname, 0).String()
	return nil
}

type CacheComponentStatus struct {
	ComponentStatus `json:",inline"`
}

func (st *DECCStatus) UpdateCacheStatus(client crclient.Client) error {
	return st.Cache.UpdateURL(client, "kaas", "mcc-cache", "https", 0)
}

type ProxyComponentStatus struct {
	ComponentStatus `json:",inline"`
}

func (st *DECCStatus) UpdateProxyStatus(client crclient.Client) error {
	return st.Proxy.UpdateURL(client, "kaas", "squid-proxy", "http", 3128)
}
