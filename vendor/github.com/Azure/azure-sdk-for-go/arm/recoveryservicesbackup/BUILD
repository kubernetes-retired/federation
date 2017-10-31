load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "backupengines.go",
        "backupjobs.go",
        "backupoperationresults.go",
        "backupoperationstatuses.go",
        "backuppolicies.go",
        "backupprotectableitems.go",
        "backupprotecteditems.go",
        "backupprotectioncontainers.go",
        "backupresourcestorageconfigs.go",
        "backupresourcevaultconfigs.go",
        "backups.go",
        "backupusagesummaries.go",
        "client.go",
        "exportjobsoperationresults.go",
        "itemlevelrecoveryconnections.go",
        "jobcancellations.go",
        "jobdetails.go",
        "joboperationresults.go",
        "jobs.go",
        "models.go",
        "operations.go",
        "protecteditemoperationresults.go",
        "protecteditemoperationstatuses.go",
        "protecteditems.go",
        "protectioncontaineroperationresults.go",
        "protectioncontainerrefreshoperationresults.go",
        "protectioncontainers.go",
        "protectionpolicies.go",
        "protectionpolicyoperationresults.go",
        "protectionpolicyoperationstatuses.go",
        "recoverypoints.go",
        "restores.go",
        "securitypins.go",
        "version.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/arm/recoveryservicesbackup",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/date:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
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
