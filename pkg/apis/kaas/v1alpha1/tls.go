package v1alpha1

import (
	"reflect"
	"strings"
)

// +gocode:public-api=true
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
type TLSSpec struct {
	//[Deprecated] The desired name of the application
	Hostname string `json:"hostname,omitempty" sensitive:"true"`
	//[Deprecated] The reference to the Certificate object
	Certificate CertificateRef `json:"certificate,omitempty"`
	// TLSConfigRef is a reference to TLCConfig for the application
	TLSConfigRef string `json:"tlsConfigRef"`
}
type TLSStatus struct {
	// ExpirationTime indicates the end of validity period of a certificate
	ExpirationTime string `json:"expirationTime"`
	// RenewalTime indicates the date MCC controller attempts to request a new certificate
	RenewalTime string `json:"renewalTime,omitempty"`
	// Hostname indicates protected by the certificate server name
	Hostname string `json:"hostname,omitempty"`
}
type CertificateRef struct {
	// The name of the Certificate object
	Name string `json:"name,omitempty"`
}

// +gocode:public-api=true
type TLS struct {
	Keycloak             *TLSSpec `json:"keycloak,omitempty"`
	UI                   *TLSSpec `json:"ui,omitempty"`
	MKE                  *TLSSpec `json:"mke,omitempty"`
	Cache                *TLSSpec `json:"cache,omitempty"`
	IAMProxyAlerta       *TLSSpec `json:"iamProxyAlerta,omitempty"`
	IAMProxyAlertManager *TLSSpec `json:"iamProxyAlertManager,omitempty"`
	IAMProxyGrafana      *TLSSpec `json:"iamProxyGrafana,omitempty"`
	IAMProxyKibana       *TLSSpec `json:"iamProxyKibana,omitempty"`
	IAMProxyPrometheus   *TLSSpec `json:"iamProxyPrometheus,omitempty"`
}

// ToMap returns a map of applications where the key is a name and the value is a tls spec defined in the struct for the app
func (tls *TLS) ToMap() map[string]*TLSSpec {
	res := map[string]*TLSSpec{}
	v := reflect.ValueOf(*tls)

	for i := 0; i < v.NumField(); i++ {
		spec, ok := v.Field(i).Interface().(*TLSSpec)
		if !ok {
			continue
		}
		appName := toAppName(v.Type().Field(i).Name)
		res[strings.ToLower(appName)] = spec
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

// +gocode:public-api=true
func toAppName(fieldName string) string {
	if strings.HasPrefix(fieldName, "IAMProxy") {
		return strings.Replace(fieldName, "IAMProxy", "IAM-Proxy-", 1)
	}
	return fieldName
}
