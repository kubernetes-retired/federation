load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "blockopt.go",
        "builder.go",
        "const.go",
        "create.go",
        "doc.go",
        "dom.go",
        "emit.go",
        "func.go",
        "lift.go",
        "lvalue.go",
        "methods.go",
        "mode.go",
        "print.go",
        "sanity.go",
        "source.go",
        "ssa.go",
        "testmain.go",
        "util.go",
        "wrappers.go",
    ],
    importpath = "golang.org/x/tools/go/ssa",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/go/ast/astutil:go_default_library",
        "//vendor/golang.org/x/tools/go/types/typeutil:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "builder_test.go",
        "example_test.go",
        "source_test.go",
        "stdlib_test.go",
        "testmain_test.go",
    ],
    data = glob(["testdata/**"]),
    importpath = "golang.org/x/tools/go/ssa_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/tools/go/ast/astutil:go_default_library",
        "//vendor/golang.org/x/tools/go/buildutil:go_default_library",
        "//vendor/golang.org/x/tools/go/loader:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa/ssautil:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/golang.org/x/tools/go/ssa/interp:all-srcs",
        "//vendor/golang.org/x/tools/go/ssa/ssautil:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
