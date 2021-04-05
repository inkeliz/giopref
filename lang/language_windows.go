package lang

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	_Kernel32 = windows.NewLazySystemDLL("kernel32")

	_DefaultUserLang   = _Kernel32.NewProc("GetUserDefaultLocaleName")
	_DefaultSystemLang = _Kernel32.NewProc("GetSystemDefaultLocaleName")

	_DefaultMaxSize = 85
)

func getLanguage() string {
	lang := make([]uint16, _DefaultMaxSize)

	r, _, _ := _DefaultUserLang.Call(uintptr(unsafe.Pointer(&lang[0])), uintptr(_DefaultMaxSize))
	if r == 0 {
		_DefaultSystemLang.Call(uintptr(unsafe.Pointer(&lang[0])), uintptr(_DefaultMaxSize))
	}

	return windows.UTF16ToString(lang)
}
