package(default_visibility = ["//visibility:public"])

genrule(
    name = "gen_e2e_node.test",
    testonly = 1,
    srcs = [":go_default_test"],
    outs = ["e2e_node.test"],
    cmd = "srcs=($(SRCS)); cp $$(dirname $${srcs[0]})/go_default_test $@;",
    output_to_bindir = 1,
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
        "//vendor/k8s.io/kubernetes/test/e2e_node/builder:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/environment:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/perftype:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/remote:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/runner/local:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/runner/remote:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/services:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/system:all-srcs",
    ],
    tags = ["automanaged"],
)
