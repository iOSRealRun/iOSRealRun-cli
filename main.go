package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/device"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/initialize"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/run"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/utils"
)

func main() {
	// redirect stderr to ./log/error.log
	log, err := os.OpenFile("./log/error.log", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("无法创建日志文件")
		fmt.Println("按回车退出")
		fmt.Scanln()
		os.Exit(1)
	}
	os.Stderr = log

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

	// connect to the device and mount DeveloperDiskImage
	initialize.Connect()

	loc := initialize.Init()
	fmt.Println("路线信息读取成功")

	if OS == "win" {
		utils.SetDisplayRequired()

	}

	fmt.Printf("已开始模拟跑步，速度大约为 %s m/s\n", strconv.FormatFloat(v, 'f', -1, 64))
	fmt.Println("会无限绕圈，要停止请按Ctrl+C")
	fmt.Println("请勿直接关闭窗口，否则无法还原正常定位")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		// catch Ctrl+C
		<-c
		device.ResetLoc()
		if OS == "win" {
			utils.ResetDisplayRequired()
		}
		fmt.Println("现在你可以正常关闭当前窗口或终端了")
		os.Exit(0)
	}()
	run.Run(loc, v, 15)
}
