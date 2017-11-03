load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "key.go",
            "syscall_windows.go",
            "win32.go",
            "zsyscall_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/shiny/driver/internal/win32",
    visibility = ["//vendor/golang.org/x/exp/shiny/driver:__subpackages__"],
    deps = select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/golang.org/x/exp/shiny/screen:go_default_library",
            "//vendor/golang.org/x/mobile/event/key:go_default_library",
            "//vendor/golang.org/x/mobile/event/lifecycle:go_default_library",
            "//vendor/golang.org/x/mobile/event/mouse:go_default_library",
            "//vendor/golang.org/x/mobile/event/paint:go_default_library",
            "//vendor/golang.org/x/mobile/event/size:go_default_library",
            "//vendor/golang.org/x/mobile/geom:go_default_library",
            "//vendor/golang.org/x/sys/windows:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
