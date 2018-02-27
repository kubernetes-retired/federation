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
        "//vendor/github.com/emicklei/go-restful/examples/google_app_engine/datastore:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
