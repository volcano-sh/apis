## Introduction
volcano-sh/api aims to provide a flexible way to generate CRD informer/clientset/lister with the Kubernetes version you 
need by separating api package with main repository. 
## Getting Started
1. Clone the repository to local.
```shell
git clone https://github.com/openshift-evangelists/crd-code-generation.git
```
2. Update the Kubernetes version in go.mod to what you want. Current release is taking v0.19.6 as an example.
```go
module volcano.sh/apis

go 1.13

require (
	k8s.io/api v0.19.6
	k8s.io/apimachinery v0.19.6
	k8s.io/apiserver v0.19.6
	k8s.io/client-go v0.19.6
	k8s.io/code-generator v0.20.4 // indirect
	k8s.io/klog v1.0.0
)

replace (
	k8s.io/api => k8s.io/api v0.19.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.6
	k8s.io/apiserver => k8s.io/apiserver v0.19.6
	k8s.io/client-go => k8s.io/client-go v0.19.6
	k8s.io/klog => k8s.io/klog v1.0.0
)
```
3. Execute the script hack/update-codegen.sh and you will get the CRD clientset in github.com/volcano-sh/apis/pkg/client,
which is under $GOPATH/src
```shell
bash hack/update-codegen.sh

# expected output as follows:
Generating deepcopy funcs
Generating clientset for batch.volcano.sh:v1alpha1 at github.com/volcano-sh/apis/pkg/client/clientset
Generating listers for batch.volcano.sh:v1alpha1 at github.com/volcano-sh/apis/pkg/client/listers
Generating informers for batch.volcano.sh:v1alpha1 at github.com/volcano-sh/apis/pkg/client/informers
Generating deepcopy funcs
Generating clientset for bus.volcano.sh:v1alpha1 at github.com/volcano-sh/apis/pkg/client/clientset
Generating listers for bus.volcano.sh:v1alpha1 at github.com/volcano-sh/apis/pkg/client/listers
Generating informers for bus.volcano.sh:v1alpha1 at github.com/volcano-sh/apis/pkg/client/informers
Generating deepcopy funcs
Generating clientset for scheduling.volcano.sh:v1beta1 at github.com/volcano-sh/apis/pkg/client/clientset
Generating listers for scheduling.volcano.sh:v1beta1 at github.com/volcano-sh/apis/pkg/client/listers
Generating informers for scheduling.volcano.sh:v1beta1 at github.com/volcano-sh/apis/pkg/client/informers
```
