load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "labels.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "windows.go",
            "windows_store.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/drivers/windows",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/Microsoft/hcsshim:go_default_library",
            "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
            "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
            "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
            "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
            "//vendor/github.com/docker/libnetwork/types:go_default_library",
            "//vendor/github.com/sirupsen/logrus:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "windows_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/drivers/windows",
    library = ":go_default_library",
    deps = select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
            "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
            "//vendor/github.com/docker/libnetwork/types:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
        "//vendor/github.com/docker/libnetwork/drivers/windows/overlay:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
