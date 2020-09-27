package com.test.auto.util;

import com.test.auto.model.InterfaceName;

import java.util.Locale;
import java.util.ResourceBundle;

public class ConfigFile {
    private static ResourceBundle bundle = ResourceBundle.getBundle("application", Locale.CANADA);

    public static String getUrl(InterfaceName name){
        String address = bundle.getString("test.url");
        String uri = "";
        String testUrl;
        if (name == InterfaceName.HTTPREQUESTFORM){
            uri = bundle.getString("getHttpRequest.uri");
        }
        testUrl = address+uri;
        return testUrl;
    }
}
