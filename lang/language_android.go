package lang

//go:generate javac -source 8 -target 8 -bootclasspath $ANDROID_HOME\platforms\android-29\android.jar -d $TEMP\giopref_lang\classes language_android.java
//go:generate jar cf language_android.jar -C $TEMP\giopref_lang\classes .

import (
	"gioui.org/app"
	"git.wow.st/gmp/jni"
)

var (
	_Lib = "com/inkeliz/giopref/language/language_android"
)

func getLanguage() (lang string) {
	err := jni.Do(jni.JVMFor(app.JavaVM()), func(env jni.Env) error {
		class, err := jni.LoadClass(env, jni.ClassLoaderFor(env, jni.Object(app.AppContext())), _Lib)
		if err != nil {
			return err
		}

		obj, err := jni.NewObject(env, class, jni.GetMethodID(env, class, "<init>", `()V`))
		if err != nil {
			return err
		}

		obj, err = jni.CallObjectMethod(env, obj, jni.GetMethodID(env, class, "getLanguage", "()Ljava/lang/String;"))
		if err != nil {
			return err
		}

		lang = jni.GoString(env, jni.String(obj))
		return nil
	})

	if err != nil {
		return ""
	}

	return lang
}
