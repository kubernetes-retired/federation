load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_kubernetes_build//defs:go.bzl", "go_genrule")
load("//build:openapi.bzl", "openapi_target_prefix", "openapi_output_prefix", "openapi_pkg_prefix")

def openapi_library(name, tags, srcs, openapi_targets=[], vendor_targets=[]):
  deps = [
      "//vendor/github.com/go-openapi/spec:go_default_library",
      "//vendor/k8s.io/kube-openapi/pkg/common:go_default_library",
  ] + ["//%s:go_default_library" % target for target in openapi_targets] + ["//vendor/%s:go_default_library" % target for target in vendor_targets]
  go_library(
      name=name,
      tags=tags,
      srcs=srcs + [":zz_generated.openapi"],
      deps=deps,
  )
  go_genrule(
      name = "zz_generated.openapi",
      srcs = srcs + ["//" + openapi_pkg_prefix + "hack/boilerplate:boilerplate.go.txt"],
      outs = ["zz_generated.openapi.go"],
      cmd = " ".join([
        "$(location //vendor/k8s.io/code-generator/cmd/openapi-gen)",
        "--v 1",
        "--logtostderr",
        "--go-header-file $(location //" + openapi_pkg_prefix + "hack/boilerplate:boilerplate.go.txt)",
        "--output-file-base zz_generated.openapi",
        "--output-package " + openapi_output_prefix + "k8s.io/kubernetes/pkg/generated/openapi",
        "--input-dirs " + ",".join([openapi_target_prefix + target for target in openapi_targets] + ["k8s.io/kubernetes/vendor/" + target for target in vendor_targets]),
        "&& cp " + openapi_pkg_prefix + "pkg/generated/openapi/zz_generated.openapi.go $(location :zz_generated.openapi.go)",
      ]),
      go_deps = deps,
      tools = ["//vendor/k8s.io/code-generator/cmd/openapi-gen"],
)
