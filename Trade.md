
### Trade APIs

#### PlaceOrder

Place an order.

**Request** `POST /v1/placeorder`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when sign in.

- Params
  - `traderAddr` Trader address
  - `pairId` Id of the trading token pair
  - `amount` Amount of token to buy or sell
  - `price` Order price
  - `action` "Buy" or "Sell"
  - `nonce` Order nonce. Recommended value is the current timestamp in millisecond
  - `expireTimeSec` Order expire time (timestamp in second)
  - `sig` Signature of signing the order with the private key of the trader address, see API Introduction for details.

**Sample Request**

```js
{
  "amount": "66.51501244",
  "price": "0.00150342",
  "action": "Buy",
  "nonce": 1522290645732,
  "expireTimeSec": 1522377045,
  "sig": "0x3d42c49bebfc912d5b6698556db273e9e38b72928b052eccd967a4bb4d31cb007b363a1a30d45ee5055b64a8a7d8bac7fc1d4d2f3009d910b856467699a58b4f00",
  "pairId": "ETH_ADX"
}
```

**Sample Response**

```js
{
  "order": {
    "orderId": "10000004",
    "pairId": "ETH_ADX",
    "action": "Buy",
    "price": "0.01000000",
    "amountTotal": "10.00000000",
    "amountFilled": "0.00000000",
    "filledAveragePrice": "0.00000000",
    "status": "Unfilled", // "Filled" or "Unfilled" or "PartiallyFilled"
    "createTimeMs": "1522554774421",
    "updateTimeMs": "1522554774421",
    "expireTimeSec": "1522641169",
    "nonce": "1522554769144"
  }
}
```

#### CancelOrder

**Request**`POST /v1/cancelorder`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when sign in.

- params
  - `traderAddr` Trader address
  - `orderId` The order id you want to cancel
  - `pairId` The id of the trading token pair
  - `nonce` Order nonce. Recommended value is the current timestamp in millisecond.

**Simple Request**

```js
{
  "orderId":"10000004",
  "pairId":"ETH_ADX",
  "nonce":1522554949196,  // nonuce use timestamp in millisecond
}
```

**Simple response**

```js
{}
```

#### Withdraw

**Request** `POST /v1/withdraw`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when sign in.

- required params
  - `tokenId`:  Withdraw token id
  - `amount`: Withdraw Amount

**Simple Request**

```js
{
  "tokenId":"ETH",
  "amount":"0.1",
}
```

**Simple response**

```js
{}
```
