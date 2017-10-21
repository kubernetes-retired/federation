package(default_visibility = ["//visibility:public"])

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
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/admit:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/alwayspullimages:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/antiaffinity:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/defaulttolerationseconds:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/deny:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/eventratelimit:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/exec:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/gc:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/imagepolicy:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/initialresources:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/namespace/autoprovision:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/namespace/exists:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/noderestriction:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/persistentvolume/label:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/persistentvolume/resize:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/podnodeselector:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/podpreset:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/priority:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/resourcequota:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/security:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/securitycontext/scdeny:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/storageclass/setdefault:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/admission/webhook:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/auth:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler:all-srcs",
    ],
    tags = ["automanaged"],
)
