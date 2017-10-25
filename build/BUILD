package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_docker//docker:docker.bzl", "docker_image", "docker_bundle")
load("@io_kubernetes_build//defs:build.bzl", "release_filegroup")

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//build/release-tars:all-srcs",
        "//build/visible_to:all-srcs",
    ],
    tags = ["automanaged"],
)

# Change the base image to hyperkube to reduce size, if need be.
docker_image(
    name = "fcp-image",
    base = "@debian-hyperkube-base-amd64//image",
    cmd = ["/fcp"],
    files = ["//cmd/fcp"],
)

docker_bundle(
    name = "fcp-amd64",
    images = {
        "gcr.io/google_containers/fcp-amd64:{STABLE_DOCKER_TAG}": "fcp-image",
    },
    stamp = True,
)

release_filegroup(
    name = "docker-artifacts",
    srcs = [":fcp-amd64.tar"],
)

genrule(
    name = "os_package_version",
    outs = ["version"],
    cmd = """
grep ^STABLE_BUILD_SCM_REVISION bazel-out/stable-status.txt \
    | awk '{print $$2}' \
    | sed -e 's/^v//' -Ee 's/-([a-z]+)/~\\1/' -e 's/-/+/g' \
    >$@
""",
    stamp = 1,
)

# KUBE_CLIENT_TARGETS
release_filegroup(
    name = "client-targets",
    srcs = [
        "//cmd/kubefed",
    ],
)

# KUBE_SERVER_TARGETS
# No need to duplicate CLIENT_TARGETS or NODE_TARGETS here,
# since we include them in the actual build rule.
release_filegroup(
    name = "server-targets",
    srcs = [
        "//cmd/fcp",
    ],
)

# kube::golang::test_targets
filegroup(
    name = "test-targets",
    srcs = [
        "//cmd/genfeddocs",
        "//test/e2e:e2e.test",
        "//vendor/github.com/onsi/ginkgo/ginkgo",
        #        "//vendor/k8s.io/kubernetes/cmd/gendocs",
        #        "//vendor/k8s.io/kubernetes/cmd/genman",
        #        "//vendor/k8s.io/kubernetes/cmd/genswaggertypedocs",
        #        "//vendor/k8s.io/kubernetes/cmd/genyaml",
        #        "//vendor/k8s.io/kubernetes/cmd/linkcheck",
    ],
)

# KUBE_TEST_PORTABLE
filegroup(
    name = "test-portable-targets",
    srcs = [
        "//develop:all-srcs",
        "//hack:e2e.go",
        "//hack:federated-ginkgo-e2e.sh",
        "//hack:get-build.sh",
        "//hack:ginkgo-e2e.sh",
        "//hack/e2e-internal:all-srcs",
        "//hack/lib:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/testing-manifests:all-srcs",
    ],
)
