import tools.utils as utils


def connect() -> int:
    import os
    import sys
    from main import seperator
    from config import imageDir
    status = utils.pair()
    while status == 1:
        print("无设备连接，Windows需要安装iTunes")
        input("确定连接后按回车")
        status = utils.pair()
    if -1 == status:
        print("遇到了未知的问题")
        print("按回车退出")
        sys.exit()

    deviceName, version = utils.getDeviceInfo()
    print("已连接到 {}".format(deviceName))
    print("系统版本：{}".format(version))
    
    imageStatus = os.path.exists("{}/{}/DeveloperDiskImage.dmg".format(imageDir, version)) and\
                  os.path.exists("{}/{}/DeveloperDiskImage.dmg.signature".format(imageDir, version))
    if not imageStatus:
        print("没有在 {} 下找到 {} 版本的开发者镜像".format(imageDir, version))
        print("请添加完再打开本脚本")
        print("按回车退出")
        sys.exit()

    imageCMD = "ideviceimagemounter {}{}{}{}DeveloperDiskImage.dmg {}{}{}{}DeveloperDiskImage.dmg.signature".format(*(2*[imageDir, seperator, version, seperator]))
    if -1 != utils.cmd(imageCMD).find("-3"):
        print("开发者镜像签名验证失败，你要重新下一遍")
        print("完成后再打开脚本")
        print("按回车退出")
        sys.exit()

def init():
    import tools.snippet as snippet
    with open("route.txt") as myFile:
        loc = snippet.split(myFile.read())
    return loc
