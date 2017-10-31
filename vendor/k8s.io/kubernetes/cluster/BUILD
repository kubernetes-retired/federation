package(default_visibility = ["//visibility:public"])

load("@io_bazel//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/k8s.io/kubernetes/cluster/addons:all-srcs",
        "//vendor/k8s.io/kubernetes/cluster/gce:all-srcs",
        "//vendor/k8s.io/kubernetes/cluster/images/etcd-version-monitor:all-srcs",
        "//vendor/k8s.io/kubernetes/cluster/images/etcd/attachlease:all-srcs",
        "//vendor/k8s.io/kubernetes/cluster/images/etcd/rollback:all-srcs",
        "//vendor/k8s.io/kubernetes/cluster/images/hyperkube:all-srcs",
        "//vendor/k8s.io/kubernetes/cluster/images/kubemark:all-srcs",
        "//vendor/k8s.io/kubernetes/cluster/lib:all-srcs",
        "//vendor/k8s.io/kubernetes/cluster/saltbase:all-srcs",
    ],
    tags = ["automanaged"],
)

# All of the manifests that are expected to be in a "gci-trusty"
# subdir of the manifests tarball.
pkg_tar(
    name = "_manifests-gci-trusty",
    package_dir = "gci-trusty",
    visibility = ["//visibility:private"],
    deps = [
        "//cluster/addons",
        "//cluster/gce:gci-trusty-manifests",
        "//cluster/saltbase:gci-trusty-salt-manifests",
    ],
)

pkg_tar(
    name = "manifests",
    mode = "0644",
    package_dir = "kubernetes",
    deps = [
        ":_manifests-gci-trusty",
        "//cluster/saltbase:salt-manifests",
    ],
)

# These tests just verify that bash can interpret the file.
sh_test(
    name = "common_test",
    srcs = ["common.sh"],
    deps = [
        "//cluster/lib",
        "//hack/lib",
    ],
)

sh_test(
    name = "clientbin_test",
    srcs = ["clientbin.sh"],
    deps = [
        "//cluster/lib",
        "//hack/lib",
    ],
)

sh_test(
    name = "kube-util_test",
    srcs = ["kube-util.sh"],
    deps = [
        "//cluster/lib",
        "//hack/lib",
    ],
)
