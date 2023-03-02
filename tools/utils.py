def cmd(i_cmd, getoutp=True):
    from main import seperator
    from main import libimobiledeviceDir
    from main import OS, env
    import subprocess
    if type(i_cmd) == str:
        i_cmd = seperator.join([libimobiledeviceDir, i_cmd])
    else:
        i_cmd[0] = seperator.join([libimobiledeviceDir, i_cmd[0]])
    if getoutp:
        return subprocess.Popen(i_cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, env=env[OS]).stdout.read().decode("utf-8")
    else:
        subprocess.run(i_cmd, env=env[OS])


def getOS():
    import sys
    OS = sys.platform
    if -1 != OS.find("win32"):
        return "win"
    elif -1 != OS.find("darwin"):
        return "darwin"
    else:
        return "linux"

def pair() -> int:
    resp = cmd(["idevicepair", "pair"])
    if -1 != resp.find("SUCCESS"):
        return 0
    if -1 != resp.find("No device found"):
        return 1
    if -1 != resp.find("passcode"):
        while -1 != resp.find("passcode"):
            input("请解锁手机后按回车")
            resp = cmd(["idevicepair", "pair"])
        if -1 != resp.find("SUCCESS"):
            return 0
    if -1 != resp.find("trust"):
        while -1 != resp.find("trust"):
            input("请在你的手机/或平板上按提示信任此电脑并按回车")
            resp = cmd(["idevicepair", "pair"])
        if -1 != resp.find("SUCCESS"):
            return 0
        else:
            return -1
    else:
        return -1

def getDeviceInfo():
    import re
    info = cmd("ideviceinfo")
    deviceName = re.search(r"DeviceName: (.+)\n", info).group(1).strip()
    version = re.search(r"ProductVersion: (.+)\n", info).group(1).strip()
    return deviceName, version


def setLoc(loc):
    cmd(["idevicesetlocation", "--", str(loc["lat"]-0.00389), str(loc["lng"]-0.01075)], False)

def resetLoc():
    cmd(["idevicesetlocation", "reset"], False)
