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
        "//vendor/golang.org/x/crypto/openpgp/armor:all-srcs",
        "//vendor/golang.org/x/crypto/openpgp/clearsign:all-srcs",
        "//vendor/golang.org/x/crypto/openpgp/elgamal:all-srcs",
        "//vendor/golang.org/x/crypto/openpgp/errors:all-srcs",
        "//vendor/golang.org/x/crypto/openpgp/packet:all-srcs",
        "//vendor/golang.org/x/crypto/openpgp/s2k:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
