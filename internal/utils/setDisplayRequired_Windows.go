//go:build windows
// +build windows

package utils

import "syscall"

const (
	ES_CONTINUOUS       = 0x80000000
	ES_DISPLAY_REQUIRED = 0x00000002
)

func SetDisplayRequired() {

	dll, _ := syscall.LoadDLL("kernel32.dll")
	proc, _ := dll.FindProc("SetThreadExecutionState")
	proc.Call(uintptr(ES_CONTINUOUS | ES_DISPLAY_REQUIRED))
}

func ResetDisplayRequired() {
	dll, _ := syscall.LoadDLL("kernel32.dll")
	proc, _ := dll.FindProc("SetThreadExecutionState")
	proc.Call(uintptr(ES_CONTINUOUS))
}
