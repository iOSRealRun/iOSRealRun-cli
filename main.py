#!/usr/bin/env python3
"""
main.py
the entrance of the program
"""

import sys
import os
import tools.utils as utils
import tools.run as run
from tools.initialize import connect, init
from tools import runtime


def main():
    if not os.path.exists("./log"):
        os.mkdir("./log")
    sys.stderr = open("./log/error.log", "w")  # redirect error message

    running = False
    display_required = False
    try:
        # connect to the device and mount DevelopDiskImage
        connect()

        loc = init()  # get the route
        print("路线信息读取成功")

        if runtime.OS == "win":
            utils.setDisplayRequired()
            display_required = True
        print("已开始模拟跑步, 速度大约为 {} m/s".format(str(runtime.v)))
        print("会无限绕圈，要停可以按Ctrl+C")
        print("请勿直接关闭窗口，否则无法还原正常定位")

        running = True
        run.run(loc, runtime.v)
    except KeyboardInterrupt:
        print("已停止模拟跑步")
    except Exception as exc:
        import traceback

        print("程序遇到错误：{}".format(exc))
        print("详细错误已写入 log/error.log")
        traceback.print_exc()
        return 1
    finally:
        if running:
            try:
                utils.resetLoc()  # reset the location
            except Exception as exc:
                import traceback

                print("恢复正常定位失败：{}".format(exc))
                traceback.print_exc()
        if display_required:
            utils.resetDisplayRequired()
        print("现在你可以关闭当前窗口或终端了")
    return 0


if __name__ == "__main__":
    sys.exit(main())
