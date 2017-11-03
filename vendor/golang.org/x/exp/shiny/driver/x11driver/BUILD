load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "buffer.go",
        "screen.go",
        "shm_other.go",
        "texture.go",
        "window.go",
        "x11driver.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "shm_linux_amd64.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/shiny/driver/x11driver",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/BurntSushi/xgb:go_default_library",
        "//vendor/github.com/BurntSushi/xgb/render:go_default_library",
        "//vendor/github.com/BurntSushi/xgb/shm:go_default_library",
        "//vendor/github.com/BurntSushi/xgb/xproto:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/drawer:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/errscreen:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/event:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/lifecycler:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/swizzle:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/x11key:go_default_library",
        "//vendor/golang.org/x/exp/shiny/screen:go_default_library",
        "//vendor/golang.org/x/image/math/f64:go_default_library",
        "//vendor/golang.org/x/mobile/event/key:go_default_library",
        "//vendor/golang.org/x/mobile/event/mouse:go_default_library",
        "//vendor/golang.org/x/mobile/event/paint:go_default_library",
        "//vendor/golang.org/x/mobile/event/size:go_default_library",
        "//vendor/golang.org/x/mobile/geom:go_default_library",
    ],
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
