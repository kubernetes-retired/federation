load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "kvm.go",
        "net_plugin.go",
        "networking.go",
        "podenv.go",
        "portfwd.go",
    ],
    importpath = "github.com/coreos/rkt/networking",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/ip:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/ns:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/utils:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/utils/sysctl:go_default_library",
        "//vendor/github.com/coreos/go-iptables/iptables:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/common/networking:go_default_library",
        "//vendor/github.com/coreos/rkt/networking/netinfo:go_default_library",
        "//vendor/github.com/coreos/rkt/networking/tuntap:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
        "//vendor/github.com/vishvananda/netlink:go_default_library",
        "//vendor/golang.org/x/sys/unix:go_default_library",
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
        "//vendor/github.com/coreos/rkt/networking/netinfo:all-srcs",
        "//vendor/github.com/coreos/rkt/networking/tuntap:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
