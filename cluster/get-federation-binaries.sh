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

# This script downloads and installs the Federation client and server
# (and optionally test) binaries,
# It is intended to be called from an extracted Federation release tarball.
#
# We automatically choose the correct client binaries to download.
#
# Options:
#  Set FEDERATION_SERVER_ARCH to choose the server
#  architecture to download:
#    * amd64 [default]
#    * arm
#    * arm64
#    * ppc64le
#    * s390x
#
#  Set FEDERATION_SKIP_CONFIRM to skip the installation confirmation prompt.
#  Set FEDERATION_RELEASE_URL to choose where to download binaries from.
#    (Defaults to https://storage.googleapis.com/kubernetes-federation-release/release).
#  Set FEDERATION_DOWNLOAD_TESTS to additionally download and extract the test
#    binaries tarball.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd)

FEDERATION_RELEASE_URL="${FEDERATION_RELEASE_URL:-https://storage.googleapis.com/kubernetes-federation-release}"

function detect_federation_release() {
  if [[ -n "${FEDERATION_VERSION:-}" ]]; then
    return 0  # Allow caller to explicitly set version
  fi

  if [[ ! -e "${KUBE_ROOT}/version" ]]; then
    echo "Can't determine Federation release." >&2
    echo "${BASH_SOURCE} should only be run from a prebuilt Federation release." >&2
    echo "Did you mean to use get-federation.sh instead?" >&2
    exit 1
  fi

  FEDERATION_VERSION=$(cat "${KUBE_ROOT}/version")
}

function detect_client_info() {
  local kernel=$(uname -s)
  case "${kernel}" in
    Darwin)
      CLIENT_PLATFORM="darwin"
      ;;
    Linux)
      CLIENT_PLATFORM="linux"
      ;;
    *)
      echo "Unknown, unsupported platform: ${kernel}." >&2
      echo "Supported platforms: Linux, Darwin." >&2
      echo "Bailing out." >&2
      exit 2
  esac

  # TODO: migrate the kube::util::host_platform function out of hack/lib and
  # use it here.
  local machine=$(uname -m)
  case "${machine}" in
    x86_64*|i?86_64*|amd64*)
      CLIENT_ARCH="amd64"
      ;;
    aarch64*|arm64*)
      CLIENT_ARCH="arm64"
      ;;
    arm*)
      CLIENT_ARCH="arm"
      ;;
    i?86*)
      CLIENT_ARCH="386"
      ;;
    s390x*)
      CLIENT_ARCH="s390x"
      ;;
    *)
      echo "Unknown, unsupported architecture (${machine})." >&2
      echo "Supported architectures x86_64, i686, arm, arm64, s390x." >&2
      echo "Bailing out." >&2
      exit 3
      ;;
  esac
}

function md5sum_file() {
  if which md5 >/dev/null 2>&1; then
    md5 -q "$1"
  else
    md5sum "$1" | awk '{ print $1 }'
  fi
}

function sha1sum_file() {
  if which sha1sum >/dev/null 2>&1; then
    sha1sum "$1" | awk '{ print $1 }'
  else
    shasum -a1 "$1" | awk '{ print $1 }'
  fi
}

function download_tarball() {
  local -r download_path="$1"
  local -r file="$2"
  url="${DOWNLOAD_URL_PREFIX}/${file}"
  mkdir -p "${download_path}"
  if [[ $(which curl) ]]; then
    curl -fL --retry 3 --keepalive-time 2 "${url}" -o "${download_path}/${file}"
  elif [[ $(which wget) ]]; then
    wget "${url}" -O "${download_path}/${file}"
  else
    echo "Couldn't find curl or wget.  Bailing out." >&2
    exit 4
  fi
  echo
  local md5sum=$(md5sum_file "${download_path}/${file}")
  echo "md5sum(${file})=${md5sum}"
  local sha1sum=$(sha1sum_file "${download_path}/${file}")
  echo "sha1sum(${file})=${sha1sum}"
  echo
  # TODO: add actual verification
}

function extract_arch_tarball() {
  local -r tarfile="$1"
  local -r platform="$2"
  local -r arch="$3"

  platforms_dir="${KUBE_ROOT}/platforms/${platform}/${arch}"
  echo "Extracting ${tarfile} into ${platforms_dir}"
  mkdir -p "${platforms_dir}"
  # Tarball looks like federation/{client,server}/bin/BINARY"
  tar -xzf "${tarfile}" --strip-components 3 -C "${platforms_dir}"
  # Create convenience symlink
  ln -sf "${platforms_dir}" "$(dirname ${tarfile})/bin"
  echo "Add '$(dirname ${tarfile})/bin' to your PATH to use newly-installed binaries."
}

detect_federation_release
DOWNLOAD_URL_PREFIX="${FEDERATION_RELEASE_URL}/${FEDERATION_VERSION}"

SERVER_PLATFORM="linux"
SERVER_ARCH="${FEDERATION_SERVER_ARCH:-amd64}"
SERVER_TAR="federation-server-${SERVER_PLATFORM}-${SERVER_ARCH}.tar.gz"

detect_client_info
CLIENT_TAR="federation-client-${CLIENT_PLATFORM}-${CLIENT_ARCH}.tar.gz"

echo "Federation release: ${FEDERATION_VERSION}"
echo "Server: ${SERVER_PLATFORM}/${SERVER_ARCH}  (to override, set FEDERATION_SERVER_ARCH)"
echo "Client: ${CLIENT_PLATFORM}/${CLIENT_ARCH}  (autodetected)"
echo

echo "Will download ${SERVER_TAR} from ${DOWNLOAD_URL_PREFIX}"
echo "Will download and extract ${CLIENT_TAR} from ${DOWNLOAD_URL_PREFIX}"

TESTS_TAR="federation-test.tar.gz"
DOWNLOAD_TESTS_TAR=false
if [[ -n "${FEDERATION_DOWNLOAD_TESTS-}" ]]; then
  DOWNLOAD_TESTS_TAR=true
  echo "Will download and extract ${TESTS_TAR} from ${DOWNLOAD_URL_PREFIX}"
fi

if [[ -z "${FEDERATION_SKIP_CONFIRM-}" ]]; then
  echo "Is this ok? [Y]/n"
  read confirm
  if [[ "${confirm}" =~ ^[nN]$ ]]; then
    echo "Aborting."
    exit 1
  fi
fi

download_tarball "${KUBE_ROOT}/server" "${SERVER_TAR}"

download_tarball "${KUBE_ROOT}/client" "${CLIENT_TAR}"
extract_arch_tarball "${KUBE_ROOT}/client/${CLIENT_TAR}" "${CLIENT_PLATFORM}" "${CLIENT_ARCH}"

if "${DOWNLOAD_TESTS_TAR}"; then
  download_tarball "${KUBE_ROOT}/test" "${TESTS_TAR}"
  echo "Extracting ${TESTS_TAR} into ${KUBE_ROOT}"
  # Strip leading "federation/"
  tar -xzf "${KUBE_ROOT}/test/${TESTS_TAR}" --strip-components 1 -C "${KUBE_ROOT}"
fi
