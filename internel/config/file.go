package config

import (
	"github.com/edwardpro/k8scontexttools/internel/utils"
	"os"
	filepath_s "path/filepath"
)

var (
	ConfigFile *os.File
)

func init() {
	filepath, _ := utils.GetHomeDir()

	KubeConfig := filepath_s.Join(filepath, ".kube", "config")
	if !utils.IsPathExist(KubeConfig) {
		KubeConfigDir := filepath_s.Join(filepath, ".kube")
		os.MkdirAll(KubeConfigDir, 0744)
		os.Create(KubeConfig)

	}
	file, _err := os.OpenFile(KubeConfig, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0744)
	if _err != nil {
		ConfigFile = file
	}

}

func InitConfigFile() {

}
