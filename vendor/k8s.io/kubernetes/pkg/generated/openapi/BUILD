# doc.go is managed by kazel, so gazelle should ignore it.
# gazelle:exclude doc.go

package(default_visibility = ["//visibility:public"])

load("//vendor/k8s.io/kubernetes/pkg/generated/openapi:def.bzl", "openapi_library")

openapi_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    openapi_targets = [
        "vendor/k8s.io/api/admission/v1alpha1",
        "vendor/k8s.io/api/admissionregistration/v1alpha1",
        "vendor/k8s.io/api/apps/v1",
        "vendor/k8s.io/api/apps/v1beta1",
        "vendor/k8s.io/api/apps/v1beta2",
        "vendor/k8s.io/api/authentication/v1",
        "vendor/k8s.io/api/authentication/v1beta1",
        "vendor/k8s.io/api/authorization/v1",
        "vendor/k8s.io/api/authorization/v1beta1",
        "vendor/k8s.io/api/autoscaling/v1",
        "vendor/k8s.io/api/autoscaling/v2beta1",
        "vendor/k8s.io/api/batch/v1",
        "vendor/k8s.io/api/batch/v1beta1",
        "vendor/k8s.io/api/batch/v2alpha1",
        "vendor/k8s.io/api/certificates/v1beta1",
        "vendor/k8s.io/api/core/v1",
        "vendor/k8s.io/api/extensions/v1beta1",
        "vendor/k8s.io/api/imagepolicy/v1alpha1",
        "vendor/k8s.io/api/networking/v1",
        "vendor/k8s.io/api/policy/v1beta1",
        "vendor/k8s.io/api/rbac/v1",
        "vendor/k8s.io/api/rbac/v1alpha1",
        "vendor/k8s.io/api/rbac/v1beta1",
        "vendor/k8s.io/api/scheduling/v1alpha1",
        "vendor/k8s.io/api/settings/v1alpha1",
        "vendor/k8s.io/api/storage/v1",
        "vendor/k8s.io/api/storage/v1beta1",
        "vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1",
        "vendor/k8s.io/apimachinery/pkg/api/resource",
        "vendor/k8s.io/apimachinery/pkg/apis/meta/v1",
        "vendor/k8s.io/apimachinery/pkg/apis/meta/v1alpha1",
        "vendor/k8s.io/apimachinery/pkg/runtime",
        "vendor/k8s.io/apimachinery/pkg/util/intstr",
        "vendor/k8s.io/apimachinery/pkg/version",
        "vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1",
        "vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1",
        "vendor/k8s.io/client-go/pkg/version",
        "vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1",
        "vendor/k8s.io/kubernetes/pkg/apis/abac/v0",
        "vendor/k8s.io/kubernetes/pkg/apis/abac/v1beta1",
        "vendor/k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1",
        "vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1",
        "vendor/k8s.io/kubernetes/pkg/version",
        "vendor/k8s.io/metrics/pkg/apis/metrics/v1alpha1",
        "vendor/k8s.io/metrics/pkg/apis/metrics/v1beta1",
    ],
    tags = ["automanaged"],
    vendor_targets = [],
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
)
