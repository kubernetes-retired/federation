#!/bin/bash

set -e

sr() (
  perl -p -i -e "s|$1|$2|g" `grep -ril $1 *`
)

pushd federation > /dev/null
  sr 'kubernetes/federation' federation
  sr 'kubernetes\.federation' federation
popd
git mv federation/Makefile Makefile.federation
git mv federation/BUILD BUILD
git mv federation/test .
git mv federation/cluster test/
git mv federation/* .
rmdir federation

git add .
