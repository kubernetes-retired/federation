package(default_visibility = ["//visibility:public"])

load("@io_kubernetes_build//defs:pkg.bzl", "pkg_tar")

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
        "//cluster/gce:all-srcs",
        "//cluster/images/etcd-version-monitor:all-srcs",
        "//cluster/images/etcd/attachlease:all-srcs",
        "//cluster/images/etcd/rollback:all-srcs",
        "//cluster/lib:all-srcs",
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
        "//cluster/gce:gci-trusty-manifests",
    ],
)

pkg_tar(
    name = "manifests",
    mode = "0644",
    package_dir = "kubernetes",
    deps = [
        ":_manifests-gci-trusty",
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
