## About

This is an example of using the MCC Public API. Implemented functionality allows to create managed cluster over existing MCC management cluster and make some operations with it one, such as upgrade and delete. 

## How to build own client

Client creation example for a managed cluster operations via MCC Public API is implemented in this _example_ folder. As the result the binary file with some _cmd_ UI will be prepared. See description [below](#example-client-use).
In this section the API client calls and used types details are pointed up.

Example implementation is based on two golang types:

-  The `Cluster` implements cluster related operations:
```
type Cluster struct {
	kubeClient     kubernetes.Interface
	clientSet      *internalclientset.Clientset
	kubeConfig     *rest.Config
	isWaitInfinite bool // for PollImmidiate  if need to wait under operator's control
}

type Interface interface {
	EnsureNamespace(namespaceName string) error
	setKubeClient() error
	GetOpenStackCredentialWithRetry(name, namespace string) (*kaasv1alpha1.OpenStackCredential, error)
	CreateOpenStackCredential(credential *kaasv1alpha1.OpenStackCredential) error
	CreatePublicKey(publicKey *kaasv1alpha1.PublicKey) error
	CreateClusterObject(cluster *clusterv1.Cluster) error
	CreateMachines(machines []*clusterv1.Machine, namespace string) error
	WaitForMachinesFold(interval, duration time.Duration, namespace, clusterName string, initialValue bool, foldfunc func(bool, *clusterv1.Machine, *kaasv1alpha1.MachineStatusMixin) bool) error
	GetMachinesForCluster(cluster *clusterv1.Cluster) ([]*clusterv1.Machine, error)
	PatchCluster(name, namespace string, data []byte) error
	DeleteCluster(namespace, name string) error
	waitForClusterDelete(namespace, name string) error
        GetCluster(namespace, name string) (*clusterv1.Cluster, error) 
}
```
Code is [there](lib/cluster.go)

-  The `MccClient` is used for cluster operations steps implementation:
```
type MccClient struct {
	managementCluster *exampleLib.Cluster
	isWaitInfinite    bool
}

type Interface interface {
	CloudConfigExists(name, namespace string) (bool, error)
	CreateCloudConfig(region, name, namespace string, cloudConfig *clientconfig.Cloud) error
	CreateManagedCluster(cluster *clusterv1.Cluster, machines []*clusterv1.Machine, publicKey *kaasv1alpha1.PublicKey) error
	waitForMachinesReady(namespace, clusterName string) error
	waitForOneMachineNotReady(namespace, clusterName string) error
	GetKubeconfig(kubeconfigOutput string, clusterName, namespace, realm, username, password string) error
	Upgrade(clusterName, namespace, releaseName string) error
	Delete(clusterName, namespace string) error
}
```
Code is [there](lib/deployer/deployer.go)

### Using the API structures and clients 
### Connect to a management cluster

Management cluster object is used for all managed cluster creation steps described below. As a main condition the management kubeconfig should be present.
Reaching different cluster related objects arte based on a management cluster connection.
                          
Please, see the [example](lib/cluster.go) with the function 'ConnectToCluster'.

### Managed cluster creation

__Cloud credentials__
The cloud config (credentials) is used while a new managed cluster creation.

The target namespace should contain this object.

```
    import (
      metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
      "k8s.io/client-go/rest"
      "github.com/gophercloud/utils/openstack/clientconfig"
      
      "github.com/Mirantis/mcc-api/pkg/client/internalclientset"
      clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
    )  

    ...
    
    var clientconfig.Cloud cloud // Unmarshal from cloud.yaml

    credentials := &kaasv1alpha1.OpenStackCredential{
		ObjectMeta: metav1.ObjectMeta{
		      Name:      name,
	      	Namespace: namespace,
	      	Labels: map[string]string{
			"kaas.mirantis.com/region": region,
	  	},
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
  // MCC API client
  clientSet :=  internalclientset.NewForConfig(rest.kubeConfig)

  _, err := clientSet.KaasV1alpha1().OpenStackCredentials(credential.Namespace).Create(context.TODO(), credential, metav1.CreateOptions{})
```

The next code allows create a new __managed cluster object__:

```
    import (
      metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
      "k8s.io/client-go/rest"
      
      "github.com/Mirantis/mcc-api/pkg/client/internalclientset"
      clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
    )
    ....

    // MCC API client
    clientSet :=  internalclientset.NewForConfig(rest.kubeConfig)
   
    // MCC API cluster object. See details below
    var  v1alpha1 clusterv1.Cluster cluster 

    _, err := clientSet.ClusterV1alpha1().Clusters(namespace).Create(context.TODO(), cluster, metav1.CreateOptions{})

```


Before managed cluster related machines are created we need __SSH public key__ to be ensured to use in the deployment.

```

    import (
    "golang.org/x/crypto/ssh"
    "k8s.io/client-go/rest"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

     "github.com/Mirantis/mcc-api/pkg/client/internalclientset"
     kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
    )

...
      kaasPublicKey := &kaasv1alpha1.PublicKey{
        ObjectMeta: metav1.ObjectMeta{
            Name: copts.KeyName,
              Namespace: cluster.Namespace,
          },
              Spec: kaasv1alpha1.PublicKeySpec{
              PublicKey: string(ssh.MarshalAuthorizedKey(ssh.publicKey)),
              },
      }

...
      // MCC API client
      clientSet :=  internalclientset.NewForConfig(rest.kubeConfig)
    
      _, err = clientSet.KaasV1alpha1().PublicKeys(publicKey.Namespace).Create(context.TODO(), publicKey, metav1.CreateOptions{})
```

__Machine objects__


```
    import (
      metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
      
      "github.com/Mirantis/mcc-api/pkg/client/internalclientset"
      clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
    )

      var createdMachine *clusterv1.Machine
      var err error
      // MCC API client
      clientSet :=  internalclientset.NewForConfig(rest.kubeConfig)

      createdMachine, err = clientSet.ClusterV1alpha1().Machines(namespace).Create(context.TODO(), machine, metav1.CreateOptions{})
```

__Wait for machine readiness__

Machine object status field Status should have value "Ready". MCC API has constant value for this:
```
    import (
        lcmv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/common/lcm/v1alpha1"
    )
    ...
    mchineStatus.Status == lcmv1alpha1.LCMMachineStateReady
```

How to get machine item object

```
    import (
      "k8s.io/client-go/rest"
      metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
      
      "github.com/Mirantis/mcc-api/pkg/client/internalclientset"
      clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
    )

    var machineslist []clusterv1.Machine
...
    cluster := &clusterv1.Cluster{
			TypeMeta: metav1.TypeMeta{
				Kind: "Cluster", // GetMachinesForCluster filters Machines by ownerReferences and checks Kind and Name fields
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      clusterName,
			},
		}
    
...
  // Define selector and add value "MachineClusterLabelName = "cluster.sigs.k8s.io/cluster-name"" 
  selector := labels.NewSelector()
  defReq, err := labels.NewRequirement(clusterv1.MachineClusterLabelName, selection.Equals, []string{cluster.Name})
  if err != nil {
    return nil, err
	}
  selector = selector.Add(*defReq)

  internalclientset.ClusterV1alpha1().Machines(cluster.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: selector.String(),
  })

```
To get each machine object item Status the next util function could be used:

```
    import (
        apisutil "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
        clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
    )

    ...
    status, err := apisutil.GetMachineStatus(apisutil.Machine)

```

The `GetMachineStatus` function is a part of MCC API and could be found  [there](github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1/util.go).

It one has an example how to decode MCC API Object Spec and Status:

```
    specObj, err := decodeExtension(machine.Spec.ProviderSpec.Value)

    // OR

    statusObj, err = decodeExtension(machine.Status.ProviderStatus)

```

Object after decoding is casted to the target type. The decodeExtention function itself:

```
func decodeExtension(ext *runtime.RawExtension) (runtime.Object, error) {
	if ext.Object != nil {
		return ext.Object, nil
	}
	s := json.NewSerializer(&json.SimpleMetaFactory{}, publicapi.Scheme, publicapi.Scheme, false)
	obj, _, err := s.Decode(ext.Raw, nil, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse RawExtension value")
	}
	ext.Object = obj
	ext.Raw = nil

	return obj, nil
}
```

All these steps are implemented in the example [part](cmd/cluster_create.go)

Details: 
- [Cluster object preparing](lib/objects/cluster.go)
- [SSH Public key](lib/ssh/sshkey.go)
- [Machine object preparing](lib/objects/cluster.go)
- Getting Object Spec and Status. Decoding RawExtention type.(github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1/util.go)


### Kubeconfig of managed cluster

The managed cluster kubeconfig could be generated from the code.

See the  [example](lib/deployer/deployer.go) with

```
func generateKubeconfig(kubeconfigOutput string, cluster *clusterv1.Cluster, realm, username, password string) error {...}
```

All these steps are implemented in the example [part](cmd/cluster_kubeconfig.go)


### Cluster upgrade

Managed cluster deployment is based on the cluster release  version, that are supported by MCC management cluster. 
Cluster upgrade process is based on the next steps:
- connect the the managed cluster:

```
    import (
        "k8s.io/client-go/rest"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

        "github.com/Mirantis/mcc-api/pkg/client/internalclientset"
        apisutil "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
        clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
    )
...

    clientSet :=  internalclientset.NewForConfig(rest.kubeConfig)

    var cluster *clusterv1.Cluster
    var err error
    cluster, err = clientSet.ClusterV1alpha1().Clusters(namespace).Get(context.TODO(), name, metav1.GetOptions{})

```

- patch cluster related  Spec.ProviderSpec.Value.Release with target release value:

```
    import (
        "k8s.io/client-go/rest"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

        "github.com/Mirantis/mcc-api/pkg/client/internalclientset"
    )

    clientSet :=  internalclientset.NewForConfig(rest.kubeConfig)

    data := []byte(fmt.Sprintf(`[{"op":"replace","path":"/spec/providerSpec/value/release","value":"%s"}]`, releaseName))

    clientSet.ClusterV1alpha1().Clusters(namespace).Patch(context.TODO(), name, types.JSONPatchType, data, metav1.PatchOptions{})
```

- wait for cluster and all its machines readiness as id described in Cluster Create section above.


All these steps with some additional useful checks are implemented in the example [part](cmd/cluster_upgrade.go)

### Managed cluster deletion

```
    import (
        "k8s.io/client-go/rest"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

        "github.com/Mirantis/mcc-api/pkg/client/internalclientset"
    )

    clientSet :=  internalclientset.NewForConfig(rest.kubeConfig)

    err := clientSet.ClusterV1alpha1().Clusters(namespace).Delete(context.TODO(), name, metav1.DeletePropagationForeground)
```

All these steps are implemented in the example [part](cmd/cluster_delete.go)

## Example client use 

### Prerequesties

#### Management cluster kubeconfig
The proper kubeconfig file to access the MCC management cluster should be present

#### OpenStack credentials:

1. Log into your OpenStack Horizon
2. Under "Project" select "API Access"
3. In dropdown to the right "Download OpenStack RC File" select "OpenStack clouds.yaml File"
4. Put resulting clouds.yaml near example binary file
<aside class="warning">
Manually written custom `clouds.yaml` file may not work as expected
</aside>
5. Edit `clouds.yaml` file and add “password” field under “clouds/openstack/auth”
   section with your OpenStack password. Final `clouds.yaml` file should look like this:

   ```yaml
   clouds:
     openstack:
       auth:
         auth_url: https://ic-eu.ssl.mirantis.net:5000/v3 # for EU cloud
         username: someusername # your username
         password: yoursecretpassword # your password - add this field
         project_id: 0123456789abcdef0123456789abcdef # your project ID
         user_domain_name: ldap-password # for LDAP users, default for service ones
       region_name: RegionOne
       interface: public
       identity_api_version: 3
   ```
#### Adjust templates to your requirements

1. In `example/yaml/cluster.yaml.template` set preferrable `dnsNameservers`.
2. In `example/yaml/machines.yaml.template` you can change flavor, image and availabilityZone.


### Binary file

Example binary file to use could be built manually in `example` folder or run the next command:
  > make  go-example-build

Generated file is /bin/example
### Cluster workflow

The next commands are implemented:
```
 ./example -h
Run example of using MCC public API

Usage:
  example [flags]
  example [command]

Available Commands:
  create      Create managed cluster
  delete      Delete managed cluster
  help        Help about any command
  kubeconfig  Get managed cluster kubeconfig
  upgrade     Upgrade managed cluster release version

```

Each command also has it own set of flags that could be checked via help command.

Sometimes cluster deploy process needs more time wait period than it hardcoded in the example. 
Using environment variable `BOOTSTRAP_INFINITE_TIMEOUT=true` we can overwrite values. This option should be used only under manual operator's attention.

#### Managed cluster creating

Creating of a managed cluster could be performed via command 'create'. The next flags could be set:
- `--management-kubeconfig` - path to the management kubeconfig, by default it is a *kubeconfig*;
- `--cluster-name` - name of a managed cluster, by default *kaas-child*;
- `--namespace` - target namespace, by default *kaas-child*;
- `--os-cloud` - cloud name in clouds.yaml;
- `--region` - openstack region, by default *region-one*;
- `--cluster` - path to the managed cluster template. It out could be found under *example/yaml/cluster.yaml*. Or you can provide yours;
- `--machines` - path to the managed machines template. It out could be found under *example/yaml/machines.yaml*. Currently the machine list should contain 3 control plane machines and 2 workers machines. This is minimal machine count requirements. Machine type controlplane is set by using *label*: `cluster.sigs.k8s.io/control-plane: "true"`;
- `--release-name` - name of the valid cluster releases inbound the management cluster. Using management kubeconfig the list of accessible cluster releases could be found: `kubectl get clusterreleases`. E.g. `mke-7-5-0-3-4-6`
- `--private-key-path` - path to ssh private key, by default *ssh_key*. The key should have PEM format

Managed cluster creation process could take more than 30 minutes, so for this operation use BOOTSTRAP_INFINITE_TIMEOUT variable set to *true*.

Result command:

```
 ./example create --management-kubeconfig kubeconfig --namespace example --cluster-name some-example --os-cloud openstack  --credentials-name cloud-config --region region-one --cluster ../example/yaml/cluster.yaml --machines ../example/yaml/machines.yaml --release-name mke-7-5-0-3-4-6 --private-key-path ssh_key
 ```


#### Managed cluster kubeconfig

Managed cluster kubeconfig is used for deploy user workloads. 
Use command `kubeconfig` 

The next flags could be set:
- `--cluster-name` - name of a managed cluster, by default *kaas-child*;
- `--namespace` - target namespace, by default *kaas-child*;
- `--kubeconfig-output` - kubeconfig output file name
- `--management-kubeconfig` - path to management cluster kubeconfig (default "kubeconfig")
- `--realm` - keycloak realm for getting auth token, by default *iam*;
- `--password` - password for getting auth token, by default *writer*. Or other user for MCC management cluster with the `writer` permissions;
- `--username` - username for getting auth token, by default *password*;

Result command:

```
./example kubeconfig --management-kubeconfig kubeconfig --namespace example --cluster-name some-example --realm iam --username writer --password password
```


 #### Upgrade cluster

Managed cluster creation is based on cluster releases that available for MCC management cluster. 
Current cluster release could be checked from current managed cluster object using management cluster kubeconfig.

`
kubectl -n example get cluster some-example -o json | jq '.spec.providerSpec.value.release'
`:

`"mke-7-6-0-rc-3-4-7"`

and from cluster status section *releaseRefs.current.name*:

`kubectl -n example get cluster some-example -o json | jq '.status.providerStatus.releaseRefs'`:

```
{
  "available": [
    {
      "name": "mke-11-0-0-rc-3-5-1",
      "version": "11.0.0-rc"
    }
  ],
  "current": {
    "allowedNodeLabels": [
      {
        "displayName": "Stacklight",
        "key": "stacklight",
        "value": "enabled"
      }
    ],
    "lcmType": "ucp",
    "name": "mke-7-6-0-rc-3-4-7",
    "version": "7.6.0-rc+3.4.7"
  },
  "previous": {
    "allowedNodeLabels": [
      {
        "displayName": "Stacklight",
        "key": "stacklight",
        "value": "enabled"
      }
    ],
    "lcmType": "ucp",
    "name": "mke-7-5-0-3-4-6",
    "version": "7.5.0+3.4.6"
  }
}
```
As we can see *available* release has value of release that could be used for managed cluster upgrade.

Updating a managed cluster release could be performed via the `upgrade` command. 
For this operation use BOOTSTRAP_INFINITE_TIMEOUT variable set to *true* to prevent timout wait fails.

Flags:
- `--cluster-name` - name of a managed cluster, by default *kaas-child*;
- `--namespace` - target namespace, by default *kaas-child*;
- `--management-kubeconfig` - path to management cluster kubeconfig (default "kubeconfig")
- `--release-name` - the name of cluster release to upgrade managed cluster

Result command:

```
./example upgrade --management-kubeconfig kubeconfig --namespace example --cluster-name some-example  --release-name mke-7-6-0-rc-3-4-7
```

 #### Delete managed cluster

 Deleting a managed cluster could be performed using the next command.

 ```
 ./example delete --management-kubeconfig kubeconfig --namespace example --cluster-name some-example
 ```

