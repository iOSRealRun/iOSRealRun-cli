package utils

import (
	"os/exec"
	"strings"
)

func Cmd(i_cmd []string, getoutp ...bool) string {
	var cmd *exec.Cmd
	if len(i_cmd) == 1 {
		cmd = exec.Command(i_cmd[0])
	} else if len(i_cmd) > 1 {
		cmd = exec.Command(i_cmd[0], i_cmd[1:]...)
	}

	if len(getoutp) > 0 && !getoutp[0] {
		// don't panic when exit code is not 0
		err := cmd.Run()
		if err != nil {
			return ""
		}
		return ""
	} else {
		// merge stderr to stdout and don't panic when exit code is not 0
		cmd.Stderr = cmd.Stdout
		output, _ := cmd.Output()
		return string(output)
	}
}

func CmdWithlibimobidevice(i_cmd []string, libimobiledeviceDir string, getoutp ...bool) string {
	i_cmd[0] = strings.Join([]string{libimobiledeviceDir, i_cmd[0]}, "/")
	if GetOS() == "win" {
		i_cmd[0] = strings.Replace(i_cmd[0], "/", "\\", -1)
	}
	return Cmd(i_cmd, getoutp...)
}
