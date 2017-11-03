# Cluster Federation

Kubernetes Cluster Federation enables users to federate multiple
Kubernetes clusters.
To know more details about the same please see the
[user guide](https://kubernetes.io/docs/concepts/cluster-administration/federation/).

# Deploying Kubernetes Cluster Federation

The prescribed mechanism to deploy Kubernetes Cluster Federation is using
[kubefed](https://kubernetes.io/docs/admin/kubefed/).
A complete guide for the same is available at
[setup cluster federation using kubefed](https://kubernetes.io/docs/tutorials/federation/set-up-cluster-federation-kubefed/).

# Building Kubernetes Cluster Federation

Building cluster federation binaries, which include fcp (short for federation
control plane) and [kubefed](https://kubernetes.io/docs/admin/kubefed/) 
should be as simple as running:
```shell
make
```

A kubernetes like release packages can also be built using:
```shell
make quick-release
```

The '`kubefed` binary can be found in `federation-client-*.tar.gz`.
The `fcp` binary, which self contains `federation-apiserver` and
`federation-controller-manager` can be found in `federation-server-*.tar.gz`.
`federation-server-*.tar.gz` includes `fcp-*.tar`, which is the fcp docker 
image in the tar format and can be consumed by the `kubefed` tool.


# A note to the reader
Kubernetes federation code is in a state of flux. Since its incubation, it 
lived in [core kubernetes repo](https://github.com/kubernetes/kubernetes).
The same now is maturing into [its own placeholder](https://github.com/kubernetes/federation).
The process of this movement is not yet complete. It already borrows a lot 
of code from its earlier parent, especially build infrastructure and utility 
scripts. This will be cleaned up and simplified. Subsequently we will also 
concentrate our efforts into cleaning issues and problems reported on existing 
features, with a focus of moving atleast a subset of all federation features 
towards GA.
Please raise an issue, in case you find problems and developers are most 
welcome to participate in this effort.

