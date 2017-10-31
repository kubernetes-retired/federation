load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "account.go",
        "activity.go",
        "agentregistrationinformation.go",
        "certificate.go",
        "client.go",
        "connection.go",
        "connectiontype.go",
        "credential.go",
        "dsccompilationjob.go",
        "dscconfiguration.go",
        "dscnode.go",
        "dscnodeconfiguration.go",
        "fields.go",
        "hybridrunbookworkergroup.go",
        "job.go",
        "jobschedule.go",
        "jobstream.go",
        "models.go",
        "module.go",
        "nodereports.go",
        "objectdatatypes.go",
        "operations.go",
        "runbook.go",
        "runbookdraft.go",
        "schedule.go",
        "statistics.go",
        "testjobs.go",
        "testjobstreams.go",
        "usages.go",
        "variable.go",
        "version.go",
        "webhook.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/arm/automation",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/date:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/validation:go_default_library",
        "//vendor/github.com/satori/uuid:go_default_library",
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
