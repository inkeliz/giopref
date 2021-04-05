package com.inkeliz.giopref.power;

import android.os.Build;
import android.content.Context;
import android.os.BatteryManager;
import android.os.PowerManager;

public class power_android {
    public int batteryLevel(Context context) {
        if (Build.VERSION.SDK_INT <= Build.VERSION_CODES.LOLLIPOP) {
            return 100;
        }

        return ((BatteryManager)context.getSystemService(Context.BATTERY_SERVICE)).getIntProperty(BatteryManager.BATTERY_PROPERTY_CAPACITY);
    }

    public boolean isCharging(Context context) {
        if (Build.VERSION.SDK_INT <= Build.VERSION_CODES.LOLLIPOP) {
            return true;
        }

        return ((BatteryManager)context.getSystemService(Context.BATTERY_SERVICE)).isCharging();
    }

    public boolean isSaving(Context context) {
        if (Build.VERSION.SDK_INT <= Build.VERSION_CODES.LOLLIPOP) {
            return true;
        }

        return ((PowerManager)context.getSystemService(Context.POWER_SERVICE)).isPowerSaveMode();
    }
}