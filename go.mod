module volcano.sh/apis

go 1.13

require (
	k8s.io/api v0.19.6
	k8s.io/apimachinery v0.19.6
	k8s.io/apiserver v0.19.6
	k8s.io/client-go v0.19.6
	k8s.io/klog v1.0.0
	k8s.io/code-generator v0.20.4
)

replace (
	k8s.io/api => k8s.io/api v0.19.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.6
	k8s.io/apiserver => k8s.io/apiserver v0.19.6
	k8s.io/client-go => k8s.io/client-go v0.19.6
	k8s.io/klog => k8s.io/klog v1.0.0
)
