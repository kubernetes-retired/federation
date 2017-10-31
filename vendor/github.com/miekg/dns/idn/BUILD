load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "code_points.go",
        "punycode.go",
    ],
    importpath = "github.com/miekg/dns/idn",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/miekg/dns:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["punycode_test.go"],
    importpath = "github.com/miekg/dns/idn",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "github.com/miekg/dns/idn_test",
    deps = [":go_default_library"],
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
