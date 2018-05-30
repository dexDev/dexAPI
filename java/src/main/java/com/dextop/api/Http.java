package com.dextop.api;

import com.github.kevinsawicki.http.HttpRequest;
import org.json.JSONObject;

import java.util.HashMap;

public class Http {
    public static JSONObject FetchJson(String url, String token) {
        HttpRequest request = HttpRequest.get(url);
        try {
            request.header("Authorization", "Bearer " + token);
            String response = request.body();
            JSONObject ret = new JSONObject(response);
            return ret;
        } catch (HttpRequest.HttpRequestException e) {
            System.out.printf("failed to call %s\n", request.toString());
            return null;
        }
    }

    public static JSONObject FetchJson(String url, HashMap<String, String> params, String token) {
        HttpRequest request = HttpRequest.post(url);
        try {
            request.header("Authorization", "Bearer " + token);
            JSONObject postData = new JSONObject(params);
            request.send(postData.toString());
            System.out.println(postData.toString());
            String response = request.body();
            JSONObject ret = new JSONObject(response);
            return ret;
        } catch (HttpRequest.HttpRequestException e) {
            System.out.printf("failed to call %s\n", request.toString());
            return null;
        }
    }
}