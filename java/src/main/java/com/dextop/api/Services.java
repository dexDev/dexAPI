package com.dextop.api;

import org.json.JSONArray;
import org.json.JSONObject;
import java.util.HashMap;


public class Services {

    public String dextop;

    public Services() {
        this.dextop = "https://testnet271828.dex.top/v1";
    }

    public Services(String dextop) {
        this.dextop = dextop;
    }

    public String Login(String email, String password) {
        String url = this.dextop + Endpoints.Login;
        HashMap<String, String> loginData = new HashMap<String, String>();
        loginData.put("email", email);
        loginData.put("password", password);
        JSONObject ret = Http.FetchJson(url, loginData,"");
        return ret.getString("token");
    }

    public Models.MarketInfo GetMarketInfo() {
        String url = this.dextop + Endpoints.GetMarket;
        JSONObject ret = Http.FetchJson(url, "");
        System.out.println(ret.toString());
        Models.MarketInfo marketInfo = new Models.MarketInfo();
        marketInfo.marketAddr = ret.getString("marketAddr");
        JSONArray cashTokensJSON = ret.getJSONObject("config").getJSONArray("cashTokens");
        Models.Token[] cashTokens = new Models.Token[cashTokensJSON.length()];
        for(int i=0;i<cashTokens.length;i++) {
            JSONObject json = cashTokensJSON.getJSONObject(i);
            Models.Token token = new Models.Token();
            token.tokenCode = json.getInt("tokenCode");
            token.tokenId = json.getString("tokenId");
            cashTokens[i] = token;
        }
        marketInfo.cashTokens = cashTokens;

        JSONArray stockTokensJSON = ret.getJSONObject("config").getJSONArray("stockTokens");
        Models.Token[] stockTokens= new Models.Token[stockTokensJSON.length()];
        for(int i=0;i<stockTokens.length;i++) {
            JSONObject json = stockTokensJSON.getJSONObject(i);
            Models.Token token = new Models.Token();
            token.tokenCode = json.getInt("tokenCode");
            token.tokenId = json.getString("tokenId");
            stockTokens[i] = token;
        }
        marketInfo.stockTokens = stockTokens;
        return marketInfo;
    }

    public JSONObject PlaceOrder(Models.DexOrder order, String privKey, String token) {
        String url = this.dextop + Endpoints.PlaceOrder;
        HashMap<String,String> data = order.ToBody(privKey);
        JSONObject ret = Http.FetchJson(url, data,token);
        return ret;
    }
}
