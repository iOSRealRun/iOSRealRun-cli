def cmd(i_cmd):
    from main import seperator
    from main import libimobiledeviceDir
    import subprocess
    i_cmd = seperator.join([libimobiledeviceDir, i_cmd])
    return subprocess.Popen(i_cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE).stdout.read().decode("utf-8")

def getOS():
    import sys
    os = sys.platform
    if -1 != os.find("win"):
        return "win"
    elif -1 != os.find("darwin"):
        return "darwin"
    else:
        return "linux"

def pair() -> int:
    resp = cmd("idevicepair pair")
    if -1 != resp.find("SUCCESS"):
        return 0
    if -1 != resp.find("No device found"):
        return 1
    if -1 != resp.find("passcode"):
        while -1 != resp.find("passcode"):
            input("请解锁手机后按回车")
            resp = cmd("idevicepair pair")
        if -1 != resp.find("SUCCESS"):
            return 0
    if -1 != resp.find("trust"):
        while -1 != resp.find("trust"):
            input("请在你的手机/或平板上按提示信任此电脑并按回车")
            resp = cmd("idevicepair pair")
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
    cmd("idevicesetlocation -- " + str(loc["lat"]-0.00389) + " " + str(loc["lng"]-0.01071))

def resetLoc():
    cmd("idevicesetlocation reset")

