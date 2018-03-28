# DEx Open API
------

### TradeInfo APIs

#### GetPairsByCash

This API for getting realtime  trade infomation of  all pair of specific cashtoken such as ETH.

**Request** `GET /v1/pairlist/:cashTokenId` 

- `cashTokenId` The basic cash token like ETH

**Simple Request**

```http
http://alpha.dex.top/v1/pairlist/ETH
```

**simple Response**

```yaml
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

**Request** `GET /v1/pairinfo/:pairId`

- `pairId` Order trade pair

**Simple Request**

```http
http://alpha.dex.top/v1/pairlist/ETH_ADX
```

**Simple Response**
​
```yaml
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

#### GetTokens

**Request** `GET /v1/tokens`

**Simple Request**

```http
http://alpha.dex.top/v1/tokens
```

**Simple Response**

```yaml
{
  "data": [
    "ETH",
    "ADX",
    "EOS",
  ],
  total: 1
}
```


#### GetTradeHistory

Get recent trades by PairId

**Request** `GET /v1/tradehistory/:pairId/:size`

- `pairId` The id of the trading token pair.
- `size` The number of levels of history to get.

**Simple Request**

```http
http://alpha.dex.top/v1/tradehistory/ETH_ADX/3
```

**Simple Response**

```yaml
{
  "records": [
    {
      "pairId": "ETH_ADX",
      "timeMs": "1522208665707",
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

Reponse:

```yaml
{
  "depth": {
    "pairId": "ETH_ADX",
    # The timestamp in millisecond of this depth data.
    "timeMs": "1513244690782",
    # `asks` are ascending by price.
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
    # `bids` are descending by price.
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

Get user's unfilled orders on current wallet address

**Request** `GET /v1/activeorders/:pairId/:size/:page` 


- cookie
  - `dex-user-jwt-token` User token get after login
  - `dex-trader-addr` User's current legal trading address 

- params
  - `pairId` The id of the trading token pair.
  - `size` The number of each page of user active orders to get.
  - `page` The pages number to get

**Sample Cookie**
```
dex-user-jwt-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ; 

dex-trader-addr=0x6a83D834951F29924559B8146D11a70EaB8E328b
```

**Simple Request**

```yaml
http://alpha.dex.top/v1/activeorders/ETH_ADX/100/1
```

**Simple response**

```yaml
{
  orders: [{
    order_id: '334213', //Unique Order Id
    pair_id: 'ETH_ADX', 
    action: 1,
    type: 'limit', // Order Type:
    price: 1000,
    amount_total: 6, //Total Amount（include filled and unfilled）
    amount_filled: 3, // Filled amount
    filled_total_price: 3000, // Filled price
    create_time_ms: 12317, // Order create time
    update_time_ms: 12317, // Last updated time
    status: 1,
    nonce: 12,
  }],
  total: 1,
  page: 1
}
```



#### GetPastOrders

Get user's filled orders on current wallet address

**Request**`POST /v1/pastorders/:pairId/:size/:page`

- cookie
  - `dex-user-jwt-token` User token get after login API
  - `dex-trader-addr` User's current legal trading address 

- params
  - `pairId` The id of the trading token pair.
  - `size` The number of each page of user past orders to get.
  - `page` The pages number to get


**Sample Cookie**
```
dex-user-jwt-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ; 

dex-trader-addr=0x6a83D834951F29924559B8146D11a70EaB8E328b
```

**Simple Request**

```yaml
http://alpha.dex.top/v1/activeorders/ETH_ADX/100/1
```

**Simple response**

```yaml
{
  orders: [{
    order_id: '334213',
    pair_id: 'ETH_ADX',
    action: 1,
    type: 'limit', // Order type
    price: 1000,
    amount_total: 6, //Total Amount（include filled and unfilled）
    amount_filled: 3, // Filled amount
    filled_total_price: 3000, // Filled price
    create_time_ms: 12317, // Order create time
    update_time_ms: 12317, // Last updated time
    status: 1,
    nonce: 12,
  }],
  total: 1,
  page: 1
}

```



------

### Trade APIs

#### PlaceOrder

Place a trade order on current market using current trading address

**Request** `POST /v1/placeorder` 

- cookie
  - `dex-user-jwt-token` User token get after login API
  - `dex-trader-addr` User's current legal trading address 

- params
  - `amount` The id of the trading token pair.
  - `price` The number of each page of user past orders to get.
  - `action` The pages number to get
  - `nonce`
  - `sig` 
  - `pairId` The pairid 


**Simple Cookie**

```
dex-user-jwt-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ; 

dex-trader-addr=0x6a83D834951F29924559B8146D11a70EaB8E328b
// trader address of this operation
```

**Simple Request**

```js
http://alpha.dex.top/v1/placeorder
{
  "amount": "1",  // Order amount
  "price": "0.0001", // Order price
  "action": "Buy", // Action Enum 1. Buy, 2, Sell
  "nonce": 1521012225317, // 
  "expireTimeSec": 1521098625,  // Order expire Time
  "sig": "0xfaf909e20220ba3085f07616417fe17ecc3e82097e2b760a48ea0a24332899430632c361c645734118f5a8dca241b4c0eae41d4e3f88f789da2179121ac1107400", // Order infomation sign by private key of dex-trader-addr
  "pairId": "ETH_EOS" // Order trade pair
}

```

**Simple Response**

```yaml
{
  "order": {
    "orderId": "10000001", // 
    "pairId": "ETH_EOS",
    "action": "Buy",
    "price": "0.00010000",
    "amountTotal": "1.00000000",
    "amountFilled": "0.00000000",
    "filledAveragePrice": "NaN", // When PartiallyFilled cal the average price
    "status": "Unfilled", // Order Status 1.Filled, 2.Unfilled, 3.PartiallyFilled
    "createTimeMs": "1521012371171", //
    "updateTimeMs": "1521012371171",
    "nonce": "1521012225317"
  }
}
```

#### CancelOrder

**Request**`POST /v1/cancelorder`

- cookie
  - `dex-user-jwt-token` User token get after login API
  - `dex-trader-addr` User's current legal trading address 

- params
  - `orderId` The id of the trading token pair.
  - `pairId` The number of each page of user past orders to get.
  - `nonce`

**Simple Request**

```yaml
{
  orderId: '111',
  pairId: 'ETH_ADX',
  nonce: 1,
}
```

**Simple response**

```javsscript
{
}
```

------

### Account APIs

#### Login

**Request** `POST /v1/placeorder` 

- cookie
  - `dex-user-jwt-token` User token get after login API
  - `dex-trader-addr` User's current legal trading address 

- params
  - `amount` The id of the trading token pair.
  - `price` The number of each page of user past orders to get.
  - `action` The pages number to get
  - `nonce`
  - `sig` 
  - `pairId` The pairid 

**Request**`POST /v1/authenticate` 

**Simple Request**

```yaml
{
  "email": '',
  "password": '',
}
```

Reponse:

```yaml
{
  "token": '',
  "user": {},
}
```

#### register

**Request**`POST /v1/register`

**Simple Request**

```yaml
{
  "email": "",
  "password": "",
}
```

**Simple Reponse**

```yaml
{
  "token": '',
}
```

#### Email Confirm

**Request** `GET /v1/email/confirm` 

 **Simple Request**

```yaml
{
  "token": "",
}
```

Reponse:

```yaml
{
}
```

#### Email Resend

`GET`

**Request**`/v1/email/resend` 

 **Simple Request**

```yaml
{
  "email": "",
}
```

Reponse:

```yaml
{
}
```

#### Transfers

`POST`

**Request** `/v1/transfers` 

 **Simple Request**

```yaml
{
  trans_type: 'all', // Trans_enum: 1
  token: 'all', // Token type
  page: 1,
  count: 20,
}
```

**Simple response**

```yaml
{
  data: [{
    trans_id: 1, // Unique trans id
    token_id: 'ETH',
    active: 1,// Trans type: 1, deposit 2,withdraw
    amount: 6,
    addr: 'xxx', //User address of trans
    txid: 'xxx', //Transaction ID
    fee: 11, // Transaction fee
    status: 1, //Transaction status 
    create_time_ms: 111,
    update_time_ms: 111,
  },{
    trans_id: 2, 
    token_id: 'ADX',
    active: 2,
    amount: 6,
    addr: 'xxx', 
    txid: 'xxx', 
    fee: 11,
    status: 1, 
    contract_id: 12, 
    create_time_ms: 111,
    update_time_ms: 111,
  }],
  total: 2,
  page: 1
}
```

#### Balance

**Request** `GET /v1/balances` 

Cookie: 

```
dex-user-jwt-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ; 

dex-trader-addr=0x6a83D834951F29924559B8146D11a70EaB8E328b
// trader address of this operation
```

**Simple Request**

```
http://alpha.dex.top/v1/balances
```

**Simple Response**

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

------

### Funding API

#### Withdraw 

**Request** `POST /v1/withdraw` 

- tokenId:  Withdraw token id
- amount: Withdraw Amount

**Simple Request**

```yaml
{
  tokenId: 'ETH', 
  amount: '11.11',
}
```

**Simple response**

```yaml
{
}
```