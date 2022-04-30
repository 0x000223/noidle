package win32

import (
	"syscall"
	"unsafe"
)

var (
	moduleUser32            = syscall.NewLazyDLL("user32.dll")
	procFindWindow          = moduleUser32.NewProc("FindWindowW")
	procSetForegroundWindow = moduleUser32.NewProc("SetForegroundWindow")
	procGetForegroundWindow = moduleUser32.NewProc("GetForegroundWindow")
)

func GetForegroundWindow() syscall.Handle {
	r1, _, _ := procGetForegroundWindow.Call()
	return syscall.Handle(r1)
}

func SetForegroundWindow(windowHandle syscall.Handle) {
	procSetForegroundWindow.Call(uintptr(windowHandle))
}

func FindWindow(className string, windowName string) syscall.Handle {
	if len(className) == 0 && len(windowName) == 0 {
		return 0
	}

	ptrClassName, _ := syscall.UTF16PtrFromString(className)
	ptrWindowName, _ := syscall.UTF16PtrFromString(windowName)

	r1, _, _ := procFindWindow.Call(uintptr(unsafe.Pointer(ptrClassName)), uintptr(unsafe.Pointer(ptrWindowName)))

	return syscall.Handle(r1)
}
