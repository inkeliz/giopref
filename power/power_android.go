package power

//go:generate javac -source 8 -target 8 -bootclasspath $ANDROID_HOME\platforms\android-29\android.jar -d $TEMP\giopref_power\classes power_android.java
//go:generate jar cf power_android.jar -C $TEMP\giopref_power\classes .

import (
	"gioui.org/app"
	"git.wow.st/gmp/jni"
)

var (
	_Lib = "com/inkeliz/giopref/power/power_android"
)

func batteryLevel() uint8 {
	return uint8(doInt("batteryLevel"))
}

func isSavingBattery() bool {
	return doBool("isSaving")
}

func isCharging() bool {
	return doBool("isCharging")
}

func doBool(f string) (enabled bool) {
	jni.Do(jni.JVMFor(app.JavaVM()), func(env jni.Env) error {
		class, err := jni.LoadClass(env, jni.ClassLoaderFor(env, jni.Object(app.AppContext())), _Lib)
		if err != nil {
			return err
		}

		obj, err := jni.NewObject(env, class, jni.GetMethodID(env, class, "<init>", `()V`))
		if err != nil {
			return err
		}

		obj, err = jni.CallObjectMethod(env, obj, jni.GetMethodID(env, class, f, "(Landroid/content/Context;)Z"), jni.Value(app.AppContext()))
		if err != nil {
			return err
		}

		enabled = obj != 0
		return nil
	})

	return enabled
}

func doInt(f string) (i int) {
	jni.Do(jni.JVMFor(app.JavaVM()), func(env jni.Env) error {
		class, err := jni.LoadClass(env, jni.ClassLoaderFor(env, jni.Object(app.AppContext())), _Lib)
		if err != nil {
			return err
		}

		obj, err := jni.NewObject(env, class, jni.GetMethodID(env, class, "<init>", `()V`))
		if err != nil {
			return err
		}

		obj, err = jni.CallObjectMethod(env, obj, jni.GetMethodID(env, class, f, "(Landroid/content/Context;)I"), jni.Value(app.AppContext()))
		if err != nil {
			return err
		}

		i = int(int32(obj))
		return nil
	})

	return i
}
