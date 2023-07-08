package device

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/utils"
)

func Pair() int {
	conf := config.Config
	resp := utils.CmdWithlibimobidevice([]string{"idevicepair", "pair"}, conf.LibimobiledeviceDir)
	if strings.Contains(resp, "SUCCESS") {
		return 0
	} else if strings.Contains(resp, "No device found") {
		return 1
	} else if strings.Contains(resp, "passcode") {
		for strings.Contains(resp, "passcode") {
			fmt.Println("请解锁手机后按回车")
			fmt.Scanln()
			resp = utils.CmdWithlibimobidevice([]string{"idevicepair", "pair"}, conf.LibimobiledeviceDir)
		}
		if strings.Contains(resp, "SUCCESS") {
			return 0
		} else {
			return -1
		}
	} else if strings.Contains(resp, "trust") {
		for strings.Contains(resp, "trust") {
			fmt.Println("请在你的手机/或平板上按提示信任此电脑并按回车")
			fmt.Scanln()
			resp = utils.CmdWithlibimobidevice([]string{"idevicepair", "pair"}, conf.LibimobiledeviceDir)
		}
		if strings.Contains(resp, "SUCCESS") {
			return 0
		} else {
			return -1
		}
	} else {
		return -1
	}
}

func GetDeviceInfo() (deviceName string, version string) {
	info := utils.CmdWithlibimobidevice([]string{"ideviceinfo"}, config.Config.LibimobiledeviceDir)
	deviceName = regexp.MustCompile(`DeviceName: (.*)`).FindStringSubmatch(info)[1]
	version = regexp.MustCompile(`ProductVersion: (.*)`).FindStringSubmatch(info)[1]
	return
}
