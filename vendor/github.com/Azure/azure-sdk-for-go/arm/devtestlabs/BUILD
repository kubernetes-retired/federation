load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "armtemplates.go",
        "artifacts.go",
        "artifactsources.go",
        "client.go",
        "costs.go",
        "customimages.go",
        "disks.go",
        "environments.go",
        "formulas.go",
        "galleryimages.go",
        "globalschedules.go",
        "labs.go",
        "models.go",
        "notificationchannels.go",
        "policies.go",
        "policysets.go",
        "schedules.go",
        "secrets.go",
        "servicerunners.go",
        "users.go",
        "version.go",
        "virtualmachines.go",
        "virtualmachineschedules.go",
        "virtualnetworks.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/arm/devtestlabs",
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
