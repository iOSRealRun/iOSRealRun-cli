package main

import (
	"fmt"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/utils"
)

func main() {
	conf := config.Config
	fmt.Println(utils.CmdWithlibimobidevice([]string{"idevicepair", "pair"}, conf.LibimobiledeviceDir))
}
