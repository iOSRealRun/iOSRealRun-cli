import tools.utils as utils


def connect() -> int:
    import os
    import sys
    from main import seperator
    from main import OS
    from config import imageDir, libimobiledeviceDir

    if OS != "win":
        os.system("chmod -R +rx " + libimobiledeviceDir)
        if OS == "linux":
            path = os.environ["PATH"].split(":")
            check = False
            for i in path:
                if os.path.exists("{}/{}".format(i, "usbmuxd")):
                    check = True
            if not check:
                print("你没有安装usbmuxd，请再次阅读README中关于Linux的部分")
                input("按回车退出")
        else:
            utils.cmd(["xattr", "-d", "-r", "com.apple.quarantine", "."], getoutp=False)

    input("请解锁手机或Pad后按回车")
    status = utils.pair()
    while status == 1:
        print("无设备连接，Windows需要安装iTunes，也可尝试解锁手机并插拔数据线，如果还是不行Mac和Windows请打开iTunes并在跑完前不要关闭")
        input("确定连接后按回车")
        status = utils.pair()
    if -1 == status:
        print("遇到了未知的问题")
        print("按回车退出")
        sys.exit()

    deviceName, version = utils.getDeviceInfo()
    print("已连接到 {}".format(deviceName))
    print("系统版本：{}".format(version))

    if int(version.split(".")[0]) >= 16:
        developerMode = -1 == utils.cmd(["idevicedevmodectl", "list"]).find("disable")
        if not developerMode:
            utils.cmd(["idevicedevmodectl", "reveal"], False)
            print("请在系统设置-隐私与安全性-开发者模式中打开开发者模式")
            print("可能需要按要求重启手机/pad")
            print("请在开启开发者模式之后重新打开本脚本，开机后请不要急，等确认所有开发者模式相关的弹出框再打开本脚本")
            input("现在按回车退出")
            sys.exit()
    
    imageStatus = os.path.exists("{}/{}/DeveloperDiskImage.dmg".format(imageDir, version)) and\
                  os.path.exists("{}/{}/DeveloperDiskImage.dmg.signature".format(imageDir, version))
    if not imageStatus:
        version = ".".join(version.split(".")[0:2])
        imageStatus = os.path.exists("{}/{}/DeveloperDiskImage.dmg".format(imageDir, version)) and\
                    os.path.exists("{}/{}/DeveloperDiskImage.dmg.signature".format(imageDir, version))

    if not imageStatus:
        print("没有在 {} 下找到 {} 版本的开发者镜像".format(imageDir, version))
        print("请添加完再打开本脚本")
        print("按回车退出")
        sys.exit()

    imageCMD = [
        "ideviceimagemounter",
        "{}{}{}{}DeveloperDiskImage.dmg".format(imageDir, seperator, version, seperator),
        "{}{}{}{}DeveloperDiskImage.dmg.signature".format(imageDir, seperator, version, seperator),
    ]
    if -1 != utils.cmd(imageCMD).find("-3"):
        print("开发者镜像签名验证失败，你要重新下一遍")
        print("完成后再打开脚本")
        print("按回车退出")
        sys.exit()

def init():
    import tools.snippet as snippet
    from config import routeConfig
    with open(routeConfig) as myFile:
        loc = snippet.split(myFile.read())
    return loc
