package(default_visibility = ["//visibility:public"])

load("@io_kubernetes_build//defs:build.bzl", "release_filegroup")
load("@io_kubernetes_build//defs:pkg.bzl", "pkg_tar")

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
    srcs = select({
        ":package_src": ["//:all-srcs"],
        "//conditions:default": ["README-src.txt"],
    }),
    extension = "tar.gz",
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
    srcs = ["//build:client-targets"],
    mode = "0755",
    package_dir = "client/bin",
    visibility = ["//visibility:private"],
)

pkg_tar(
    name = "federation-client-%s" % PLATFORM_ARCH_STRING,
    srcs = [":license-targets"],
    extension = "tar.gz",
    mode = "0644",
    package_dir = "federation",
    deps = [
        ":_client-bin",
    ],
)

pkg_tar(
    name = "_server-bin",
    srcs = [
        "//build:client-targets",
        "//build:docker-artifacts",
        "//build:server-targets",
    ],
    mode = "0755",
    package_dir = "server/bin",
    visibility = ["//visibility:private"],
)

pkg_tar(
    name = "_server_image-bin",
    srcs = [
        "//build:docker-artifacts",
    ],
    mode = "0755",
    package_dir = "server/bin/image",
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
    srcs = [":license-targets"],
    extension = "tar.gz",
    mode = "0644",
    package_dir = "federation",
    deps = [
        ":_server-bin",
        ":_server_image-bin",
    ],
)

pkg_tar(
    name = "_test-bin",
    srcs = ["//build:test-targets"],
    mode = "0755",
    package_dir = "platforms/" + PLATFORM_ARCH_STRING.replace("-", "/"),
    # TODO: how to make this multiplatform?
    visibility = ["//visibility:private"],
)

pkg_tar(
    name = "federation-test",
    srcs = ["//build:test-portable-targets"],
    extension = "tar.gz",
    package_dir = "federation",
    strip_prefix = "//",
    deps = [
        # TODO: how to make this multiplatform?
        ":_test-bin",
    ],
)

# TODO (irfan): do we need to include server binaries here?
# We can also consider this as one final tar and include client bins too.
pkg_tar(
    name = "federation",
    srcs = [
        "//:LICENSE",
        "//:README.md",
        "//:k8s-version",
        "//:version",
        "//cluster:all-srcs",
        "//deploy:all-srcs",
        "//docs:all-srcs",
        "//hack/lib:all-srcs",
    ],
    extension = "tar.gz",
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
