load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "authorizationpolicies.go",
        "client.go",
        "connectormappings.go",
        "connectors.go",
        "hubs.go",
        "images.go",
        "interactions.go",
        "kpi.go",
        "links.go",
        "models.go",
        "profiles.go",
        "relationshiplinks.go",
        "relationships.go",
        "roleassignments.go",
        "roles.go",
        "version.go",
        "views.go",
        "widgettypes.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/arm/customer-insights",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/date:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/validation:go_default_library",
        "//vendor/github.com/shopspring/decimal:go_default_library",
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
