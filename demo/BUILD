load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "code",
    srcs = ["code.go"],
    importpath = "demo",
)

go_test(
    name = "code_test",
    srcs = ["code_test.go"],
    embed = [":code"],
)


