load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "init.go",
            "kvm.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/stage1/init",
    visibility = ["//visibility:private"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/appc/goaci/proj2aci:go_default_library",
            "//vendor/github.com/appc/spec/schema:go_default_library",
            "//vendor/github.com/appc/spec/schema/types:go_default_library",
            "//vendor/github.com/coreos/go-systemd/util:go_default_library",
            "//vendor/github.com/coreos/pkg/dlopen:go_default_library",
            "//vendor/github.com/coreos/rkt/common:go_default_library",
            "//vendor/github.com/coreos/rkt/common/cgroup:go_default_library",
            "//vendor/github.com/coreos/rkt/common/cgroup/v1:go_default_library",
            "//vendor/github.com/coreos/rkt/common/cgroup/v2:go_default_library",
            "//vendor/github.com/coreos/rkt/common/networking:go_default_library",
            "//vendor/github.com/coreos/rkt/networking:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/flag:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/fs:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/mountinfo:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
            "//vendor/github.com/coreos/rkt/stage1/common:go_default_library",
            "//vendor/github.com/coreos/rkt/stage1/common/types:go_default_library",
            "//vendor/github.com/coreos/rkt/stage1/init/common:go_default_library",
            "//vendor/github.com/coreos/rkt/stage1/init/kvm:go_default_library",
            "//vendor/github.com/coreos/rkt/stage1/init/kvm/hypervisor/hvlkvm:go_default_library",
            "//vendor/github.com/coreos/rkt/stage1/init/kvm/hypervisor/hvqemu:go_default_library",
            "//vendor/github.com/godbus/dbus:go_default_library",
            "//vendor/github.com/godbus/dbus/introspect:go_default_library",
            "//vendor/github.com/hashicorp/errwrap:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_binary(
    name = "init",
    importpath = "github.com/coreos/rkt/stage1/init",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
        "//vendor/github.com/coreos/rkt/stage1/init/common:all-srcs",
        "//vendor/github.com/coreos/rkt/stage1/init/kvm:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
