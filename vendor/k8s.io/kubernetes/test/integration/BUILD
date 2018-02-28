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
        "//vendor/k8s.io/kubernetes/test/integration/apiserver:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/auth:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/client:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/configmap:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/defaulttolerationseconds:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/deployment:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/etcd:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/evictions:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/examples:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/framework:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/garbagecollector:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/master:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/metrics:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/objectmeta:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/openshift:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/pods:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/quota:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/replicaset:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/replicationcontroller:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/scheduler:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/scheduler_perf:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/secrets:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/serviceaccount:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/storageclasses:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/ttlcontroller:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/volume:all-srcs",
    ],
    tags = ["automanaged"],
)
