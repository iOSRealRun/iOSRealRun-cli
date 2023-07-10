package device

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/logger"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/utils"
)

var myLogger = logger.NewMyLogger("log.log", config.Config.LogLevel)

func Pair() int {
	myLogger.Infoln("Start pairing ...")
	conf := config.Config
	resp := utils.CmdWithlibimobidevice([]string{"idevicepair", "pair"}, conf.LibimobiledeviceDir)
	myLogger.Debugln("resp:", resp)
	if strings.Contains(resp, "SUCCESS") {
		myLogger.Infoln("Pair success")
		return 0
	}
	if strings.Contains(resp, "No device found") {
		myLogger.Warnln("No device found")
		return 1
	}
	if strings.Contains(resp, "passcode") {
		for strings.Contains(resp, "passcode") {
			fmt.Println("请解锁手机后按回车")
			fmt.Scanln()
			resp = utils.CmdWithlibimobidevice([]string{"idevicepair", "pair"}, conf.LibimobiledeviceDir)
			myLogger.Debugln("resp:", resp)
		}
		if strings.Contains(resp, "SUCCESS") {
			myLogger.Infoln("Pair success")
			return 0
		}
	}
	if strings.Contains(resp, "trust") {
		for strings.Contains(resp, "trust") {
			fmt.Println("请在你的手机/或平板上按提示信任此电脑并按回车")
			fmt.Scanln()
			resp = utils.CmdWithlibimobidevice([]string{"idevicepair", "pair"}, conf.LibimobiledeviceDir)
			myLogger.Debugln("resp:", resp)
		}
		if strings.Contains(resp, "SUCCESS") {
			myLogger.Infoln("Pair success")
			return 0
		} else {
			myLogger.Errorln("Unknown error")
			return -1
		}
	} else {
		myLogger.Errorln("Unknown error")
		return -1
	}
}

func GetDeviceInfo() (deviceName string, version string) {
	info := utils.CmdWithlibimobidevice([]string{"ideviceinfo"}, config.Config.LibimobiledeviceDir)
	deviceName = regexp.MustCompile(`DeviceName: (.*)`).FindStringSubmatch(info)[1]
	version = regexp.MustCompile(`ProductVersion: (.*)`).FindStringSubmatch(info)[1]
	myLogger.Debugln("deviceName:", deviceName)
	myLogger.Debugln("version:", version)
	myLogger.Debugln("info:", info)
	// strip
	deviceName = strings.Trim(deviceName, " \n\r\t")
	version = strings.Trim(version, " \n\r\t")
	return
}
