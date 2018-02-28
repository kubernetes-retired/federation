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
        "//vendor/google.golang.org/genproto/googleapis/api/annotations:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/configchange:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/distribution:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/httpbody:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/label:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/metric:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/monitoredres:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/serviceconfig:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/servicecontrol/v1:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/servicemanagement/v1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
