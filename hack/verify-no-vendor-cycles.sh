#!/bin/bash

# Copyright 2017 The Kubernetes Authors.
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

failed=false
for i in $(find vendor/ -type d); do
  deps=$(go list -f '{{range .Deps}}{{.}}{{"\n"}}{{end}}' ./$i 2> /dev/null || echo "")
  deps_on_main=$(echo "${deps}" | grep -v "k8s.io/federation/vendor/" | grep "k8s.io/federation" || echo "")
  if [ -n "${deps_on_main}" ]; then
    echo "Package ${i} has a cyclic dependency on the main repository."
    failed=true
  fi
done

if [[ "${failed}" == "true" ]]; then
  exit 1
fi
