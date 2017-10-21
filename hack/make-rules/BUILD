# Scripts runnable from make, e.g.
#
#    cd $GOPATH/src/k8s.io/federation
#    make test
#
# The sh_binary rules below exist only to validate
# dependencies; if a shell dependency is accidentally
# deleted, a presubmit BUILD will fail.
#
# If the scripts sourced their dependencies from
# $RUNFILES (rather than $BASH_SOURCE/../.. or
# whatever), then bazel build hack/...  would install
# runnable, hermetically sealed shell "binaries".
# E.g. the following command would work:
#
#   ./bazel-bin/hack/make-rules/test
#

sh_binary(
    name = "test-integration",
    srcs = ["test-integration.sh"],
    deps = [
        "//hack/lib",
    ],
)

sh_binary(
    name = "build",
    srcs = ["build.sh"],
    deps = [
        "//hack/lib",
    ],
)

sh_binary(
    name = "cross",
    srcs = ["cross.sh"],
    deps = [
        "//hack/lib",
    ],
)

sh_binary(
    name = "test",
    srcs = ["test.sh"],
    deps = [
        "//hack/lib",
    ],
)

sh_binary(
    name = "vet",
    srcs = ["vet.sh"],
    deps = [
        "//hack/lib",
    ],
)

sh_binary(
    name = "verify",
    srcs = ["verify.sh"],
    deps = [
        "//hack/lib",
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
