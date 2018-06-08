This file gives 4 full data samples of order signing. You can use them to test and
debug your implementation.

## Market Config

The samples use this fake market config:
```
market_addr: "0x4afF2E056D2fE5CeBed442Dc924eAeBe25909550"
config: <
  cash_tokens: <
    token_id: "ETH"
    scale_factor: 1000000000000000000
  >
  stock_tokens: <
    token_id: "ADX"
    token_code: 100
    token_addr: "0xE154B8AE3aa072C3376C6162D28d31800273Ed0a"
    scale_factor: 1
  >
>
```

Get the real market config of the mainnet environment at https://dex.top/v1/market .

## Sample 1

### Order parameter

```yaml
pairId: ETH_ADX
action: buy
action: 0.1
amount: 15
expireTimeSec: 1527323455
nonce: 1527319855888
marketAddr: 0x4afF2E056D2fE5CeBed442Dc924eAeBe25909550
```

### Encoded order

```yaml
PairId: 100,
Action: 0,
Ioc: 0,                   # not used, always zero
PriceE8: 10000000,        # 0.1e8
AmountE8: 1500000000,     # 15e8
ExpireTimeSec: 1527323455
```

### Signature Information

Bytes to sign:
`0x19457468657265756d205369676e6564204d6573736167653a0a373044457832204f726465723a204aff2e056d2fe5cebed442dc924eaebe25909550000001639b5b8310000000005b091b3f0000000059682f000000000000989680000000000064`

Keccak256 hash the bytes to Sign:
`0x63b7204c7b80b2872d2e0af67f81ba4ca04d86d4e4a8f9cdbdbede1474656cfe`

Trader private key:
`0xd5b39ee354b4e06f0829b55ce2af3a49693126593e6ed8c8f6c9e1f8a1781a0d`

Trader public key:
`0x0479858808f1fcf2782eda27993b7759078a06db51f48cb80865b73a1b12eba32ed53bfdeb33e09ca8dc1662e4f2cb642dc38bb745c09e7f2afb81f614847c0f02` (not used, just FYI)

Trader address:
`0x5EA802D03039DaAA53E32f457D9CD42dac0BDDf9`

Signature:
`0x16944b324f3a9b3b14d5eaee5a96b55b68da2eaf09b831f177ce98374b4e098f5729f3e774b7331422e6b40f420d588df37fd2f8279f80d3d37eaca83bdf26d600`
