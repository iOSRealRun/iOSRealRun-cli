package utils

import (
	"os/exec"
	"strings"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/config"
	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/logger"
)

var myLogger = logger.NewMyLogger("log.log", config.Config.LogLevel)

func Cmd(i_cmd []string, getoutp ...bool) string {
	myLogger.Debugln("i_cmd: ", i_cmd)
	myLogger.Debugln("getoutp: ", getoutp)
	var cmd *exec.Cmd
	if len(i_cmd) == 1 {
		cmd = exec.Command(i_cmd[0])
	} else if len(i_cmd) > 1 {
		cmd = exec.Command(i_cmd[0], i_cmd[1:]...)
	}

	if len(getoutp) > 0 && !getoutp[0] {
		// don't panic when exit code is not 0
		err := cmd.Run()
		myLogger.Debugln("err: ", err)
		return ""
	} else {
		// merge stderr to stdout and don't panic when exit code is not 0
		cmd.Stderr = cmd.Stdout
		output, err := cmd.Output()
		myLogger.Debugln("err: ", err)
		return string(output)
	}
}

func CmdWithlibimobidevice(i_cmd []string, libimobiledeviceDir string, getoutp ...bool) string {
	myLogger.Debugln("i_cmd: ", i_cmd)
	myLogger.Debugln("getoutp: ", getoutp)
	i_cmd[0] = strings.Join([]string{libimobiledeviceDir, i_cmd[0]}, "/")
	return Cmd(i_cmd, getoutp...)
}
