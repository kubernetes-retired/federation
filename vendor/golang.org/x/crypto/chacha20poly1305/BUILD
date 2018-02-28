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
        "//vendor/golang.org/x/crypto/chacha20poly1305/internal/chacha20:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
