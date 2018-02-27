load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:windows": [
            "advapi32.go",
            "combobox.go",
            "comctl32.go",
            "comdlg32.go",
            "datetimepicker.go",
            "edit.go",
            "gdi32.go",
            "gdiplus.go",
            "header.go",
            "kernel32.go",
            "listbox.go",
            "listview.go",
            "menu.go",
            "ole32.go",
            "oleaut32.go",
            "opengl32.go",
            "pdh.go",
            "shdocvw.go",
            "shell32.go",
            "shobj.go",
            "statusbar.go",
            "syslink.go",
            "tab.go",
            "toolbar.go",
            "tooltip.go",
            "treeview.go",
            "updown.go",
            "user32.go",
            "uxtheme.go",
            "win.go",
            "winspool.go",
        ],
        "//conditions:default": [],
    }) + select({
        "@io_bazel_rules_go//go/platform:windows_386": [
            "oleaut32_386.go",
            "shobj_386.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "oleaut32_amd64.go",
            "shobj_amd64.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/lxn/win",
    visibility = ["//visibility:public"],
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
