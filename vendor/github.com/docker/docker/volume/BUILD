load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "validate.go",
        "validate_test_unix.go",
        "volume.go",
        "volume_copy.go",
        "volume_copy_unix.go",
        "volume_propagation_unsupported.go",
        "volume_unsupported.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "volume_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "volume_linux.go",
            "volume_propagation_linux.go",
            "volume_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "validate_test_windows.go",
            "volume_copy_windows.go",
            "volume_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/volume",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types/mount:go_default_library",
        "//vendor/github.com/docker/docker/pkg/idtools:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringid:go_default_library",
        "//vendor/github.com/opencontainers/selinux/go-selinux/label:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "validate_test.go",
        "volume_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "volume_linux_test.go",
            "volume_propagation_linux_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/volume",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/docker/api/types/mount:go_default_library"],
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
        "//vendor/github.com/docker/docker/volume/drivers:all-srcs",
        "//vendor/github.com/docker/docker/volume/local:all-srcs",
        "//vendor/github.com/docker/docker/volume/store:all-srcs",
        "//vendor/github.com/docker/docker/volume/testutils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
