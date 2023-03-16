#!/usr/bin/env python3

import sys
import os
import tools.utils as utils
import tools.run as run
from tools.initialize import connect, init
from config import v, libimobiledeviceDir

if not os.path.exists("./log"):
    os.mkdir("./log")
sys.stderr=open("./log/error.log", "w")  # redirect error message


OS = utils.getOS()
# path separators for different systems
seperator = {"win": "\\", "darwin": "/", "linux": "/"}
seperator = seperator[OS]
libimobiledeviceDir += seperator + OS
# environment variables
env = {
    "win": None,
    "darwin": {"DYLD_LIBRARY_PATH": os.getcwd() + "/" + libimobiledeviceDir},
    "linux": {"LD_LIBRARY_PATH": os.getcwd() + "/" + libimobiledeviceDir}
}

# connect to the device and mount DevelopDiskImage
connect()

loc = init()  # get the route
print("路线信息读取成功")


if OS == "win":
    utils.setDisplayRequired()
print("已开始模拟跑步, 速度大约为 {} m/s".format(str(v)))
print("会无限绕圈，要停可以按Ctrl+C")
print("请勿直接关闭窗口，否则无法还原正常定位")

try:
    run.run(loc, v)
finally:
    utils.resetLoc()
    if OS == "win":
        utils.resetDisplayRequired()
    print("现在你可以关闭当前窗口或终端了")