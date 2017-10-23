package(default_visibility = ["//visibility:public"])

load("@io_bazel//tools/build_defs/pkg:pkg.bzl", "pkg_tar")
load("@io_kubernetes_build//defs:build.bzl", "release_filegroup")

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
)

config_setting(
    name = "package_src",
    values = {
        "define": "PACKAGE_SRC=true",
    },
    visibility = ["//visibility:private"],
)

genrule(
    name = "federation-src-readme",
    outs = ["README-src.txt"],
    cmd = """
echo For build efficiency, the src was not included in this release.>$@
echo The full source code can be viewed at >>$@
echo -n https://github.com/kubernetes/federation/tree/ >>$@
grep ^STABLE_BUILD_GIT_COMMIT bazel-out/stable-status.txt | cut -d' ' -f2 >>$@
""",
    stamp = 1,
)

pkg_tar(
    name = "federation-src",
    extension = "tar.gz",
    files = select({
        ":package_src": ["//:all-srcs"],
        "//conditions:default": ["README-src.txt"],
    }),
    package_dir = "federation",
    strip_prefix = select({
        ":package_src": "//",
        "//conditions:default": ".",
    }),
)

# FIXME: this should be configurable/auto-detected
PLATFORM_ARCH_STRING = "linux-amd64"

# Included in server tarballs.
filegroup(
    name = "license-targets",
    srcs = [
        ":federation-src.tar.gz",
        "//:LICENSE",
    ],
    visibility = ["//visibility:private"],
)

pkg_tar(
    name = "_client-bin",
    files = ["//build:client-targets"],
    mode = "0755",
    package_dir = "client/bin",
    visibility = ["//visibility:private"],
)

pkg_tar(
    name = "federation-client-%s" % PLATFORM_ARCH_STRING,
    extension = "tar.gz",
    package_dir = "federation",
    deps = [
        ":_client-bin",
    ],
)

pkg_tar(
    name = "_server-bin",
    files = [
        "//build:client-targets",
        "//build:server-targets",
    ],
    mode = "0755",
    package_dir = "server/bin",
    visibility = ["//visibility:private"],
)

genrule(
    name = "dummy",
    outs = [".dummy"],
    cmd = "touch $@",
    visibility = ["//visibility:private"],
)

pkg_tar(
    name = "federation-server-%s" % PLATFORM_ARCH_STRING,
    extension = "tar.gz",
    files = [":license-targets"],
    mode = "0644",
    package_dir = "federation",
    deps = [
        ":_server-bin",
    ],
)

pkg_tar(
    name = "_test-bin",
    files = ["//build:test-targets"],
    mode = "0755",
    package_dir = "platforms/" + PLATFORM_ARCH_STRING.replace("-", "/"),
    # TODO: how to make this multiplatform?
    visibility = ["//visibility:private"],
)

pkg_tar(
    name = "federation-test",
    extension = "tar.gz",
    files = ["//build:test-portable-targets"],
    package_dir = "federation",
    strip_prefix = "//",
    deps = [
        # TODO: how to make this multiplatform?
        ":_test-bin",
    ],
)

pkg_tar(
    name = "federation",
    extension = "tar.gz",
    files = [
        "//:LICENSE",
        "//:README.md",
        "//:version",
        "//cluster:all-srcs",
        "//develop:all-srcs",
        "//deploy:all-srcs",
        "//docs:all-srcs",
        "//hack/lib:all-srcs",
    ],
    package_dir = "federation",
    strip_prefix = "//",
    deps = [
        ":_server-bin",
    ],
)

release_filegroup(
    name = "release-tars",
    srcs = [
        ":federation.tar.gz",
        ":federation-client-%s.tar.gz" % PLATFORM_ARCH_STRING,
        ":federation-server-%s.tar.gz" % PLATFORM_ARCH_STRING,
        ":federation-src.tar.gz",
        ":federation-test.tar.gz",
    ],
)
