load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "demo",
    srcs = ["code.go"],
    importpath = "bdl.com/demo",
    visibility = ["//visibility:public"],
)

go_test(
    name = "demo_test",
    srcs = ["code_test.go"],
    embed = [":demo"],
)
