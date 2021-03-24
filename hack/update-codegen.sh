#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ${GOPATH}/src/k8s.io/code-generator)}
bash $CODEGEN_PKG/generate-groups.sh "deepcopy,client,informer,lister" \
  volcano.sh/apis/pkg/client volcano.sh/apis/pkg/apis \
  batch:v1alpha1 bus:v1alpha1 scheduling:v1beta1 \
  --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt
