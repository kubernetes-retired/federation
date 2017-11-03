#!/bin/bash

# Copyright 2014 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

kube::golang::setup_env

BUILD_TARGETS=(
  vendor/k8s.io/code-generator/cmd/client-gen
)
make -C "${KUBE_ROOT}" WHAT="${BUILD_TARGETS[*]}"

clientgen=$(kube::util::find-binary "client-gen")

# Please do not add any logic to this shell script. Add logic to the go code
# that generates the set-gen program.
#

# This can be called with one flag, --verify-only, so it works for both the
# update- and verify- scripts.
${clientgen} --clientset-name=federation_clientset --clientset-path=k8s.io/federation/client/clientset_generated --input-base="k8s.io/federation/vendor/k8s.io/api" --input="../../../apis/federation/v1beta1","core/v1","extensions/v1beta1","batch/v1","autoscaling/v1" --included-types-overrides="core/v1/Service,core/v1/Namespace,extensions/v1beta1/ReplicaSet,core/v1/Secret,extensions/v1beta1/Ingress,extensions/v1beta1/Deployment,extensions/v1beta1/DaemonSet,core/v1/ConfigMap,core/v1/Event,batch/v1/Job,autoscaling/v1/HorizontalPodAutoscaler" --go-header-file="${KUBE_ROOT}/hack/boilerplate/boilerplate.go.txt" "$@"
