# mcc-api

This is a Mirantis Container Cloud (MCC) public API. Contains a set of types and clientset.

Use the public API for operates managed cluster over exsisting MCC management cluster, e.g. create, delete, etc.

Could be used for managed clusters across multiple cloud provider platforms, both on premises and in the cloud. 
Please, see the [Mirantis Container Cloud product overview](https://docs.mirantis.com/container-cloud/latest/overview.html) for details.
A client implementation example could be found in the [example](/example) package.

### Getting started

MCC API could be useful to implement own client for managed cluster operation, such as creating, upgrading, etc. 
The [README.md](example/README.md) describes a provided client example of using API. 

### What's included

* The `pkg/apis/common` package contains the set of Mirantis projects related types are used in public API.
* The  `pkg/apis/public` package contains all public API related types.
* The `pkg/apis/util` package contains the public API related utils, e.g. getting ProviderSpec or ProviderStatus methods, etc.
* The `pkg/client` package contains the a MCC Public API clientset, scheme and fake client to use in tests.
* The `pkg/errors` package contains the errors wrappers.
* The `pkg/util` package contains the different functions could be useful.
* The `example` package contains the implemented client for a managed cluster operation (create, upgrade and delete) for the Openstack cloud provider. 


