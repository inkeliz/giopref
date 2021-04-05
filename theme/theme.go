package theme

// IsDark returns true if the user prefers dark-mode
func IsDark() bool {
	return isDark()
}

// IsReducedMotion returns true if the user want low animation/disabled animation.
func IsReducedMotion() bool {
	return isReducedMotion()
}
