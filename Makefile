
CRD_OPTIONS ?= "crd:crdVersions=v1,generateEmbeddedObjectMeta=true"

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.6.0 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif


# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) paths="./pkg/apis/scheduling/v1beta1;./pkg/apis/batch/v1alpha1;./pkg/apis/bus/v1alpha1;./pkg/apis/nodeinfo/v1alpha1/;./pkg/apis/flow/v1alpha1" output:crd:artifacts:config=manifests/crd/bases
	$(CONTROLLER_GEN) "crd:crdVersions=v1beta1" paths="./pkg/apis/scheduling/v1beta1;./pkg/apis/batch/v1alpha1;./pkg/apis/bus/v1alpha1;./pkg/apis/flow/v1alpha1;./pkg/apis/nodeinfo/v1alpha1" output:crd:artifacts:config=manifests/crd/v1beta1