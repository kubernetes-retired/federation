load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "dns_config.go",
        "log.go",
        "mount.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "app.go",
            "path.go",
            "pod.go",
            "seccomp.go",
            "seccomp_wildcards.go",
            "units.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/stage1/init/common",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/fileutil:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/fs:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/coreos/go-systemd/unit:go_default_library",
            "//vendor/github.com/coreos/rkt/common/cgroup:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/acl:go_default_library",
            "//vendor/github.com/coreos/rkt/stage1/common/types:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = [
        "dns_config_test.go",
        "mount_test.go",
        "pod_test.go",
    ],
    importpath = "github.com/coreos/rkt/stage1/init/common",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/common/types:go_default_library",
        "//vendor/github.com/kr/pretty:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
