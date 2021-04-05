package power

// Level returns the battery level as percent level, between 0~100.
// If the device doesn't run on battery it will return 100.
func Level() uint8 {
	return batteryLevel()
}

// IsSavingBattery returns true if the user is saving battery, when want to reduce the power consumption.
func IsSavingBattery() bool {
	return isSavingBattery()
}

// IsCharging returns true if the device is charging.
// If the device doesn't relies on batteries it will be always true.
func IsCharging() bool {
	return isCharging()
}