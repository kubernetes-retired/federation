package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "k8s.io/federation/pkg/federation-controller",
)

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
        "//pkg/federation-controller/cluster:all-srcs",
        "//pkg/federation-controller/ingress:all-srcs",
        "//pkg/federation-controller/job:all-srcs",
        "//pkg/federation-controller/service:all-srcs",
        "//pkg/federation-controller/sync:all-srcs",
        "//pkg/federation-controller/util:all-srcs",
    ],
    tags = ["automanaged"],
)
