package openstack

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/utils/openstack/clientconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/net"

	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/errors"
)

const (
	// RequestTimeout default provider-to-cloud request timeout
	RequestTimeout = 30 * time.Second
)

type Client struct {
	Compute       *gophercloud.ServiceClient
	Network       *gophercloud.ServiceClient
	LoadBalancer  *gophercloud.ServiceClient
	BlockStorage  *gophercloud.ServiceClient
	Orchestration *gophercloud.ServiceClient
	Image         *gophercloud.ServiceClient
}

func NewOpenStackClientFromCloud(cloud *clientconfig.Cloud, requestTimeout time.Duration) (*Client, error) {
	clientOpts := new(clientconfig.ClientOpts)

	if cloud.AuthInfo != nil {
		clientOpts.AuthInfo = cloud.AuthInfo
		clientOpts.AuthType = cloud.AuthType
		clientOpts.Cloud = cloud.Cloud
		clientOpts.RegionName = cloud.RegionName
	}

	opts, err := clientconfig.AuthOptions(clientOpts)
	if err != nil {
		return nil, err
	}
	opts.AllowReauth = true
	provider, err := openstack.NewClient(opts.IdentityEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create provider client: %w", err)
	}
	tlsConf := &tls.Config{InsecureSkipVerify: false}
	if cloud.Verify != nil {
		tlsConf.InsecureSkipVerify = !*cloud.Verify
	}

	if cloud.CACertFile != "" {
		caCert, err := base64.StdEncoding.DecodeString(cloud.CACertFile)
		if err != nil {
			return nil, fmt.Errorf("error reading CA Cert: %w", err)
		}

		caCertPool := x509.NewCertPool()
		success := caCertPool.AppendCertsFromPEM([]byte(caCert))
		if !success {
			return nil, errors.New("failed to parse PEM sertificates by a AppendCertsFromPEM")
		}
		tlsConf.RootCAs = caCertPool
	}

	if cloud.ClientCertFile != "" && cloud.ClientKeyFile != "" {
		clientCert, err := base64.StdEncoding.DecodeString(cloud.ClientCertFile)
		if err != nil {
			return nil, fmt.Errorf("error reading Client Cert: %w", err)
		}
		clientKey, err := base64.StdEncoding.DecodeString(cloud.ClientKeyFile)
		if err != nil {
			return nil, fmt.Errorf("error reading Client Key: %w", err)
		}

		cert, err := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
		if err != nil {
			return nil, err
		}

		tlsConf.Certificates = []tls.Certificate{cert}
	}

	transport := net.SetTransportDefaults(&http.Transport{
		TLSClientConfig: tlsConf,
	})
	provider.HTTPClient = http.Client{
		Transport: transport,
		Timeout:   requestTimeout,
	}
	err = openstack.Authenticate(provider, *opts)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate provider client: %v", err)
	}

	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create compute client: %w", err)
	}
	imageClient, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create image client: %w", err)
	}
	networkingClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create networking client: %w", err)
	}
	lbClient, err := openstack.NewLoadBalancerV2(provider, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create loadbalancer client: %w", err)
	}
	bsClient, err := openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create block storage client: %w", err)
	}
	stackClient, err := openstack.NewOrchestrationV1(provider, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create orchestration client: %w", err)
	}

	return &Client{
		Compute:       computeClient,
		Network:       networkingClient,
		LoadBalancer:  lbClient,
		BlockStorage:  bsClient,
		Orchestration: stackClient,
		Image:         imageClient,
	}, nil
}

func CredentialFromCloud(cloud *clientconfig.Cloud, meta metav1.ObjectMeta) *kaasv1alpha1.OpenStackCredential {
	return &kaasv1alpha1.OpenStackCredential{
		ObjectMeta: meta,
		Spec: kaasv1alpha1.OpenStackCredentialSpec{
			AuthInfo: &kaasv1alpha1.OpenStackAuthInfo{
				AuthURL:  cloud.AuthInfo.AuthURL,
				Username: cloud.AuthInfo.Username,
				UserID:   cloud.AuthInfo.UserID,
				Password: &kaasv1alpha1.SecretValue{
					Value: &cloud.AuthInfo.Password,
				},
				ProjectName:       cloud.AuthInfo.ProjectName,
				ProjectID:         cloud.AuthInfo.ProjectID,
				UserDomainName:    cloud.AuthInfo.UserDomainName,
				UserDomainID:      cloud.AuthInfo.UserDomainID,
				ProjectDomainName: cloud.AuthInfo.ProjectName,
				ProjectDomainID:   cloud.AuthInfo.ProjectDomainID,
				DomainName:        cloud.AuthInfo.DomainName,
				DomainID:          cloud.AuthInfo.DomainID,
				DefaultDomain:     cloud.AuthInfo.DefaultDomain,
			},
			AuthType:   string(cloud.AuthType),
			RegionName: cloud.RegionName,
			CACert:     []byte(cloud.CACertFile),
		},
	}
}
