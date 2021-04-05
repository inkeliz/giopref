package theme

//go:generate javac -source 8 -target 8 -bootclasspath $ANDROID_HOME\platforms\android-29\android.jar -d $TEMP\giopref_theme\classes theme_android.java
//go:generate jar cf theme_android.jar -C $TEMP\giopref_theme\classes .

import (
	"gioui.org/app"
	"git.wow.st/gmp/jni"
)

var (
	_Lib = "com/inkeliz/giopref/theme/theme_android"
)

func isDark() bool {
	return do("isDark")
}

func isReducedMotion() bool {
	return do("isReducedMotion")
}

func do(f string) (enabled bool) {
	jni.Do(jni.JVMFor(app.JavaVM()), func(env jni.Env) error {
		class, err := jni.LoadClass(env, jni.ClassLoaderFor(env, jni.Object(app.AppContext())), _Lib)
		if err != nil {
			return err
		}

		obj, err := jni.NewObject(env, class, jni.GetMethodID(env, class, "<init>", `()V`))
		if err != nil {
			return err
		}

		obj, err = jni.CallObjectMethod(env, obj, jni.GetMethodID(env, class, f, "()Z"))
		if err != nil {
			return err
		}

		enabled = obj != 0
		return nil
	})

	return enabled
}
