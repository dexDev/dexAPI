# DEx Open API

------

### TradeInfo APIs

#### GetPairsByCash 货币对列表

URL: `/api/v1/pairlist/:cashTokenId` 

Method: `GET` 

Request: 

```http
http://alpha.dex.top/api/v1/pairlist/ETH
```

Response:

​	status code: 200

```javascript
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

Note: 

This API for getting realtime  trade infomation of  all pair of specific cashtoken such as ETH.



#### GetPairInfo 货币对详情

URL: `/api/v1/pairinfo/:pairId`

Method: `GET` 

Request: 

```http
http://alpha.dex.top/api/v1/pairlist/ETH_ADX
```

Response:

​	status code: 200

```javascript
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

Note:



#### GetTokens 查询系统支持的所有币种

URL: `/api/v1/tokens` 

Method: `GET`

Request: 

```http
http://alpha.dex.top/api/v1/tokens
```

Response:

```javascript
{
  "data": [
    "ETH",
    "ADX",
    "EOS",
  ],
  total: 1
}
```

Note:



#### GetTradeHistory 货币对最近交易

Method:  `GET` 

URL: `/api/v1/tradehistory/:pairId/:size`

Request: 

```http
http://alpha.dex.top/api/v1/tradehistory/ETH_ADX/5
```

Response:

```javascript
{
  "records": [{
     "pairId": "ETH_ADX",
     "timeMs": "1513244690782",
     "action": "Sell", // Buy, Sell
     "price": "0.11",
     "amount": "253"
  }]
}
```



#### GetPairDepth 货币对深度

Method:  `GET` 

URL: `/api/v1/depth/:pairId/:size` 

Request: 

```http
http://alpha.dex.top/api/v1/depth/ETH_ADX/5
```

Reponse:

```javascript
{
  "depth": {
    "pairId": "ETH_ADX",
    "timeMs": "1513244690782",
    "asks": [
      {
        price: 1000.00 // price升序
        amount: 100
      }
    ],
    "bids": [
      {
        price: 1000.00// 按price降序
        amount: 100 
      }
    ]
  }
}
```



#### GetKlineByPairId 货币对图表数据

Method: `GET` 

URL: `/api/v1/kline/history`

Response: 

```javascript
{
    
}
```

#### GetActiveOrders 查询当前订单

Methed: `GET`

URL:  `/api/v1/activeorders/:traderAddr/:pairId/:size/:page` 

Request: 

```javascript
http://alpha.dex.top/api/v1/activeorders/ETH_ADX/100/1
```

response:

```javascript
{
  orders: [{
    order_id: '334213',
    pair_id: 'ETH_ADX',
    action: 1,
    type: 'limit',
    price: 1000,
    amount_total: 6, //总数量（包括成交、未成交）
    amount_filled: 3, // 已成交数量
    filled_total_price: 3000, // 已成交总金额
    create_time_ms: 12317, // 创建时间
    update_time_ms: 12317, // 结束时间
    status: 1,
    nonce: 12,
  }],
  total: 1,
  page: 1
}
```



#### GetPastOrders 查询历史订单

Method:  `POST` 

URL: `/api/v1/pastorders/:pairId/:size/:page`

Request:  **sign needed**

```javascript
{
  filter: { // optional，不加filter查询所有
    action: '1', // all for all action
    type: 'limit', // limit, market, and all for all type
  },
}
```

response:

```javascript
{
  orders: [{
    order_id: '334213',
    pair_id: 'ETH_ADX',
    action: 1,
    type: 'limit',
    price: 1000,
    amount_total: 6, //总数量（包括成交、未成交）
    amount_filled: 3, // 已成交数量
    filled_total_price: 3000, // 已成交总金额
    create_time_ms: 12317, // 创建时间
    update_time_ms: 12317, // 结束时间
    status: 1,
    nonce: 12,
  }],
  total: 1,
  page: 1
}

```



------

### Trade APIs

#### PlaceOrder 下单

Method: `POST`

URL:  `/api/v1/placeorder` 

Cookie: 

```
dex-user-jwt-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ; 
// User token
dex-user_id=2; 

dex-user_email=flynn%40dex.top; 

dex-trader-addr=0x6a83D834951F29924559B8146D11a70EaB8E328b
// trader address of this operation
```



Request:  **sign needed**

```js
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

Response:

```javascript
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

#### CancelOrder 撤单

Method: `POST` 

URL: `/api/v1/cancelorder `

Request: 

```javascript
{
  orderId: '111',
  pairId: 'ETH_ADX',
  nonce: 1,
}
```

response:

```javsscript
{
}
```

------

### Account APIs

#### Login 登录（验证账户）

Method:  `POST` 

URL: `/v1/authenticate` 

Request:  

```javascript
{
  "email": '',
  "password": '',
}
```

Reponse:

```javascript
{
  "token": '',
  "user": {},
}
```

#### register 注册

Method: `POST` 

URL: `/v1/register`

Request:  

```javascript
{
  "email": "",
  "password": "",
}
```

Reponse:

```javascript
{
  "token": '',
}
```

#### Email Confirm 验证邮箱

Method: `GET`

URL `/v1/email/confirm` 

Request:  

```javascript
{
  "token": "",
}
```

Reponse:

```javascript
{
}
```

#### Email Resend 重发邮件

Method: `GET`

URL: `/v1/email/resend` 

Request:  

```javascript
{
  "email": "",
}
```

Reponse:

```javascript
{
}
```

#### Transfers 帐号充值提现历史

Method: `POST`

URL:  `/api/v1/transfers` 

Request:  **sign needed**

```javascript
{
  trans_type: 'all', // Trans_enum: 1
  token: 'all', // 币种类型
  page: 1,
  count: 20,
}
```

response:

```javascript
{
  data: [{
    trans_id: 1, // 转账id
    token_id: 'ETH',
    active: 1,//充值、提现
    amount: 6,
    addr: 'xxx', //充值提现地址
    txid: 'xxx', //Transaction ID
    fee: 11,
    status: 1, //充值提现对应状态
    create_time_ms: 111,
    update_time_ms: 111,
  },{
    trans_id: 2, // 转账id
    token_id: 'ADX',
    active: 2,//充值、提现
    amount: 6,
    addr: 'xxx', //充值提现地址
    txid: 'xxx', //Transaction ID
    fee: 11,
    status: 1, //充值提现对应状态
    contract_id: 12, // TODO
    create_time_ms: 111,
    update_time_ms: 111,
  }],
  total: 2,
  page: 1
}
```

#### Balance 账户资产

Method: `GET`

URL: `/api/v1/balances` 

Cookie: 

```
dex-user-jwt-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGRDbGFpbXMiOnsiYXVkIjoiREV4IFNlcnZlcnMiLCJleHAiOjE1MjEyNjU3MTcsImp0aSI6ImEwY2VmNGQ0LTBmNWQtNDVmMS05MzYwLTkwOGQ0MmY1ZWMzYiIsImlhdCI6MTUyMTAwNjUxNywiaXNzIjoiREV4Iiwic3ViIjoiMiJ9fQ.r0kwUhMh8pZCGazZt0Lp4gPl1JEOdQIGyXlNpi5zHQ90NloUXNuEhlSRvSrTu5rug6nhkO_cbvIGc2okeC9zLQ; 
// User token
dex-user_id=2; 

dex-user_email=flynn%40dex.top; 

dex-trader-addr=0x6a83D834951F29924559B8146D11a70EaB8E328b
// trader address of this operation
```

Request: 

```
http://alpha.dex.top/api/v1/balances
```

Response:

```javascript
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

#### Deposit 充值

Method: `POST`

URL: `/api/v1/deposit`

Request:  

```javascript
{
    // the address will sign 
    traderAddr: "0x6a83D834951F29924559B8146D11a70EaB8E328b",
    // token wants to trans
    tokenId: "ETH",
    // amount 
    amount: "10.3"
    
}
```

Reponse:

```
{
}
```

 Note: 此方法应用于Eth



#### Withdraw 提现

Method: `POST`

URL:  `/api/v1/withdraw` 

Request:  **sign needed**

```javascript
{
  tokenId: 'ETH',
  amount: '11.11', //提现数量
}
```

response:

```javascript
{
}
```