package utils

import "syscall"

const (
	ES_CONTINUOUS       = 0x80000000
	ES_DISPLAY_REQUIRED = 0x00000002
)

func SetDisplayRequired() {
	// python:
	// import ctypes
	// ctypes.windll.kernel32.SetThreadExecutionState(ES_CONTINUOUS | ES_DISPLAY_REQUIRED)
	// go version:
	dll, _ := syscall.LoadDLL("kernel32.dll")
	proc, _ := dll.FindProc("SetThreadExecutionState")
	proc.Call(uintptr(ES_CONTINUOUS | ES_DISPLAY_REQUIRED))
}

func ResetDisplayRequired() {
	// python:
	// import ctypes
	// ctypes.windll.kernel32.SetThreadExecutionState(ES_CONTINUOUS)
	// go version:
	dll, _ := syscall.LoadDLL("kernel32.dll")
	proc, _ := dll.FindProc("SetThreadExecutionState")
	proc.Call(uintptr(ES_CONTINUOUS))
}
