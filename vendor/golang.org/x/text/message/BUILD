load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "catalog.go",
        "format.go",
        "message.go",
        "print.go",
    ],
    importpath = "golang.org/x/text/message",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/text/internal/number:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/message/catalog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "fmt_test.go",
        "message_test.go",
    ],
    importpath = "golang.org/x/text/message",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/text/internal:go_default_library",
        "//vendor/golang.org/x/text/internal/format:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/message/catalog:go_default_library",
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
        "//vendor/golang.org/x/text/message/catalog:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
