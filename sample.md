## Trading API Methods

### Getting Public Trading Data

#### Method `GetPairDepth`

Get the depth data of a trading token pair.

**Request** `GET /api/v1/depth/{pairId}/{size}`

- `pairId`: The id of the trading token pair.
- `size`: The number of levels of depth to get.

**Sample Request**

```http
http://alpha.dex.top/api/v1/depth/ETH_ADX/3
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
      { price: 1050.98, amount: 50.50 }
      { price: 1100.00, amount: 10.70 },
      { price: 1200.37, amount: 80.25 },
    ],
    # `bids` are descending by price.
    "bids": [
      { price: 1000.00, amount: 10.00 },
      { price: 999.99, amount: 95.50 },
      { price: 983.80, amount: 15.32 }
    ]
  }
}
```
