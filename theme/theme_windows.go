package theme

import (
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"unsafe"
)

var (
	_RegistryPersonalize = `SOFTWARE\Microsoft\Windows\CurrentVersion\Themes\Personalize`

	_User32           = windows.NewLazySystemDLL("user32")
	_SystemParameters = _User32.NewProc("SystemParametersInfoW")
)

const (
	_SPI_GETCLIENTAREAANIMATION = 0x1042
)

func isDark() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, _RegistryPersonalize, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()

	v, _, err := k.GetIntegerValue("AppsUseLightTheme")
	if err != nil {
		return false
	}

	return v == 0
}

func isReducedMotion() bool {
	disabled := true
	_SystemParameters.Call(uintptr(_SPI_GETCLIENTAREAANIMATION), 0, uintptr(unsafe.Pointer(&disabled)), 0)
	return !disabled
}
