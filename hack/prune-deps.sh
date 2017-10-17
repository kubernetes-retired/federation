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

# Sourced from https://github.com/openshift/client-go/blob/master/hack/prune-deps.sh

# Staging paths are vendored
rm -rf vendor/k8s.io/kubernetes/staging

dep init
dep prune

# we shouldn't have modified anything
git diff-index --name-only --diff-filter=M HEAD | xargs -r git checkout -f
# we need to preserve code-generator and friends that aren't referenced in the code
git diff-index --name-only HEAD | grep -F \
  -e 'github.com/jteeuwen/go-bindata' \
  -e 'github.com/onsi/ginkgo/ginkgo' \
  -e 'k8s.io/gengo' \
  -e 'k8s.io/kube-openapi' \
  -e 'k8s.io/kubernetes/cluster' \
  -e 'k8s.io/kubernetes/cmd/gen' \
  -e 'k8s.io/kubernetes/examples' \
  -e 'k8s.io/kubernetes/hack' \
  -e 'k8s.io/kubernetes/pkg/util/template' \
  -e 'k8s.io/kubernetes/test/e2e/testing-manifests' \
  -e 'k8s.io/kubernetes/test/fixtures' \
  -e 'k8s.io/kubernetes/test/images' \
  -e 'k8s.io/kubernetes/translations' \
  -e 'vendor/k8s.io/apimachinery/pkg/util/sets/types' \
  -e 'vendor/k8s.io/code-generator' \
  -e 'vendor/k8s.io/client-go/util/cert/testdata' \
  | grep -v 'vendor/github.com/jteeuwen/go-bindata/testdata' \
  | xargs -r git checkout -f

# now cleanup what's dangling
git clean -x  -f -d
