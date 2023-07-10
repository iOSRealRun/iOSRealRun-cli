package utils

import (
	"runtime"
)

func GetOS() string {
	switch runtime.GOOS {
	case "windows":
		return "win"
	case "darwin":
		return "darwin"
	case "linux":
		return "linux"
	default:
		return "unknown"
	}
}
