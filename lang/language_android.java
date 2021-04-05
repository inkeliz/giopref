package com.inkeliz.giopref.language;

import android.content.res.Resources;
import android.os.Build;
import java.util.Locale;

public class language_android {
	public String getLanguage() {
	    if (Build.VERSION.SDK_INT <= Build.VERSION_CODES.LOLLIPOP) {
	        return "";
	    }

        if (Build.VERSION.SDK_INT <= Build.VERSION_CODES.N) {
           return Resources.getSystem().getConfiguration().locale.toLanguageTag();
        }

        return Resources.getSystem().getConfiguration().getLocales().get(0).toLanguageTag();
	}
}