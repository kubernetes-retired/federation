load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "agreements.go",
        "certificates.go",
        "client.go",
        "integrationaccounts.go",
        "maps.go",
        "models.go",
        "partners.go",
        "schemas.go",
        "sessions.go",
        "version.go",
        "workflowrunactions.go",
        "workflowruns.go",
        "workflows.go",
        "workflowtriggerhistories.go",
        "workflowtriggers.go",
        "workflowversions.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/arm/logic",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/date:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/validation:go_default_library",
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
