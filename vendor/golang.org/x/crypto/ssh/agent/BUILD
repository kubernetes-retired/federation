load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "forward.go",
        "keyring.go",
        "server.go",
    ],
    importpath = "golang.org/x/crypto/ssh/agent",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/crypto/ed25519:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "client_test.go",
        "keyring_test.go",
        "server_test.go",
        "testdata_test.go",
    ],
    importpath = "golang.org/x/crypto/ssh/agent",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
        "//vendor/golang.org/x/crypto/ssh/testdata:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "golang.org/x/crypto/ssh/agent_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
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
