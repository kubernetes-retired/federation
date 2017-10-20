# Package groups defined for use in kubernetes visibility rules.
#
# See associated README.md for explanation.
#
# Style suggestions:
#
#  - Sort package group definitions by name.
#
#  - Prefer obvious package group names.
#
#    E.g "pkg_kubectl_cmd_util_CONSUMERS" names a group
#    of packages allowed to depend on (consume) the
#    //pkg/kubectl/cmd/util package.
#
#
#  - A group name ending in _BAD wants to be deleted.
#
#    Such a group wants to contract, rather than expand.
#    It likely exists to permit a legacy unintentional
#    dependency that requires more work to remove.
#
#  - Prefer defining new groups to expanding groups.
#
#    The former permits tight targeting, the latter can
#    allow unnecessary visibility and thus bad deps.
#

package_group(
    name = "COMMON_generators",
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/gendocs",
        "//vendor/k8s.io/kubernetes/cmd/genman",
        "//vendor/k8s.io/kubernetes/cmd/genyaml",
    ],
)

package_group(
    name = "COMMON_testing",
    packages = [
        "//hack",
        "//hack/lib",
        "//hack/make-rules",
        "//vendor/k8s.io/kubernetes/test/e2e",
        "//vendor/k8s.io/kubernetes/test/e2e/framework",
        "//vendor/k8s.io/kubernetes/test/e2e/kubectl",
        "//vendor/k8s.io/kubernetes/test/e2e/workload",
        "//vendor/k8s.io/kubernetes/test/integration/etcd",
        "//vendor/k8s.io/kubernetes/test/integration/framework",
        "//vendor/k8s.io/kubernetes/test/integration/kubectl",
    ],
)

package_group(
    name = "FEDERATION_BAD",
    packages = [
        "//cmd/genfeddocs",
        "//cmd/kubefed/app",
        "//pkg/kubefed",
        "//pkg/kubefed/init",
        "//pkg/kubefed/testing",
        "//pkg/kubefed/util",
    ],
)

package_group(
    name = "cluster",
    packages = [
        "//cluster/...",
        "//vendor/k8s.io/kubernetes/cluster",
        ],
)

package_group(
    name = "KUBEADM_BAD",
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/cmd",
    ],
)

package_group(
    name = "cmd_kubectl_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/cmd",
    ],
)

package_group(
    name = "cmd_kubectl_app_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/kubectl",
    ],
)

package_group(
    name = "pkg_kubectl_CONSUMERS_BAD",
    includes = [
        ":FEDERATION_BAD",
        ":KUBEADM_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/clicheck",
        "//vendor/k8s.io/kubernetes/cmd/hyperkube",
        "//vendor/k8s.io/kubernetes/pkg",
    ],
)

package_group(
    name = "pkg_kubectl_CONSUMERS",
    includes = [
        ":COMMON_generators",
        ":pkg_kubectl_CONSUMERS_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/kubectl",
        "//vendor/k8s.io/kubernetes/cmd/kubectl/app",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/auth",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/config",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/rollout",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/set",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/testing",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util/editor",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_CONSUMERS_BAD",
    includes = [
        ":FEDERATION_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/clicheck",
        "//vendor/k8s.io/kubernetes/cmd/hyperkube",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_CONSUMERS",
    includes = [
        ":COMMON_generators",
        ":pkg_kubectl_cmd_CONSUMERS_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/kubectl",
        "//vendor/k8s.io/kubernetes/cmd/kubectl/app",
        "//vendor/k8s.io/kubernetes/pkg/kubectl",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_auth_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/rollout",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_config_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_rollout_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_set_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/rollout",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_templates_CONSUMERS_BAD",
    packages = [
        "//pkg/kubefed/init",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_templates_CONSUMERS",
    includes = [
        ":COMMON_generators",
        ":COMMON_testing",
        ":FEDERATION_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/kubectl",
        "//vendor/k8s.io/kubernetes/cmd/kubectl/app",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/auth",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/config",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/rollout",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/set",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/templates",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util/sanity",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_testdata_edit_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_testing_CONSUMERS_BAD",
    packages = [
        "//pkg/kubefed",
        "//pkg/kubefed/init",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_testing_CONSUMERS",
    includes = [
        ":pkg_kubectl_cmd_testing_CONSUMERS_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/auth",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/set",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_util_CONSUMERS_BAD",
    includes = [
        ":FEDERATION_BAD",
        ":KUBEADM_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/clicheck",
        "//vendor/k8s.io/kubernetes/cmd/hyperkube",
        "//vendor/k8s.io/kubernetes/cmd/kube-proxy/app",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_util_CONSUMERS",
    includes = [
        ":COMMON_generators",
        ":COMMON_testing",
        ":pkg_kubectl_cmd_util_CONSUMERS_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/kubectl",
        "//vendor/k8s.io/kubernetes/cmd/kubectl/app",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/auth",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/config",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/rollout",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/set",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/testing",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util/editor",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_util_editor_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_util_jsonmerge_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util",
    ],
)

package_group(
    name = "pkg_kubectl_cmd_util_sanity_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/clicheck",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util",
    ],
)

package_group(
    name = "pkg_kubectl_metricsutil_CONSUMERS_BAD",
    includes = [
        ":FEDERATION_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/clicheck",
        "//vendor/k8s.io/kubernetes/cmd/hyperkube",
    ],
)

package_group(
    name = "pkg_kubectl_metricsutil_CONSUMERS",
    includes = [
        ":COMMON_generators",
        ":pkg_kubectl_metricsutil_CONSUMERS_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/kubectl",
        "//vendor/k8s.io/kubernetes/cmd/kubectl/app",
        "//vendor/k8s.io/kubernetes/pkg/kubectl",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
    ],
)

package_group(
    name = "pkg_kubectl_resource_CONSUMERS_BAD",
    packages = [
        "//pkg/kubefed",
    ],
)

package_group(
    name = "pkg_kubectl_resource_CONSUMERS",
    includes = [
        ":COMMON_generators",
        ":COMMON_testing",
        ":pkg_kubectl_resource_CONSUMERS_BAD",
    ],
    packages = [
        "//vendor/k8s.io/kubernetes/cmd/kubectl",
        "//vendor/k8s.io/kubernetes/cmd/kubectl/app",
        "//vendor/k8s.io/kubernetes/pkg/kubectl",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/auth",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/config",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/rollout",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/set",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/testing",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util/editor",
    ],
)

package_group(
    name = "pkg_kubectl_testing_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl",
        "//vendor/k8s.io/kubernetes/pkg/printers/internalversion",
    ],
)

package_group(
    name = "pkg_kubectl_util_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/proxy",
    ],
)

package_group(
    name = "pkg_kubectl_validation_CONSUMERS",
    packages = [
        "//vendor/k8s.io/kubernetes/pkg/kubectl",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/testing",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/cmd/util",
        "//vendor/k8s.io/kubernetes/pkg/kubectl/resource",
    ],
)

# Added by ./hack/verify-bazel.sh; should be excluded from
# that script since it makes no sense here.
filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
)

# Added by ./hack/verify-bazel.sh; should be excluded from
# that script since it makes no sense here.
filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
