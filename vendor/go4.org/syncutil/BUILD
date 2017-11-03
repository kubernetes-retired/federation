load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "gate.go",
        "group.go",
        "once.go",
        "sem.go",
        "syncutil.go",
    ],
    importpath = "go4.org/syncutil",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["once_test.go"],
    importpath = "go4.org/syncutil",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["sem_test.go"],
    importpath = "go4.org/syncutil_test",
    deps = [":go_default_library"],
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
        "//vendor/go4.org/syncutil/singleflight:all-srcs",
        "//vendor/go4.org/syncutil/syncdebug:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
