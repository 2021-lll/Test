package com.test.auto.config;

import lombok.Data;
import org.apache.http.client.CookieStore;

@Data
public class TestConfig {
    public static String httpRequestUrl;

    public static CookieStore store;
}
