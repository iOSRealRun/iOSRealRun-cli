package utils

import (
	"runtime"
)

func GetOS() string {
	if runtime.GOOS == "windows" {
		return "win"
	} else if runtime.GOOS == "darwin" {
		return "darwin"
	} else if runtime.GOOS == "linux" {
		return "linux"
	} else {
		return "unknown"
	}
}
