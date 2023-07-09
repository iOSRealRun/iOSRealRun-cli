package initialize

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/device"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/parseRoute"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/utils"
)

func Connect() {
	conf := config.Config
	OS := utils.GetOS()

	if OS != "win" {
		utils.Cmd([]string{"chmod", "-R", "+rx", conf.LibimobiledeviceDir}, false)
		if OS == "linux" {
			path := strings.Split(os.Getenv("PATH"), ":")
			check := false
			for _, i := range path {
				if utils.FileExists(fmt.Sprintf("%s/%s", i, "usbmuxd")) {
					check = true
					break
				}
			}
			if !check {
				fmt.Println("你没有安装usbmuxd，请再次阅读README中关于Linux的部分")
				fmt.Println("按回车退出")
				fmt.Scanln()
			}
		} else {
			quarantine := strings.Contains(utils.Cmd([]string{"xattr", "libimobiledevice/darwin/ideviceinfo"}), "quarantine")
			if quarantine {
				utils.Cmd([]string{"sudo", "xattr", "-d", "-r", "com.apple.quarantine", "."}, false)
			}
		}
	}

	fmt.Println("请解锁手机或iPad，然后按回车键继续")
	fmt.Scanln()
	status := device.Pair()
	for status == 1 {
		fmt.Println("无设备连接，Windows需要安装iTunes，也可尝试解锁手机并插拔数据线，如果还是不行Mac和Windows请打开iTunes并在跑完前不要关闭")
		fmt.Println("确定连接后按回车键继续")
		fmt.Scanln()
		status = device.Pair()
	}
	if status == -1 {
		fmt.Println("遇到了位置的问题")
		fmt.Println("按回车退出")
		os.Exit(1)
	}

	deviceName, version := device.GetDeviceInfo()
	fmt.Printf("已连接到%s\n", deviceName)
	fmt.Printf("系统版本：%s\n", version)

	majorVersion, _ := strconv.ParseInt(strings.Split(version, ".")[0], 10, 64)
	if majorVersion >= 16 {
		developerMode := !strings.Contains(utils.Cmd([]string{"idevicedevmodectl", "list"}), "disabled")
		if !developerMode {
			utils.Cmd([]string{"idevicedevmodectl", "reveal"})
			fmt.Println("请在系统设置-隐私与安全性-开发者模式中打开开发者模式")
			fmt.Println("可能需要按要求重启手机/pad")
			fmt.Println("请在开启开发者模式吹常在重新打开本脚本，开机后请不要急，等确认所有开发者模式相关的弹出框再打开本脚本")
			fmt.Println("现在按回车退出")
			fmt.Scanln()
			os.Exit(0)
		}

		imageStatus := utils.FileExists(fmt.Sprintf("%s/%s/DeveloperDiskImage.dmg", conf.ImageDir, version)) &&
			utils.FileExists(fmt.Sprintf("%s/%s/DeveloperDiskImage.dmg.signature", conf.ImageDir, version))
		if !imageStatus {
			version = strings.Join(strings.Split(version, ".")[:2], ".")
			imageStatus = utils.FileExists(fmt.Sprintf("%s/%s/DeveloperDiskImage.dmg", conf.ImageDir, version)) &&
				utils.FileExists(fmt.Sprintf("%s/%s/DeveloperDiskImage.dmg.signature", conf.ImageDir, version))
		}

		if !imageStatus {
			fmt.Printf("没有在 %s 下找到 %s 版本的开发者镜像\n", conf.ImageDir, version)
			fmt.Println("请添加完后再次运行本脚本")
			fmt.Println("现在按回车退出")
			fmt.Scanln()
			os.Exit(0)
		}

		imageCMD := []string{
			"ideviceimagemounter",
			fmt.Sprintf("%s/%s/DeveloperDiskImage.dmg", conf.ImageDir, version),
			fmt.Sprintf("%s/%s/DeveloperDiskImage.dmg.signature", conf.ImageDir, version),
		}
		if strings.Contains(utils.Cmd(imageCMD), "-3") {
			fmt.Println("开发者镜像签名验证失败，你要重新下一遍")
			fmt.Println("完成后再次运行本脚本")
			fmt.Println("现在按回车退出")
			fmt.Scanln()
			os.Exit(0)
		}
	}
}

func Init() []map[string]float64 {
	conf := config.Config
	// read file
	content, err := os.ReadFile(conf.RouteConfig)
	if err != nil {
		panic(err)
	}
	return parseRoute.Split(content)
}
