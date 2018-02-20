load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "acidentifier.go",
        "ackind.go",
        "acname.go",
        "annotations.go",
        "app.go",
        "date.go",
        "dependencies.go",
        "doc.go",
        "environment.go",
        "errors.go",
        "event_handler.go",
        "exec.go",
        "hash.go",
        "isolator.go",
        "isolator_linux_specific.go",
        "isolator_resources.go",
        "isolator_unix.go",
        "labels.go",
        "mountpoint.go",
        "port.go",
        "semver.go",
        "url.go",
        "user_annotations.go",
        "user_labels.go",
        "uuid.go",
        "volume.go",
    ],
    importpath = "github.com/appc/spec/schema/types",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema/common:go_default_library",
        "//vendor/github.com/appc/spec/schema/types/resource:go_default_library",
        "//vendor/github.com/coreos/go-semver/semver:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "acidentifier_test.go",
        "ackind_test.go",
        "acname_test.go",
        "annotations_test.go",
        "app_test.go",
        "date_test.go",
        "dependencies_test.go",
        "environment_test.go",
        "exec_test.go",
        "hash_test.go",
        "isolator_linux_specific_test.go",
        "isolator_resources_test.go",
        "isolator_test.go",
        "isolator_unix_test.go",
        "labels_test.go",
        "mountpoint_test.go",
        "port_test.go",
        "semver_test.go",
        "url_test.go",
        "uuid_test.go",
        "volume_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/appc/spec/schema/types",
    deps = [
        "//vendor/github.com/appc/spec/schema/types/resource:go_default_library",
        "//vendor/github.com/coreos/go-semver/semver:go_default_library",
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
        "//vendor/github.com/appc/spec/schema/types/resource:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
