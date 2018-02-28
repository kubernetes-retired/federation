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
