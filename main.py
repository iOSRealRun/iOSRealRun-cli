import tools.utils as utils
import tools.run as run
from tools.initialize import connect, init
from config import v, libimobiledeviceDir

def final():
    utils.resetLoc()  # reset location


# path separators for different systems
seperator = {"win": "\\", "darwin": "/", "linux": "/"}
seperator = seperator[utils.getOS()]
libimobiledeviceDir += seperator + utils.getOS()

# connect to the device and mount DevelopDiskImage
connect()

loc = init()  # get the route
print("路线信息读取成功")


print("已开始模拟跑步")
print("会无限绕圈，要停可以按Ctrl+C，按一次不行的画就一直按着")
print("请勿直接关闭窗口，否则可能定位异常")
try:
    run.run(loc, v)
finally:
    utils.resetLoc()  # reset location
