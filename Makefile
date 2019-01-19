K8SCONTEXTTOOLS_BIN_NAME=k8stokentools

K8SCONTEXTTOOLS_VERSION=1.0.0

VERSION ?=1.0.0

build:
    go build -ldflags -o out/bin/${K8SCONTEXTTOOLS_BIN_NAME} github.com/edwardpro/k8scontexttools/cmd/k8scontexttools
