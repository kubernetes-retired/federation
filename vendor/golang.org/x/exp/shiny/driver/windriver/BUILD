load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "other.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "buffer.go",
            "screen.go",
            "syscall_windows.go",
            "texture.go",
            "window.go",
            "windraw.go",
            "windriver.go",
            "zsyscall_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/shiny/driver/windriver",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/exp/shiny/driver/internal/errscreen:go_default_library",
        "//vendor/golang.org/x/exp/shiny/screen:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/golang.org/x/exp/shiny/driver/internal/drawer:go_default_library",
            "//vendor/golang.org/x/exp/shiny/driver/internal/event:go_default_library",
            "//vendor/golang.org/x/exp/shiny/driver/internal/swizzle:go_default_library",
            "//vendor/golang.org/x/exp/shiny/driver/internal/win32:go_default_library",
            "//vendor/golang.org/x/image/math/f64:go_default_library",
            "//vendor/golang.org/x/mobile/event/key:go_default_library",
            "//vendor/golang.org/x/mobile/event/lifecycle:go_default_library",
            "//vendor/golang.org/x/mobile/event/mouse:go_default_library",
            "//vendor/golang.org/x/mobile/event/paint:go_default_library",
            "//vendor/golang.org/x/mobile/event/size:go_default_library",
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
