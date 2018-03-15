## Trading API Methods

### Getting Public Trading Data

#### GetPairDepth

Get the depth data of a trading token pair.

**Request** `GET /v1/depth/{pairId}/{size}`

- `pairId` The id of the trading token pair.
- `size` The number of levels of depth to get.

**Sample Request**

```http
http://alpha.dex.top/v1/depth/ETH_ADX/3
```

**Sample Response**
```yaml
{
  "depth": {
    "pairId": "ETH_ADX",
    # The timestamp in millisecond of this depth data.
    "timeMs": "1513244690782",
    # `asks` are ascending by price.
    "asks": [
      { price: "1050.98", amount: "50.5" }，
      { price: "1100", amount: "10.70" },
      { price: "1200.37", amount: "80.25" }
    ],
    # `bids` are descending by price.
    "bids": [
      { price: "1000", amount: "10" },
      { price: "999.99", amount: "95.50" },
      { price: "983.8", amount: "15.32" }
    ]
  }
}
```



#### PlaceOrder

Place a trade order on current market

**Request** `POST /api/v1/placeorder` 

- cookie
  - `dex-user-jwt-token` User token get after login
  - `dex-user_id` Unique user id
  - `dex-trader-addr` The eth address legal in list
  - `dex-user_email` The user email encode

- params
  - `amount` Order amount
  - `price` Order price
  - `action` Action Enum 1. Buy, 2, Sell
  - `nonce`
  - `expireTimeSec` Order expire Time
  - `sig` Order infomation sign by private key of dex-trader-addr
  - `pairId` Order trade pair

**Sample Cookie**
```
dex-user-jwt-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ; 

dex-user_id=2; 

dex-user_email=flynn%40dex.top; 

dex-trader-addr=0x6a83D834951F29924559B8146D11a70EaB8E328b
```


**Sample Request**

```js
{
  "amount": "1",   
  "price": "0.0001",  
  "action": "buy",  
  "nonce": 1521012225317,  
  "expiretimesec": 1521098625,   
  "sig": "0xfaf909e20220ba3085f07616417fe17ecc3e82097e2b760a48ea0a24332899430632c361c645734118f5a8dca241b4c0eae41d4e3f88f789da2179121ac1107400",  
  "pairid": "eth_eos" 
}
```

**Sample Response**

```yaml
{
  "order": {
    "orderId": "10000001",  
    "pairId": "ETH_EOS",
    "action": "Buy",
    "price": "0.00010000",
    "amountTotal": "1.00000000",
    "amountFilled": "0.00000000",
    "filledAveragePrice": "NaN", // When PartiallyFilled cal the average price
    "status": "Unfilled", // Order Status 1.Filled, 2.Unfilled, 3.PartiallyFilled
    "createTimeMs": "1521012371171", 
    "updateTimeMs": "1521012371171",
    "nonce": "1521012225317"
  }
}
```

#### GetBalance

Get user's all token balance information, this api request limit 5 times per second

**Request** `GET /api/v1/balance` 

- cookie
  - `dex-user-jwt-token` User token get after login
  - `dex-user_id` Unique user id
  - `dex-trader-addr` The eth address legal in list
  - `dex-user_email` The user email encode

- params
  - none

**Sample Cookie**
```
dex-user-jwt-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ; 

dex-user_id=2; 

dex-user_email=flynn%40dex.top; 

dex-trader-addr=0x6a83D834951F29924559B8146D11a70EaB8E328b
```


**Sample Request**

```http
http://alpha.dex.top/api/v1/balances
```

**Sample Response**

```yaml
{
  "balances": [
    {
      "tokenId": "ADX", // TokenId
      "total": "0", // Token total amount
      "active": "0", // Token amount can withdraw
      "locked": "0" // Token amount is locked at this moment
    },
    {
      "tokenId": "EOS",
      "total": "0",
      "active": "0",
      "locked": "0"
    },
    {
      "tokenId": "ETH",
      "total": "1.00000000",
      "active": "0.99790000",
      "locked": "0.00210000"
    }
  ],
  "estimatedValue": "" // Account total estimate value
}
```