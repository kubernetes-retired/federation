## `Federation transition from k/k`

# Build and test

- Clone this repo to a go path
  - `mkdir -p fcp/src/k8s.io`
  - `cd fcp/src/k8s.io`
  - `git clone https://github.com/marun/federation`
  - `cd federation`
  - `git checkout -t origin/fed-move-out`
- Build
  - make
- Test the binaries
  - ./_output/bin/kubefed
  - ./_output/bin/fcp
  - ./_output/bin/e2e.test
- Run unit tess
  - make test
- Run integration tests
  - make test-integration

# Updating dependencies

- Install glide
  - https://github.com/Masterminds/glide
- Updated vendored dependencies
  - `glide up -v`
    - `-v` ensures the removal of dependencies `vendored` path

# Notes

 - The fcp entrypoint copies supporting HyperKube source from k/k.  It
   may make sense to move that infrastructure from cmd to pkg in k/k
   so it can be imported rather than copied.
 - k/k is vendored from the [vendorable branch of marun/kubernetes](https://github.com/marun/kubernetes/tree/vendorable).
   - The vendorable branch is based on the [remove-federation branch
     of
     marun/kubernetes](https://github.com/marun/kubernetes/tree/remove-federation).
   - All federation code has been removed from the `remove-federation`
     branch, and this represenst the state of k/k that will be
     vendored by k/f.
   - The last commit - pre-removal of federation - of the `vendorable`
     branch is the one referred to by `Kubernetes-commit` in HEAD of
     the staging repos
     (e.g. [kubernetes/apimachinery](https://github.com/kubernetes/apimachinery/commit/e9a29eff7d472df0f7da9ead5ab59b74e74a07ec)).
     This is the easiest way to ensure that the versions of k/k and
     its staging dependencies are compatible.  When `glide up -v`
     brings in new commits from the staging repos such that the
     vendored k/k gets behind, it will be a simple matter to discover
     the new k/k commit to update the vendorable branch to.
   - Once the [removal
     PR](https://github.com/kubernetes/kubernetes/pull/53816) lands,
     it should be possible for k/f to pin to a version of
     k8s.io/kubernetes rather than relying on a fork.

## Next steps

 - Get bazel build and test working
