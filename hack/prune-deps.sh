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

# Inspired by https://github.com/openshift/client-go/blob/master/hack/prune-deps.sh

# The staging areas will be vendored as repos
rm -rf vendor/k8s.io/kubernetes/staging
# The vendored kube root build file will break federation's bazel build
rm -f vendor/k8s.io/kubernetes/BUILD.bazel
# Symlinks to missing files will break hack/verify-bazel.sh
rm -f vendor/k8s.io/kubernetes/.bazelrc
rm -f vendor/k8s.io/kubernetes/.kazelcfg.json
rm -f vendor/k8s.io/kubernetes/Makefile
rm -f vendor/k8s.io/kubernetes/Makefile.generated
rm -rf vendor/k8s.io/kubernetes/WORKSPACE

glide-vc --use-lock-file

# we shouldn't have modified anything
git diff-index --name-only --diff-filter=M HEAD | xargs -r git checkout -f
# we need to preserve code that is not referenced in the code
git diff-index --name-only HEAD | grep -F \
  -e 'BUILD' \
  -e 'LICENSE' \
  -e 'github.com/jteeuwen/go-bindata' \
  -e 'github.com/onsi/ginkgo/ginkgo' \
  -e 'k8s.io/.*/fake' \
  -e 'k8s.io/apiserver/pkg/storage/etcd/testing' \
  -e 'k8s.io/gengo' \
  -e 'k8s.io/kube-openapi' \
  -e 'k8s.io/kubernetes/cluster' \
  -e 'k8s.io/kubernetes/cmd/clicheck' \
  -e 'k8s.io/kubernetes/cmd/gen' \
  -e 'k8s.io/kubernetes/examples' \
  -e 'k8s.io/kubernetes/hack' \
  -e 'k8s.io/kubernetes/pkg/api/testing' \
  -e 'k8s.io/kubernetes/pkg/kubectl/cmd/testing' \
  -e 'k8s.io/kubernetes/pkg/registry/registrytest' \
  -e 'k8s.io/kubernetes/pkg/util/template' \
  -e 'k8s.io/kubernetes/test/e2e/testing-manifests' \
  -e 'k8s.io/kubernetes/test/fixtures' \
  -e 'k8s.io/kubernetes/test/images' \
  -e 'k8s.io/kubernetes/translations' \
  -e 'vendor/k8s.io/apimachinery/pkg/util/sets/types' \
  -e 'vendor/k8s.io/client-go/util/cert/testdata' \
  -e 'vendor/k8s.io/code-generator' \
  | grep -v 'vendor/github.com/jteeuwen/go-bindata/testdata' \
  | grep -v 'vendor/k8s.io/kubernetes/staging' \
  | grep -v 'vendor/k8s.io/kubernetes/BUILD.bazel' \
  | grep -v 'vendor/k8s.io/kubernetes/.bazelrc' \
  | grep -v 'vendor/k8s.io/kubernetes/.kazelcfg.json' \
  | grep -v 'vendor/k8s.io/kubernetes/Makefile' \
  | grep -v 'vendor/k8s.io/kubernetes/Makefile.generated' \
  | grep -v 'vendor/k8s.io/kubernetes/WORKSPACE' \
  | xargs -r git checkout -f

# now cleanup what's dangling
git clean -x  -f -d
