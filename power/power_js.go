package power

import (
	"errors"
	"math"
	"syscall/js"
)

var (
	_GetBattery = js.Global().Get("navigator").Get("getBattery")
)

func batteryLevel() uint8 {
	value, err := do("level")
	if err != nil || !value.Truthy() {
		return 100
	}

	b := uint8(math.Ceil(value.Float() * 100))
	switch {
	case b > 100:
		return 100
	case b < 0:
		return 1
	default:
		return b
	}
}

func isSavingBattery() bool {
	return false
}

func isCharging() bool {
	value, err := do("charging")
	if err != nil || !value.Truthy() {
		return false
	}

	return value.Bool()
}

func do(name string) (js.Value, error) {
	if !_GetBattery.Truthy() {
		return js.Value{}, errors.New("API not available")
	}

	var (
		success, failure js.Func

		value = make(chan js.Value, 1)
		err   = make(chan error, 1)
	)

	success = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		success.Release()
		failure.Release()

		value <- args[0].Get(name)

		return nil
	})

	failure = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		success.Release()
		failure.Release()

		err <- errors.New("failure getting battery")

		return nil
	})

	go func() {
		js.Global().Get("navigator").Call("getBattery").Call("then", success, failure)
	}()

	select {
	case value := <-value:
		return value, nil
	case err := <-err:
		return js.Value{}, err
	}
}
