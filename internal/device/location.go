package device

import (
	"strconv"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/utils"
)

func SetLoc(loc map[string]float64) {
	// python: utils.CmdWithlibimobidevice(["idevicesetlocation", "--", str(loc["lat"]-0.00389), str(loc["lng"]-0.01075)], config.Config.LibimobiledeviceDir, False)
	// go version
	lat := loc["lat"]
	lng := loc["lng"]
	utils.CmdWithlibimobidevice([]string{"idevicesetlocation", "--", strconv.FormatFloat(lat-0.00389, 'f', -1, 64), strconv.FormatFloat(lng-0.01075, 'f', -1, 64)}, config.Config.LibimobiledeviceDir, false)
}

func ResetLoc() {
	utils.CmdWithlibimobidevice([]string{"idevicesetlocation", "reset"}, config.Config.LibimobiledeviceDir, false)
}
