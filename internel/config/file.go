package config

import (
	"github.com/edwardpro/k8scontexttools/internel/utils"
	filepath2 "path/filepath"
)

var (
	KubeConfig string
)

func init() {
	filepath, _ := utils.GetHomeDir()
	KubeConfig = filepath2.Join(filepath, ".kube", "config")
}

func InitConfigFile() {

}
