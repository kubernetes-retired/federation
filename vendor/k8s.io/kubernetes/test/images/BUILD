package(default_visibility = ["//visibility:public"])

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
        "//vendor/k8s.io/kubernetes/test/images/clusterapi-tester:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/entrypoint-tester:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/fakegitserver:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/goproxy:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/liveness:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/logs-generator:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/mounttest:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/n-way-http:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/net:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/netexec:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/nettest:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/no-snat-test:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/no-snat-test-proxy:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/nonewprivs:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/pets/peer-finder:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/port-forward-tester:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/porter:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/resource-consumer:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/serve-hostname:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images/test-webserver:all-srcs",
    ],
    tags = ["automanaged"],
)
