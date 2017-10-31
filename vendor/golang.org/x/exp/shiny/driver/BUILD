load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "driver.go",
        "driver_fallback.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "driver_darwin.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "driver_x11.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "driver_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/shiny/driver",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/exp/shiny/driver/internal/errscreen:go_default_library",
        "//vendor/golang.org/x/exp/shiny/screen:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "//vendor/golang.org/x/exp/shiny/driver/gldriver:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/golang.org/x/exp/shiny/driver/x11driver:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/golang.org/x/exp/shiny/driver/windriver:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/gldriver:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/internal/drawer:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/internal/errscreen:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/internal/event:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/internal/lifecycler:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/internal/swizzle:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/internal/win32:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/internal/x11key:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/windriver:all-srcs",
        "//vendor/golang.org/x/exp/shiny/driver/x11driver:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
