load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "apitypes.go",
        "auth.go",
        "availabilityzones.go",
        "client.go",
        "deployments.go",
        "disks.go",
        "flavors.go",
        "hosts.go",
        "images.go",
        "info.go",
        "networks.go",
        "projects.go",
        "resourcetickets.go",
        "restclient.go",
        "routers.go",
        "services.go",
        "status.go",
        "subnets.go",
        "tasks.go",
        "tenants.go",
        "util.go",
        "virtualnetworks.go",
        "vms.go",
    ],
    importpath = "github.com/vmware/photon-controller-go-sdk/photon",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/vmware/photon-controller-go-sdk/photon/lightwave:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "auth_test.go",
        "availabilityzones_test.go",
        "client_test.go",
        "deployments_test.go",
        "disks_test.go",
        "error_test.go",
        "flavors_test.go",
        "helpers_test.go",
        "hosts_test.go",
        "images_test.go",
        "info_test.go",
        "mocks_test.go",
        "networks_test.go",
        "projects_test.go",
        "resourcetickets_test.go",
        "routers_test.go",
        "services_test.go",
        "status_test.go",
        "subnets_test.go",
        "tenants_test.go",
        "util_test.go",
        "virtualnetworks_test.go",
        "vms_test.go",
    ],
    importpath = "github.com/vmware/photon-controller-go-sdk/photon",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
        "//vendor/github.com/vmware/photon-controller-go-sdk/photon/internal/mocks:go_default_library",
        "//vendor/github.com/vmware/photon-controller-go-sdk/photon/lightwave:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["photon_suite_test.go"],
    importpath = "github.com/vmware/photon-controller-go-sdk/photon_test",
    deps = [
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
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
        "//vendor/github.com/vmware/photon-controller-go-sdk/photon/lightwave:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
