package(default_visibility = ["//visibility:public"])

licenses(["notice"])

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
        "//vendor/k8s.io/kubernetes/third_party/forked/etcd221/pkg/fileutil:all-srcs",
        "//vendor/k8s.io/kubernetes/third_party/forked/etcd221/wal:all-srcs",
        "//vendor/k8s.io/kubernetes/third_party/forked/etcd237/pkg/fileutil:all-srcs",
        "//vendor/k8s.io/kubernetes/third_party/forked/etcd237/wal:all-srcs",
        "//vendor/k8s.io/kubernetes/third_party/forked/golang/expansion:all-srcs",
        "//vendor/k8s.io/kubernetes/third_party/forked/golang/reflect:all-srcs",
        "//vendor/k8s.io/kubernetes/third_party/forked/golang/template:all-srcs",
        "//vendor/k8s.io/kubernetes/third_party/forked/gonum/graph:all-srcs",
        "//vendor/k8s.io/kubernetes/third_party/htpasswd:all-srcs",
    ],
    tags = ["automanaged"],
)
