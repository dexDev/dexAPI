
### TradeInfo Public APIs

#### GetPairsByCash

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

#### GetPairInfo

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

#### GetTradeHistory

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


#### GetPairDepth


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
      {
        "price": "0.00140912",
        "amount": "753.35833123"
      }
    ]
  }
}
```


#### GetActiveOrders

#### GetPastOrders

