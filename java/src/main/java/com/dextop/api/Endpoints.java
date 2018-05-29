package com.dextop.api;

public interface Endpoints {
    // Trade APIs
    String PlaceOrder = "/placeorder";
    String CancelOrder = "/cancelorder";
    String CancelAllOrders = "/cancelallorders";
    String Withdraw = "/withdraw";

    // Public APIs
    String GetMarket = "/market";
    String GetPairs = "/pairlist/%d";
    String GetPairInfo = "/pairinfo/%s";
    String GetTradeHistory = "/tradehistory/%d/%d";
    String GetPairDepth = "/depth/%d/%d";

    // User APIs
    String GetActiveOrders = "/activeorders/%s/%s/%d/%d";
    String GetPastOrders = "/pastorders/%s/%s/%d/%d";
    String GetOrderById = "/orderbyid/%d";
    String GetTrades = "/trades/%s/%d/%d";

    // Account APIs
    String Login = "/authenticate";
    String GetBalance = "/balances/%s";
}
