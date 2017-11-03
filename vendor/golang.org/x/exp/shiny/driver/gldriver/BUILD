load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "buffer.go",
        "egl.go",
        "gldriver.go",
        "other.go",
        "screen.go",
        "texture.go",
        "window.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "cocoa.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "x11.c",
            "x11.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "win32.go",
        ],
        "//conditions:default": [],
    }),
    cgo = True,
    clinkopts = select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "-framework Cocoa -framework OpenGL",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "-lEGL -lGLESv2 -lX11",
        ],
        "//conditions:default": [],
    }),
    copts = select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "-x objective-c",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/shiny/driver/gldriver",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/exp/shiny/driver/internal/drawer:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/errscreen:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/event:go_default_library",
        "//vendor/golang.org/x/exp/shiny/driver/internal/lifecycler:go_default_library",
        "//vendor/golang.org/x/exp/shiny/screen:go_default_library",
        "//vendor/golang.org/x/image/math/f64:go_default_library",
        "//vendor/golang.org/x/mobile/event/lifecycle:go_default_library",
        "//vendor/golang.org/x/mobile/event/size:go_default_library",
        "//vendor/golang.org/x/mobile/gl:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "//vendor/golang.org/x/mobile/event/key:go_default_library",
            "//vendor/golang.org/x/mobile/event/mouse:go_default_library",
            "//vendor/golang.org/x/mobile/event/paint:go_default_library",
            "//vendor/golang.org/x/mobile/geom:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/golang.org/x/exp/shiny/driver/internal/x11key:go_default_library",
            "//vendor/golang.org/x/mobile/event/key:go_default_library",
            "//vendor/golang.org/x/mobile/event/mouse:go_default_library",
            "//vendor/golang.org/x/mobile/event/paint:go_default_library",
            "//vendor/golang.org/x/mobile/geom:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/golang.org/x/exp/shiny/driver/internal/win32:go_default_library",
            "//vendor/golang.org/x/mobile/event/key:go_default_library",
            "//vendor/golang.org/x/mobile/event/mouse:go_default_library",
            "//vendor/golang.org/x/mobile/event/paint:go_default_library",
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
