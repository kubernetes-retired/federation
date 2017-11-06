#!/usr/bin/env bash

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

# Download a Federation release.
# Usage:
#   wget -q -O - https://storage.googleapis.com/kubernetes-federation-release | bash
# or
#   curl -fsSL https://storage.googleapis.com/kubernetes-federation-release | bash
#
# Advanced options
#  Set FEDERATION_RELEASE to choose a specific release instead of the current
#    stable release, (e.g. 'v1.3.7').
#    See https://github.com/kubernetes/federation/releases for release options.
#  Set FEDERATION_RELEASE_URL to choose where to download binaries from.
#    (Defaults to https://storage.googleapis.com/kubernetes-federation-release/release).
#
#  Set FEDERATION_SKIP_CONFIRM to skip the installation confirmation prompt.
#  Set FEDERATION_SKIP_RELEASE_VALIDATION to skip trying to validate the
#      federation release string. This implies that you know what you're doing
#      and have set FEDERATION_RELEASE and FEDERATION_RELEASE_URL properly.

set -o errexit
set -o nounset
set -o pipefail

# If FEDERATION_RELEASE_URL is overridden but FEDERATION_CI_RELEASE_URL is not then set FEDERATION_CI_RELEASE_URL to FEDERATION_RELEASE_URL.
FEDERATION_CI_RELEASE_URL="${FEDERATION_CI_RELEASE_URL:-${FEDERATION_RELEASE_URL:-https://storage.googleapis.com/kubernetes-federation-release/ci}}"
FEDERATION_RELEASE_URL="${FEDERATION_RELEASE_URL:-https://storage.googleapis.com/kubernetes-federation-release}"

FEDERATION_RELEASE_VERSION_REGEX="^v(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)(-([a-zA-Z0-9]+)\\.(0|[1-9][0-9]*))?$"
FEDERATION_CI_VERSION_REGEX="^v(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)-([a-zA-Z0-9]+)\\.(0|[1-9][0-9]*)(\\.(0|[1-9][0-9]*)\\+[-0-9a-z]*)?$"

# Sets FEDERATION_VERSION variable if an explicit version number was provided (e.g. "v1.0.6",
# "v1.2.0-alpha.1.881+376438b69c7612") or resolves the "published" version
# <path>/<version> (e.g. "release/stable",' "ci/latest-1") by reading from GCS.
#
# See the docs on getting builds for more information about version
# publication.
#
# Args:
#   $1 version string from command line
# Vars set:
#   FEDERATION_VERSION
function set_binary_version() {
  if [[ "${1}" =~ "/" ]]; then
    export FEDERATION_VERSION=$(curl -fsSL --retry 5 "https://storage.googleapis.com/kubernetes-federation-release/${1}.txt")
  else
    export FEDERATION_VERSION=${1}
  fi
}

# Use the script from inside the Federation tarball to fetch the client and
# server binaries (if not included in federation.tar.gz).
function download_federation_binaries {
  (
    cd federation
    if [[ -x ./cluster/get-federation-binaries.sh ]]; then
      # Make sure to use the same download URL in get-federation-binaries.sh
      FEDERATION_RELEASE_URL="${FEDERATION_RELEASE_URL}" \
        ./cluster/get-federation-binaries.sh
    fi
  )
}

file=federation.tar.gz
release=${FEDERATION_RELEASE:-"release/stable"}

# Validate Federation release version.
# Translate a published version <bucket>/<version> (e.g. "release/stable") to version number.
set_binary_version "${release}"
if [[ -z "${FEDERATION_SKIP_RELEASE_VALIDATION-}" ]]; then
  if [[ ${FEDERATION_VERSION} =~ ${FEDERATION_CI_VERSION_REGEX} ]]; then
    # Override FEDERATION_RELEASE_URL to point to the CI bucket;
    # this will be used by get-federation-binaries.sh.
    FEDERATION_RELEASE_URL="${FEDERATION_CI_RELEASE_URL}"
  elif ! [[ ${FEDERATION_VERSION} =~ ${FEDERATION_RELEASE_VERSION_REGEX} ]]; then
    echo "Version doesn't match regexp" >&2
    exit 1
  fi
fi
federation_tar_url="${FEDERATION_RELEASE_URL}/${FEDERATION_VERSION}/${file}"

need_download=true
if [[ -r "${PWD}/${file}" ]]; then
  downloaded_version=$(tar -xzOf "${PWD}/${file}" federation/version 2>/dev/null || true)
  echo "Found preexisting ${file}, release ${downloaded_version}"
  if [[ "${downloaded_version}" == "${FEDERATION_VERSION}" ]]; then
    echo "Using preexisting federation.tar.gz"
    need_download=false
  fi
fi

if "${need_download}"; then
  echo "Downloading federation release ${FEDERATION_VERSION}"
  echo "  from ${federation_tar_url}"
  echo "  to ${PWD}/${file}"
fi

if [[ -e "${PWD}/federation" ]]; then
  # Let's try not to accidentally nuke something that isn't a federation
  # release dir.
  if [[ ! -f "${PWD}/federation/version" ]]; then
    echo "${PWD}/federation exists but does not look like a Federation release."
    echo "Aborting!"
    exit 5
  fi
  echo "Will also delete preexisting 'federation' directory."
fi

if [[ -z "${FEDERATION_SKIP_CONFIRM-}" ]]; then
  echo "Is this ok? [Y]/n"
  read confirm
  if [[ "${confirm}" =~ ^[nN]$ ]]; then
    echo "Aborting."
    exit 0
  fi
fi

if "${need_download}"; then
  if [[ $(which curl) ]]; then
    curl -fL --retry 5 --keepalive-time 2 "${federation_tar_url}" -o "${file}"
  elif [[ $(which wget) ]]; then
    wget "${federation_tar_url}"
  else
    echo "Couldn't find curl or wget.  Bailing out."
    exit 1
  fi
fi

echo "Unpacking federation release ${FEDERATION_VERSION}"
rm -rf "${PWD}/federation"
tar -xzf ${file}

download_federation_binaries
