package utils

import (
	"os"
)

func FileExists(path string) bool {
	myLogger.Traceln("path: ", path)
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	myLogger.Traceln("err: ", err)
	return !os.IsNotExist(err)
}
