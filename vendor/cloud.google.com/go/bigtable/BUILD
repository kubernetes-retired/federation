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
        "//vendor/cloud.google.com/go/bigtable/bttest:all-srcs",
        "//vendor/cloud.google.com/go/bigtable/cmd/cbt:all-srcs",
        "//vendor/cloud.google.com/go/bigtable/cmd/emulator:all-srcs",
        "//vendor/cloud.google.com/go/bigtable/cmd/loadtest:all-srcs",
        "//vendor/cloud.google.com/go/bigtable/cmd/scantest:all-srcs",
        "//vendor/cloud.google.com/go/bigtable/internal/cbtrc:all-srcs",
        "//vendor/cloud.google.com/go/bigtable/internal/gax:all-srcs",
        "//vendor/cloud.google.com/go/bigtable/internal/option:all-srcs",
        "//vendor/cloud.google.com/go/bigtable/internal/stat:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
