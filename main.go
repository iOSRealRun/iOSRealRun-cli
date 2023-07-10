package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/device"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/initialize"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/logger"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/run"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/utils"
)

var myLogger = logger.NewMyLogger("log.log", config.Config.LogLevel)

func main() {
	conf := config.Config
	OS := utils.GetOS()
	CWD, _ := os.Getwd()
	v := conf.V

	// environment variables
	switch OS {
	case "darwin":
		os.Setenv("DYLD_LIBRARY_PATH", fmt.Sprintf("%s/%s", CWD, conf.LibimobiledeviceDir))
	case "linux":
		os.Setenv("LD_LIBRARY_PATH", fmt.Sprintf("%s/%s", CWD, conf.LibimobiledeviceDir))
	}
	myLogger.Infoln("Set environment variables for library path")

	// connect to the device and mount DeveloperDiskImage
	initialize.Connect()
	myLogger.Infoln("Connected to the device and mounted DeveloperDiskImage")

	loc := initialize.Init()
	fmt.Println("路线信息读取成功")
	myLogger.Infoln("Route information read successfully")

	if OS == "win" {
		utils.SetDisplayRequired()
		myLogger.Infoln("Set display required")
	}

	fmt.Printf("已开始模拟跑步，速度大约为 %s m/s\n", strconv.FormatFloat(v, 'f', -1, 64))
	fmt.Println("会无限绕圈，要停止请按Ctrl+C")
	fmt.Println("请勿直接关闭窗口，否则无法还原正常定位")

	// catch Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		device.ResetLoc()
		if OS == "win" {
			utils.ResetDisplayRequired()
			myLogger.Infoln("Reset display required")
		}
		fmt.Println("现在你可以正常关闭当前窗口或终端了")
		os.Exit(0)
	}()

	myLogger.Infoln("Startig running")
	run.Run(loc, v, 15)
}
