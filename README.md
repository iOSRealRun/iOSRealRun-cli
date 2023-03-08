iOSRealRun-cli
---

## 背景
原 [iOSFakeRun](https://github.com/Mythologyli/iOSFakeRun) 只能在windows上用，且每圈的路径是一样的。同时对于很多iOS/iPadOS16及以上的设备，不能方便地打开开发者模式，但是由于我不会C#无法给其提PR，所以有了这个项目

## 功能
- 已实现
  - [x] 让你方便地打开开发者模式，而不需要借助爱思助手等第三方软件  
  - [x] 和 [iOSFakeRun](https://github.com/Mythologyli/iOSFakeRun) 相同的路径格式 (**但是请只画一圈**)  
  - [x] 自定义配速  
  - [x] 一种阴间的随机方式让每一圈不完全相同  
  - [x] 支持 Windows 和 MacOS  
  - [x] 每跑完一次设定的路径随机一次速度  
  - [x] 支持 Linux/BSD  
- 暂未实现
  - [ ] 阳间的随机方式  

## 原理
不讲了，懂的自然懂

## 要求
- Python 3  
- DeveloperDiskImage (已预置了15.4及以上，其余版本请按照下文使用方法中的指引操作)  
- usbmuxd (对于Linux用户)  

## 支持的系统
- 已支持
  - [x] Windows 64位 (tested on Windows 11, Python 3.9)  
  - [x] MacOS (Apple Silicon) (tested on M2 Mac, Python 3.9) 
  - [x] MacOS (Intel) (不知道行不行，等人试)  
  - [x] linux/bsd 64位 
- 不支持
  - [ ] 任何x86的32位系统

## 使用方法
  大概步骤:  
  1. 安装依赖：
     1. 自行安装Python 3  
     2. 若为Linux用户，请用你的包管理器安装`usbmuxd`  
        Debian/Ubuntu: `sudo apt install usbmuxd`  
        Red Hat/CentOS/Fedora: `sudo dnf install usbmuxd`  
  2. 从 [Release](https://github.com/Geniucker/iOSFakeRun-cli/releases) 下载代码到本地并解压到你喜欢的地方，接下来称`main.py`所在文件夹位**脚本目录  
  3. 虚拟定位需要开发者镜像。我预置了 15.4 及以上的开发者镜像  
     如果你不是上述版本，你可以去网上找 DeveloperDiskImage  
     - 打开 [DeveloperDiskImage](https://github.com/mspvirajpatel/Xcode_Developer_Disk_Images/releases) 仓库  
     - 查看自己的 iOS 版本，下载对应的 `DeveloperDiskImage.dmg` 和 `DeveloperDiskImage.dmg.signature` 文件  
     - 进入脚本目录 `DeveloperDiskImage` 文件夹中建立以版本号为名称的文件夹，将刚才下载的两个文件放入此文件夹。  
       例如，你是 15.1 版本的系统，你需要下载并解压 `DeveloperDiskImage.dmg` 和 `DeveloperDiskImage.dmg.signature` 文件，把他们放到 `DeveloperDiskImage/15.1` 里面  
  4. 接下来和 [iOSFakeRun](https://github.com/Mythologyli/iOSFakeRun) 一样要获取你要的跑步路径，格式和其使用的格式完全相同，**但是请只画一圈**，项目预置了一个画的不太行的海宁操场路径，建议所有人都自己画路径  
     > 打开[路径拾取网站](https://fakerun.myth.cx/)。通过点击地图构造路径。点击时无需考虑间距，会自动用直线连接。路径点击完成后，单击上方的路径坐标——复制，将坐标数据复制到剪贴板  
  5. 打开脚本目录里的 `route.txt` 文件，将刚复制的文件原封不动的粘贴进去，保存并退出  
  6. 对于 Windows，你需要安装 iTunes，以确保驱动正常运行  
  7. 在脚本目录中的 `config.py` 文件中设置 `v` 变量以设置速度(m/s)，给个参考，3.3大概是5分到5分半的配速（我也没仔细看）  
  8. Mac用户和Windows用户可能需要打开iTunes，**且在跑完前不应关闭**   
  8. 用数据线将电脑连接到 iPhone 或 iPad  
  9. Mac 用户打开 `main.command`，Windows 用户打开 `main.py`  
  10. 按照提示完成设备连接和开发者模式的开启  
  11. 都好了大概就开跑了，默认无限循环  
  12. 跑完之后请 **务必使用 Ctrl + C** 来停止，而非直接把窗口叉掉，否则不能自动恢复手机或pad的正常定位  

## FAQ
- 有一定的概率在第7、8步左右的时候，也就显示出你的系统版本之后会卡住，  
  ![图片](https://user-images.githubusercontent.com/61449208/222521731-fa7947f6-174b-46ab-8b8a-871aac7f6e5b.png)  
  这不是我的问题，我也没法修，解决方法是 Ctrl + C，这个时候脚本不会停，而是继续了，如果定位成功被修改了，那就不用管，如果定位模拟失败，那就继续 Ctrl + C ，再重新打开脚本  
- 刚打开 `main.py` 的时候可能不会马上有东西显示（尤其是MacOS），请耐心等待  

## 免责声明
本项目仅供 Python 和 C 学习交流作者对软件的用途不做任何说明或暗示。对使用本软件造成的一切后果概不负责  

## 致谢
- [iOSFakeRun](https://github.com/Mythologyli/iOSFakeRun)  
- [libimobiledevice](https://github.com/libimobiledevice/libimobiledevice)
