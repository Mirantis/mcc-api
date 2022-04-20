/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/Mirantis/mcc-api/pkg/client/internalclientset/typed/kaas/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKaasV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKaasV1alpha1) AWSCredentials(namespace string) v1alpha1.AWSCredentialInterface {
	return &FakeAWSCredentials{c, namespace}
}

func (c *FakeKaasV1alpha1) AWSResourceses(namespace string) v1alpha1.AWSResourcesInterface {
	return &FakeAWSResourceses{c, namespace}
}

func (c *FakeKaasV1alpha1) AzureCredentials(namespace string) v1alpha1.AzureCredentialInterface {
	return &FakeAzureCredentials{c, namespace}
}

func (c *FakeKaasV1alpha1) AzureResourceses(namespace string) v1alpha1.AzureResourcesInterface {
	return &FakeAzureResourceses{c, namespace}
}

func (c *FakeKaasV1alpha1) BYOCredentials(namespace string) v1alpha1.BYOCredentialInterface {
	return &FakeBYOCredentials{c, namespace}
}

func (c *FakeKaasV1alpha1) Certificates(namespace string) v1alpha1.CertificateInterface {
	return &FakeCertificates{c, namespace}
}

func (c *FakeKaasV1alpha1) CertificateConfigurations(namespace string) v1alpha1.CertificateConfigurationInterface {
	return &FakeCertificateConfigurations{c, namespace}
}

func (c *FakeKaasV1alpha1) ClusterReleases() v1alpha1.ClusterReleaseInterface {
	return &FakeClusterReleases{c}
}

func (c *FakeKaasV1alpha1) EquinixMetalCredentials(namespace string) v1alpha1.EquinixMetalCredentialInterface {
	return &FakeEquinixMetalCredentials{c, namespace}
}

func (c *FakeKaasV1alpha1) EquinixMetalResourceses(namespace string) v1alpha1.EquinixMetalResourcesInterface {
	return &FakeEquinixMetalResourceses{c, namespace}
}

func (c *FakeKaasV1alpha1) KaaSCephClusters(namespace string) v1alpha1.KaaSCephClusterInterface {
	return &FakeKaaSCephClusters{c, namespace}
}

func (c *FakeKaasV1alpha1) KaaSCephOperationRequests(namespace string) v1alpha1.KaaSCephOperationRequestInterface {
	return &FakeKaaSCephOperationRequests{c, namespace}
}

func (c *FakeKaasV1alpha1) KaaSReleases() v1alpha1.KaaSReleaseInterface {
	return &FakeKaaSReleases{c}
}

func (c *FakeKaasV1alpha1) Licenses() v1alpha1.LicenseInterface {
	return &FakeLicenses{c}
}

func (c *FakeKaasV1alpha1) MCCCertificateRequests(namespace string) v1alpha1.MCCCertificateRequestInterface {
	return &FakeMCCCertificateRequests{c, namespace}
}

func (c *FakeKaasV1alpha1) MCCUpgrades() v1alpha1.MCCUpgradeInterface {
	return &FakeMCCUpgrades{c}
}

func (c *FakeKaasV1alpha1) OpenStackCredentials(namespace string) v1alpha1.OpenStackCredentialInterface {
	return &FakeOpenStackCredentials{c, namespace}
}

func (c *FakeKaasV1alpha1) OpenStackResourceses(namespace string) v1alpha1.OpenStackResourcesInterface {
	return &FakeOpenStackResourceses{c, namespace}
}

func (c *FakeKaasV1alpha1) Proxies(namespace string) v1alpha1.ProxyInterface {
	return &FakeProxies{c, namespace}
}

func (c *FakeKaasV1alpha1) PublicKeys(namespace string) v1alpha1.PublicKeyInterface {
	return &FakePublicKeys{c, namespace}
}

func (c *FakeKaasV1alpha1) RHELLicenses(namespace string) v1alpha1.RHELLicenseInterface {
	return &FakeRHELLicenses{c, namespace}
}

func (c *FakeKaasV1alpha1) UnsupportedClusterses() v1alpha1.UnsupportedClustersInterface {
	return &FakeUnsupportedClusterses{c}
}

func (c *FakeKaasV1alpha1) VsphereCredentials(namespace string) v1alpha1.VsphereCredentialInterface {
	return &FakeVsphereCredentials{c, namespace}
}

func (c *FakeKaasV1alpha1) VsphereResourceses(namespace string) v1alpha1.VsphereResourcesInterface {
	return &FakeVsphereResourceses{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKaasV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
