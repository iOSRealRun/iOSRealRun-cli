package config

import (
	"os"
	"strings"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/utils"
	"gopkg.in/yaml.v3"
)

type config struct {
	V                   int    `yaml:"v"`
	RouteConfig         string `yaml:"routeConfig"`
	LibimobiledeviceDir string `yaml:"libimobiledeviceDir"`
	ImageDir            string `yaml:"imageDir"`
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

	// set imageDir
	os := utils.GetOS()
	conf.LibimobiledeviceDir = strings.Join([]string{conf.LibimobiledeviceDir, os}, "/")
	if os == "win" { // substitute / to \
		conf.LibimobiledeviceDir = strings.Replace(conf.LibimobiledeviceDir, "/", "\\", -1)
	}

	return
}

var Config config = SetupConfig()
