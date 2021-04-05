package power

import (
	"golang.org/x/sys/windows"
	"runtime"
	"unsafe"
)

var (
	_Kernel32             = windows.NewLazySystemDLL("kernel32")
	_GetSystemPowerStatus = _Kernel32.NewProc("GetSystemPowerStatus")
)

func batteryLevel() uint8 {
	resp, err := powerStatus()
	if err != nil || resp.BatteryLifePercent == 255 {
		return 100
	}

	return resp.BatteryLifePercent
}

func isSavingBattery() bool {
	resp, err := powerStatus()
	if err != nil {
		return false
	}

	return resp.SystemStatusFlag == 1
}

func isCharging() bool {
	resp, err := powerStatus()
	if err != nil {
		return true
	}

	return resp.BatteryFlag == 8 || resp.BatteryFlag >= 128
}

type _SystemPowerStatus struct {
	ACLineStatus        byte
	BatteryFlag         byte
	BatteryLifePercent  byte
	SystemStatusFlag    byte
	BatteryLifeTime     int32
	BatteryFullLifeTime int32
}

func powerStatus() (resp _SystemPowerStatus, err error) {
	runtime.KeepAlive(resp)
	r, _, err := _GetSystemPowerStatus.Call(uintptr(unsafe.Pointer(&resp)))
	if r == 0 {
		return resp, err
	}
	return resp, nil
}
