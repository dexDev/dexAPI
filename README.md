# DEx Open API
---

# Preparations

1. register an account by email on [https://dex.top](https://dex.top)
2. bind the trader address on the account page.
3. deposit eth or tokens on the balance page.
4. You can use [https://testnet271828.dex.top](https://testnet271828.dex.top) for test which base on Kovan testnet, You can contact admin on telegram ask for test tokens
5. APIs rate limit is 100/sec per ip or account

# Trade APIs

**Note**

API rate limit per IP address is `100/second`, and per user is `600/minute`.

## PlaceOrder

Place a new order.

**Request** `POST /v1/placeorder`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

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
    "status": "Unfilled", // "Filled" or "Unfilled" or "PartiallyFilled" or "Cancelled" or "Expired"
    "createTimeMs": "1522554774421",
    "updateTimeMs": "1522554774421",
    "expireTimeSec": "1522641169",
    "nonce": "1522554769144"
  }
}
```

### Note

- Full data samples of order signing, including the used private keys, can be found at
[here](https://github.com/dexDev/dexAPI/blob/master/samples/signing_orders.md).

- Signing Scheme 1 (Friendly to API usage)

  The bytes to be hashed (using keccak256) for signing are the concatenation of the following (uints are in big-endian order):
  1. Prefix `"\x19Ethereum Signed Message:\n70"`.
  2. String `"DEx2 Order: "` (Note the trailing whitespace)
  3. The `market address`.
  
     This is for replay attack protection when we deploy a new market.
  4. `nonce(64)`
  5. `expireTimeSec(64) amountE8(64) priceE8(64) 0x00(8) action(8) pairId(32)`


## CancelOrder

Cancel an existing order.

**Request** `POST /v1/cancelorder`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

- params
  - `traderAddr` Trader address
  - `orderId` The order id you want to cancel
  - `pairId` The id of the trading token pair
  - `nonce` Order nonce. The recommended value is the current timestamp in millisecond.

**Sample Request**

```js
{
  "orderId":"10000004",
  "pairId":"ETH_ADX",
  "nonce":1522554949196,  // nonce use timestamp in millisecond
}
```

**Sample response**

```js
{}
```

## CancelAllOrders

Cancel all existing orders using given parameters.

**Request** `POST /v1/cancelallorders`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

- params
  - `traderAddr` Trader address
  - `pairId` The id of the trading token pair
  - `nonce` Order nonce. The recommended value is the current timestamp in millisecond.

**Sample Request**

```js
{
  "pairId":"ETH_ADX",
  "nonce":1522554949196,  // nonce use timestamp in millisecond
}
```

**Sample response**

```js
{
    cancelledOrderIds: [1000373, 1000374] // ids of all cancelled orders
}
```

## Withdraw

Withdraw a token with the specified amount.

**Request** `POST /v1/withdraw`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

- required params
  - `tokenId`:  Withdraw token id
  - `amount`: Withdraw Amount

**Sample Request**

```js
{
  "tokenId":"ETH",
  "amount":"0.1",
}
```

**Sample response**

```js
{}
```



# TradeInfo Public APIs

## GetMarketInfo

Get relevant market information such as market contract address and token codes for signature.

**Request** `GET /v1/market`


**Sample Request**

```http
http://dex.top/v1/market
```

**Sample Response**

```js
{
  "marketAddr": "0x91952Ce7d434E84E8932AAB0F8A3f2E6A72F89a1", // Market Address for signature and deposit
  "label": "NoneLabel",
  "config": {
    "makerFeeRateE4": "0",
    "takerFeeRateE4": "0",
    "withdrawFeeRateE4": "0",
    "cashTokens": [
      {
        "tokenId": "ETH",
        "tokenCode": 0,  // Cash Token Code for signature
        "tokenAddr": "",
        "scaleFactor": "1000000000000000000"
      }
    ],
    "stockTokens": [
      {
        "tokenId": "BTM",
        "tokenCode": 300, // Stock Token Code for signature
        "tokenAddr": "0x45ddA1360019D2D26Bfa482A629a03f31A3a1A37",
        "scaleFactor": "1000"
      },
      {
        "tokenId": "OMG",
        "tokenCode": 400,
        "tokenAddr": "0x693f7fD5A8153A56bEc70189a1ae1943E07Be495",
        "scaleFactor": "1000"
      }
    ]
  },
  "networkName": ""
}
```


## GetPairsByCash

Get the real-time trading information of all available trading pairs of the specified cash token (e.g. "ETH").

**Request** `GET /v1/pairlist/:cashTokenId`

- `cashTokenId` The basic cash token like ETH

**Sample Request**

```http
http://dex.top/v1/pairlist/ETH
```

**Sample Response**

```js
{
  "pairs":[
    {
      "pairId": "ETH_ADX", // request trade pair's Id
      "timeMs": "1517975573850", // response timestamp
      "lastPrice": "5.095449", // price of timeMs
      "volume24": "414.056370", // total volume of this pair in 24 hours
      "change24": "4.421090", // price change of this pair in 24 hours
      "changePercent24": "", // price change rate of this pair in 24 hours
      "high24": "9.612255", // highest price of this pair in last 24 hours
      "low24": "3.170465" // lowest price of this pair in 24 hours
    }
  ],
}
```

## GetPairInfo

Get the real-time information of a trading pair.

**Request** `GET /v1/pairinfo/:pairId`

- `pairId` Order trade pair

**Sample Request**

```http
http://dex.top/v1/pairinfo/ETH_ADX
```

**Sample Response**

```js
{
  "pairId": "ETH_ADX",
  "timeMs": "1517974573647",
  "lastPrice": "5.813280",
  "volume24": "402.666631",
  "change24": "4.497973",
  "changePercent24": "",
  "high24": "7.978321",
  "low24": "3.491216"
}
```

## GetTradeHistory

Get recent trades by pairId sort by trade time.

**Request** `GET /v1/tradehistory/:pairId/:size`

- `pairId` The id of the trading token pair.
- `size` The number of levels of history to get.

**Sample Request**

```http
http://dex.top/v1/tradehistory/ETH_ADX/3
```

**Sample Response**

```js
{
  "records": [
    {
      "pairId": "ETH_ADX",
      "timeMs": "1522208665707", // The timestamp in millisecond of this trade.
      "action": "Buy",
      "price": "0.00010000",
      "amount": "934.00000000"
    },
    {
      "pairId": "ETH_ADX",
      "timeMs": "1522208545271",
      "action": "Buy",
      "price": "0.00010000",
      "amount": "66.00000000"
    },
    {
      "pairId": "ETH_ADX",
      "timeMs": "1522208486796",
      "action": "Buy",
      "price": "0.00010000",
      "amount": "1000.00000000"
    }
  ]
}
```


## GetPairDepth

Get the depth data of a trading pair.

**Request** `GET /v1/depth/:pairId/:size`

- `pairId` The id of the trading token pair.
- `size` The number of levels of depth to get.

**Sample Request**

```http
http://dex.top/v1/depth/ETH_ADX/5
```

**Sample Response**:

```js
{
  "depth": {
    "pairId": "ETH_ADX",
    "timeMs": "1513244690782", // The timestamp in millisecond of this depth data.
    // `asks` are ascending by price.
    "asks": [
      {
        "price": "0.00157384",
        "amount": "66.00000000"
      },
      {
        "price": "0.00159438",
        "amount": "80.00000000"
      },
      {
        "price": "0.00163499",
        "amount": "100.00000000"
      },
      {
        "price": "0.00168030",
        "amount": "666.00000000"
      },
      {
        "price": "0.00175212",
        "amount": "800.00000000"
      }
    ],
    // `bids` are descending by price.
    "bids": [
      {
        "price": "0.00155000",
        "amount": "665.00000000"
      },
      {
        "price": "0.00151903",
        "amount": "66.00000000"
      },
      {
        "price": "0.00150021",
        "amount": "1000.00000000"
      },
      {
        "price": "0.00144436",
        "amount": "78.00000000"
      },
        "price": "0.00140912",
        "amount": "753.35833123"
      }
    ]
  }
}
```


# TradeInfo User APIs

## GetActiveOrders

Get unfilled or partially filled orders that have not been cancelled or expired of a trader.

**Request** `GET /v1/activeorders/:addr/:pairId/:size/:page`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

- params
  - `pairId` The id of the trading token pair.
  - `size` The number of each page of user active orders to get.
  - `page` The pages number to get
  - `addr` Trader eth address

**Sample Request**

```js
http://dex.top/v1/activeorders/0x6a83D834951F29924559B8146D11a70EaB8E328b/ETH_ADX/100/1
```

**Sample response**

```js
{
  "orders": [
    {
      "orderId": "10126257",
      "pairId": "ETH_ADX",
      "action": "Buy",
      "price": "0.00010000",
      "amountTotal": "500.00000000",
      "amountFilled": "0.00000000",
      "filledAveragePrice": "0.00000000",
      "status": "Unfilled", // "Filled" or "Unfilled" or "PartiallyFilled" or "Cancelled" or "Expired"
      "createTimeMs": "1528790525164",
      "updateTimeMs": "1528790525164",
      "expireTimeSec": "1529395319",
      "nonce": "1528790519061"
    }
  ],
  "page": 1,
  "total": 1
}
```

## GetPastOrders

Get past orders on current wallet address.

**Request** `GET /v1/pastorders/:addr/:pairId/:size/:page?from_time_sec=&to_time_sec=`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

- params
  - `pairId` The id of the trading token pair.
  - `size` The number of each page of user past orders to get.
  - `page` The pages number to get
  - `addr` Trader eth address

- optional params
  - `from_time_sec` Timestamp in sec to get past orders from INCLUSIVE
  - `to_time_sec` Timestamp in sec to get past orders until INCLUSIVE

**Sample Request**

```js
http://dex.top/v1/pastorders/0x6a83D834951F29924559B8146D11a70EaB8E328b/ETH_ADX/100/1?from_time_sec=1498793709&to_time_sec=1498794709
```

**Sample response**

```js
{
  "orders": [
    {
      "orderId": "10000470",
      "pairId": "ETH_ADX",
      "action": "Buy",
      "price": "0.00004100",
      "amountTotal": "1500.00000000",
      "amountFilled": "1500.00000000",
      "filledAveragePrice": "0.00003924",
      "status": "Filled",
      "createTimeMs": "1528112182218",
      "updateTimeMs": "1528112182218",
      "expireTimeSec": "1528716977",
      "nonce": "1528112177402"
    }
  ],
  "page": 1,
  "total": 1
}
```

## GetOrderById

Get the details of an order by order id.

**Request** `GET /v1/orderbyid/:traderAddr/:orderId`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

- params
  - `traderAddr` Trader eth address
  - `orderId` The id of the order.

**Sample Request**

```js
http://dex.top/v1/orderbyid/0x6a83D834951F29924559B8146D11a70EaB8E328b/10000005
```

**Sample response**

```js
{
  "order": {
    "orderId": "10000005",
    "pairId": "ETH_ADX",
    "action": "Sell",
    "price": "0.00165000",
    "amountTotal": "200.00000000",
    "amountFilled": "200.00000000",
    "filledAveragePrice": "0.00165000",
    "status": "Filled",
    "createTimeMs": "1524809137398",
    "updateTimeMs": "1524809653759",
    "expireTimeSec": "1525413935",
    "nonce": "1524809135488"
  }
}
```

## GetKlineHistory

**Request** `GET /v1/kline/history?symbol={pairID}&resolution={resolution}&from={startTimeStamp}&to={endTimeStamp}`

- params
   - `pairId` The id of the trading token pair.
   - `resolution` duration for each bar, only support `5`, `15`, `30`, `60`, `1D`, `1W`.
   - `startTimeStamp` the start time of history data.
   - `endTimeStamp` the end time of history data.

**Sample Request**

```js
http://dex.top/v1/kline/history?symbol=ETH_YEE&resolution=60&from=1527807437&to=1527843437
```

**Sample response**

```js
{
    s: "ok",
    errmsg: "",
    // Timestamp
    t: [
        "1527825600",
        "1527829200",
        "1527832800",
        "1527836400",
        "1527840000"
    ],
    // Close
    c: [
        0.00004501,
        0.00004324,
        0.00004311,
        0.0000437,
        0.00004211
    ],
    // Open
    o: [
        0.00004392,
        0.00004501,
        0.00004324,
        0.00004311,
        0.0000437
    ],
    // High
    h: [
        0.00004504,
        0.00004401,
        0.00004376,
        0.0000437,
        0.0000431
    ],
    // Low
    l: [
        0.00004392,
        0.00004324,
        0.00004311,
        0.0000437,
        0.00004211
    ],
    // Volume
    v: [
        0.9682618499999999,
        73.98493420589,
        3.018485526814107,
        0.0809761,
        143.76744365959996
    ]
}
```

## GetTrades

Get recent trades on current wallet address.

**Request** `GET /v1/trades/:addr/:pairId/:size`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

- params
  - `pairId` The id of the trading token pair.
  - `size` The number of recent trades to get.
  - `addr` Trader eth address

**Sample Request**

```js
http://dex.top/v1/trades/0x6a83D834951F29924559B8146D11a70EaB8E328b/ETH_ADX/2
```

**Sample response**

```js
{
  "trades": [
    {
      "pairId": "ETH_ADX",
      "timeMs": "1525231321804",
      "orderId": "10000776",
      "isBuyer": true,
      "isMaker": false,
      "price": "0.00156080",
      "amount": "163.06101648",
      "fee": "0.32612203"
    },
    {
      "pairId": "ETH_ADX",
      "timeMs": "1525231321804",
      "orderId": "10000776",
      "isBuyer": true,
      "isMaker": false,
      "price": "0.00155810",
      "amount": "242.11581591",
      "fee": "0.48423163"
    }
  ]
}

```

# Account APIs

## Login

API used to login to exchange.

**Request** `POST /v1/authenticate`

**Sample Request**

```js
{
  "email": 'flynn@dex.top',
  "password": '**********',
}
```

Response:

```js
{
  "token": 'eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ',
  "user": {

  },
}
```

## Balance

Get the balances of all tokens of a trader.

**Request** `GET /v1/balances/:traderAddr`

- HTTP Request Header
  - `Authorization: Bearer <token>` Token obtained when signing in.

- params
  - `traderAddr` Trader address

**Sample Request**

```
http://dex.top/v1/balances
```

**Sample Response**

```js
{
  "balances": [
    {
      "tokenId": "ADX", // TokenId
      "total": "0", // Token total amount
      "active": "0", // Token amount can withdraw
      "locked": "0", // Token amount is locked at this moment
      "withdrawing": "0"
    },
    {
      "tokenId": "EOS",
      "total": "0",
      "active": "0",
      "locked": "0",
      "withdrawing": "0"
    },
    {
      "tokenId": "ETH",
      "total": "1.00000000",
      "active": "0.99790000",
      "locked": "0.00210000",
      "withdrawing": "0"
    }
  ],
  "estimatedValue": "" // Account total estimate value
}
```
