# DEx Open API
------

# Trade APIs

## PlaceOrder

Place an order .

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

## CancelOrder

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

## Withdraw

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



# TradeInfo Public APIs

## GetPairsByCash

This API for getting realtime  trade infomation of  all pair of specific cashtoken such as ETH. This API data can get all tradepair realtime trade infomation.

**Request** `GET /v1/pairlist/:cashTokenId` 

- `cashTokenId` The basic cash token like ETH

**Simple Request**

```http
http://alpha.dex.top/v1/pairlist/ETH
```

**simple Response**

```js
{
  "pairs":[
    {
      "pairId": "ETH_ADX", // request trade pair's Id
      "timeMs": "1517975573850", // reponse timestamp
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

Simple TradePair realtime data API by pairId

**Request** `GET /v1/pairinfo/:pairId`

- `pairId` Order trade pair

**Simple Request**

```http
http://alpha.dex.top/v1/pairlist/ETH_ADX
```

**Simple Response**
​
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

Get recent trades by PairId sort by trade time

**Request** `GET /v1/tradehistory/:pairId/:size`

- `pairId` The id of the trading token pair.
- `size` The number of levels of history to get.

**Simple Request**

```http
http://alpha.dex.top/v1/tradehistory/ETH_ADX/3
```

**Simple Response**

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


**Request** `GET /v1/depth/:pairId/:size` 

- `pairId` The id of the trading token pair.
- `size` The number of levels of depth to get.

**Simple Request**

```http
http://alpha.dex.top/v1/depth/ETH_ADX/5
```

**Simple Reponse**:

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
      s        "price": "0.00140912",
        "amount": "753.35833123"
      }
    ]
  }
}
```


# TradeInfo User APIs

## GetActiveOrders

## GetPastOrders


### Account APIs

#### Login

#### Balance