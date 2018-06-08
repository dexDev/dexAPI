This file gives several full data samples of order signing. You can use them to test and
debug your implementation.

See the specification of order signing and example `PlaceOrderRequest` at
https://github.com/dexDev/dexAPI#placeorder.

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

- Bytes to sign:
  `0x19457468657265756d205369676e6564204d6573736167653a0a373044457832204f726465723a204aff2e056d2fe5cebed442dc924eaebe25909550000001639b5b8310000000005b091b3f0000000059682f000000000000989680000000000064`

- Keccak256 hash the bytes to Sign:
  `0x63b7204c7b80b2872d2e0af67f81ba4ca04d86d4e4a8f9cdbdbede1474656cfe`

- Trader private key:
  `0xd5b39ee354b4e06f0829b55ce2af3a49693126593e6ed8c8f6c9e1f8a1781a0d`

- Trader public key:
  `0x0479858808f1fcf2782eda27993b7759078a06db51f48cb80865b73a1b12eba32ed53bfdeb33e09ca8dc1662e4f2cb642dc38bb745c09e7f2afb81f614847c0f02` (not used, just FYI)

- Trader address:
  `0x5EA802D03039DaAA53E32f457D9CD42dac0BDDf9`

- Signature:
  `0x16944b324f3a9b3b14d5eaee5a96b55b68da2eaf09b831f177ce98374b4e098f5729f3e774b7331422e6b40f420d588df37fd2f8279f80d3d37eaca83bdf26d600`

## Sample 2

### Order parameter

```yaml
pairId: ETH_ADX
action: sell
action: 0.1
amount: 10
expireTimeSec: 1527323456
nonce: 1527319856421
marketAddr: 0x4afF2E056D2fE5CeBed442Dc924eAeBe25909550
```

### Encoded order

```yaml
PairId: 100,
Action: 1,
Ioc: 0,                   # not used, always zero
PriceE8: 10000000,        # 0.1e8
AmountE8: 1000000000,     # 10e8
ExpireTimeSec: 1527323456
```

### Signature Information

- Bytes to sign:
  `0x19457468657265756d205369676e6564204d6573736167653a0a373044457832204f726465723a204aff2e056d2fe5cebed442dc924eaebe25909550000001639b5b8525000000005b091b40000000003b9aca000000000000989680000100000064`

- Keccak256 hash the bytes to Sign:
  `0x2d3959297cef26d2492e6436b5216bda19ab44ceea6f1b92550ae8d22ac421d1`

- Trader private key:
  `0xbc36103586deb12003a937a57288032d307a34b5382cac6ae55c109f079b3b85`

- Trader public key:
  `0x04c316d359d7ca589d89f31701d04eef8af372a488e11a80093172b035111839dfb6602a6faac62650b75355d4dbac216f7ccda8b08f038458dbd67af939235e26` (not used, just FYI)

- Trader address:
  `0x11C2e2970B1b12f30f343711967F97EA3acf115c`

- Signature:
  `0x8ab426ad00fcde9897eab436c361427a87b80b53e30f1fa8c77df46751fe8c2255137858a0f6e43018c9ae820ff635f5615cee2f3eb11768f3a17958f51edadf00`

## Sample 3

### Order parameter

```yaml
pairId: ETH_ADX
action: sell
action: 0.1
amount: 3
expireTimeSec: 1527323458
nonce: 1527319858944
marketAddr: 0x4afF2E056D2fE5CeBed442Dc924eAeBe25909550
```

### Encoded order

```yaml
PairId: 100,
Action: 1,
Ioc: 0,                   # not used, always zero
PriceE8: 10000000,        # 0.1e8
AmountE8: 300000000,      # 3e8
ExpireTimeSec: 1527323458
```

### Signature Information

- Bytes to sign:
  `0x19457468657265756d205369676e6564204d6573736167653a0a373044457832204f726465723a204aff2e056d2fe5cebed442dc924eaebe25909550000001639b5b8f00000000005b091b420000000011e1a3000000000000989680000100000064`

- Keccak256 hash the bytes to Sign:
  `0x7a41bbf4c52bbc60a0f34c305fa12dcea20b1c6c921d8f65c08fe78e08885049`

- Trader private key (the same to Sample 2):
  `0xbc36103586deb12003a937a57288032d307a34b5382cac6ae55c109f079b3b85`

- Trader public key:
  `0x04c316d359d7ca589d89f31701d04eef8af372a488e11a80093172b035111839dfb6602a6faac62650b75355d4dbac216f7ccda8b08f038458dbd67af939235e26` (not used, just FYI)

- Trader address:
  `0x11C2e2970B1b12f30f343711967F97EA3acf115c`

- Signature:
  `0x8c8e4cb390bba482efe2dfee28945778f6bfefcc834702c4bede84d938074d9f33305394f661a3527e9b77717157263b8fadc24bbb976ff9ae03c330d6029c1501`

## Sample 4

### Order parameter

```yaml
pairId: ETH_ADX
action: sell
action: 0.1
amount: 7
expireTimeSec: 1527323458
nonce: 1527319858952
marketAddr: 0x4afF2E056D2fE5CeBed442Dc924eAeBe25909550
```

### Encoded order

```yaml
PairId: 100,
Action: 1,
Ioc: 0,                   # not used, always zero
PriceE8: 10000000,        # 0.1e8
AmountE8: 700000000,      # 7e8
ExpireTimeSec: 1527323458
```

### Signature Information

- Bytes to sign:
  `0x19457468657265756d205369676e6564204d6573736167653a0a373044457832204f726465723a204aff2e056d2fe5cebed442dc924eaebe25909550000001639b5b8f08000000005b091b420000000029b927000000000000989680000100000064`

- Keccak256 hash the bytes to Sign:
  `0x485324dfaa0c279328703fda95392a362e282e5763491c4ef36bceacfe39cf60`

- Trader private key (the same to Sample 2):
  `0xbc36103586deb12003a937a57288032d307a34b5382cac6ae55c109f079b3b85`

- Trader public key:
  `0x04c316d359d7ca589d89f31701d04eef8af372a488e11a80093172b035111839dfb6602a6faac62650b75355d4dbac216f7ccda8b08f038458dbd67af939235e26` (not used, just FYI)

- Trader address:
  `0x11C2e2970B1b12f30f343711967F97EA3acf115c`

- Signature:
  `0x27bc2cff485a9a6a3c3b2d9df604bf1103b202975fd4b4d1e64de3508364ff1f333269e556334c97cc97bbf8502c04832f643ba682c2b01fbb02c672e45cafdd01`
