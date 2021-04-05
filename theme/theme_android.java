package com.inkeliz.giopref.theme;

import android.content.res.Resources;
import android.content.res.Configuration;
import android.provider.Settings.Global;
import android.os.Build;

public class theme_android {
	public boolean isDark() {
	    return (Resources.getSystem().getConfiguration().uiMode & Configuration.UI_MODE_NIGHT_MASK) == Configuration.UI_MODE_NIGHT_YES;
	}

	public boolean isReducedMotion() {
	    return Global.TRANSITION_ANIMATION_SCALE == "0";
	}
}