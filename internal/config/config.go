package config

import (
	"os"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

type config struct {
	V                   float64 `yaml:"v"`
	RouteConfig         string  `yaml:"routeConfig"`
	LibimobiledeviceDir string  `yaml:"libimobiledeviceDir"`
	ImageDir            string  `yaml:"imageDir"`
	LogLevel            string  `yaml:"log-level"`
}

func SetupConfig() (conf config) {
	conf = config{}
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		panic(err)
	}

	var OS string
	switch runtime.GOOS {
	case "windows":
		OS = "win"
	case "darwin":
		OS = "darwin"
	case "linux":
		OS = "linux"
	default:
		OS = "unknown"
	}

	conf.LibimobiledeviceDir = strings.Join([]string{conf.LibimobiledeviceDir, OS}, "/")

	return
}

var Config config = SetupConfig()
