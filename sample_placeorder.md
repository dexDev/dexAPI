#### PlaceOrder

Place an order.

**Request** `POST /v1/placeorder`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when sign in.

- params
  - `traderAddr` Trader address
  - `pairId` Id of the trading token pair
  - `amount` Amount of token to buy or sell
  - `price` Order price
  - `action` "Buy" or "Sell"
  - `nonce` Order nonce. Recommended value is the current timestamp in millisecond.
            See API Introduction for details.
  - `expireTimeSec` Order expire time (timestamp in second).
  - `sig` Signature of signing the order with the private key of the trader address.
          See API Introduction for details.

**Sample Request**

```js
{
  "amount": "1",
  "price": "0.0001",
  "action": "Buy",
  "nonce": 1521012225317,
  "expiretimesec": 1521098625,
  "sig": "0xfaf909e20220ba3085f07616417fe17ecc3e82097e2b760a48ea0a24332899430632c361c645734118f5a8dca241b4c0eae41d4e3f88f789da2179121ac1107400",
  "pairid": "ETH_ADX"
}
```

**Sample Response**

```js
{
  "order": {  // the placed order
    "orderId": "10000001",
    "pairId": "ETH_EOS",
    "action": "Buy",
    "price": "0.00010000",
    "amountTotal": "1.00000000",
    "amountFilled": "0.00000000",
    "filledAveragePrice": "NaN",  // the average price
    "status": "Unfilled",  // "Filled" or "Unfilled" or "PartiallyFilled"
    "createTimeMs": "1521012371171",
    "updateTimeMs": "1521012371171",
    "nonce": "1521012225317"
  }
}
```
