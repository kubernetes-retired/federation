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

# This script builds ui assets into a single go datafile

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

cd "${KUBE_ROOT}"

if ! which go-bindata > /dev/null 2>&1 ; then
  echo "Cannot find go-bindata. Install with \"go get github.com/jteeuwen/go-bindata/...\""
  exit 1
fi

readonly TMP_DATAFILE="/tmp/datafile.go"
readonly SWAGGER_SRC="third_party/swagger-ui/..."
readonly SWAGGER_PKG="swagger"

function kube::hack::build_ui() {
  local pkg="$1"
  local src="$2"
  local output_file="pkg/genericapiserver/server/routes/data/${pkg}/datafile.go"

  go-bindata -nocompress -o "${output_file}" -prefix ${PWD} -pkg "${pkg}" "${src}"

  local year=$(date +%Y)
  cat hack/boilerplate/boilerplate.go.txt | sed "s/YEAR/${year}/" > "${TMP_DATAFILE}"
  echo -e "// generated by hack/build-ui.sh; DO NOT EDIT\n" >> "${TMP_DATAFILE}"
  cat "${output_file}" >> "${TMP_DATAFILE}"

  gofmt -s -w "${TMP_DATAFILE}"

  mv "${TMP_DATAFILE}" "${output_file}"
}

kube::hack::build_ui "${SWAGGER_PKG}" "${SWAGGER_SRC}"
