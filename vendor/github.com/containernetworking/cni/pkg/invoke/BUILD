load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "args.go",
        "delegate.go",
        "exec.go",
        "find.go",
        "raw_exec.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "os_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "os_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "os_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/containernetworking/cni/pkg/invoke",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "exec_test.go",
        "find_test.go",
        "get_version_integration_test.go",
        "invoke_suite_test.go",
        "raw_exec_test.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/invoke_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/invoke/fakes:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types/current:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version/testhelpers:go_default_library",
        "//vendor/github.com/containernetworking/cni/plugins/test/noop/debug:go_default_library",
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/ginkgo/extensions/table:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
        "//vendor/github.com/onsi/gomega/gexec:go_default_library",
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
