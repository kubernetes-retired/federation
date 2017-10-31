load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "log.go",
        "network.go",
        "resources.go",
    ],
    importpath = "github.com/coreos/rkt/stage1/init/kvm",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/coreos/go-systemd/unit:go_default_library",
        "//vendor/github.com/coreos/rkt/networking:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "network_test.go",
        "resources_test.go",
    ],
    importpath = "github.com/coreos/rkt/stage1/init/kvm",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
    ],
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
        "//vendor/github.com/coreos/rkt/stage1/init/kvm/hypervisor:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
