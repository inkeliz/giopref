package theme

import (
	"syscall/js"
)

var (
	_MatchMedia = js.Global().Get("matchMedia")
)

func isDark() bool {
	return do("(prefers-color-scheme: dark)")
}

func isReducedMotion() bool {
	return do("(prefers-reduced-motion: reduce)")
}

func do(name string) bool {
	if !_MatchMedia.Truthy() {
		return false
	}

	return _MatchMedia.Invoke(name).Get("matches").Bool()
}
