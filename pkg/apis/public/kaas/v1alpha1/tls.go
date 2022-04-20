package v1alpha1

import (
	"reflect"
	"strings"
)

type TLS struct {
	Keycloak *TLSSpec `json:"keycloak,omitempty"`
	UI       *TLSSpec `json:"ui,omitempty"`
}

//ToMap returns a map of applications where the key is a name and the value is a tls spec defined in the struct for the app
func (tls *TLS) ToMap() map[string]*TLSSpec {
	res := map[string]*TLSSpec{}
	v := reflect.ValueOf(*tls)

	for i := 0; i < v.NumField(); i++ {
		spec, ok := v.Field(i).Interface().(*TLSSpec)
		if !ok {
			continue
		}
		res[strings.ToLower(v.Type().Field(i).Name)] = spec
	}
	return res
}

func (tls *TLS) SupportedApplications() []string {
	res := []string{}
	for app := range tls.ToMap() {
		res = append(res, app)
	}
	return res
}

func (tls *TLS) IsSupported(app string) bool {
	for _, v := range tls.SupportedApplications() {
		if v == app {
			return true
		}
	}
	return false
}

type TLSSpec struct {
	// The desired name of the application
	Hostname string `json:"hostname"`
	// The reference to the Certificate object
	Certificate CertificateRef `json:"certificate"`
}

type CertificateRef struct {
	// The name of the Certificate object
	Name string `json:"name"`
}

type MCCTLSStatus struct {
	Keycloak             *TLSStatus `json:"keycloak,omitempty"`
	UI                   *TLSStatus `json:"ui,omitempty"`
	Admission            *TLSStatus `json:"admission,omitempty"`
	Cache                *TLSStatus `json:"cache,omitempty"`
	TelemeterServer      *TLSStatus `json:"telemeterServer,omitempty"`
	IAMProxyAlerta       *TLSStatus `json:"iamProxyAlerta,omitempty"`
	IAMProxyAlertManager *TLSStatus `json:"iamProxyAlertManager,omitempty"`
	IAMProxyGrafana      *TLSStatus `json:"iamProxyGrafana,omitempty"`
	IAMProxyKibana       *TLSStatus `json:"iamProxyKibana,omitempty"`
	IAMProxyPrometheus   *TLSStatus `json:"iamProxyPrometheus,omitempty"`
}

type TLSStatus struct {
	// ExpirationTime indicates the end of validity period of a certificate
	ExpirationTime string `json:"expirationTime"`
	// RenewalTime indicates the date MCC controller attempts to request a new certificate
	RenewalTime string `json:"renewalTime,omitempty"`
	// Hostname indicates protected by the certificate server name
	Hostname string `json:"hostname,omitempty"`
}
