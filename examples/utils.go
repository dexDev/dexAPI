package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"strings"
	"encoding/binary"
	"github.com/dexDev/dexAPI/examples/models"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

func httpRequest(method string, url string, params map[string]string, auth bool) ([]byte, error) {
	client := &http.Client{}
	var reqBody string
	if params != nil {
		paramsBytes, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		reqBody = string(paramsBytes)
	}

	req, err := http.NewRequest(method, url, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	if auth && authenticateToken != "" {
		req.Header.Add("Authorization", "Bearer "+authenticateToken)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// The returned bytes are to be hashed (using keccak256) for signing. The content are the
// concatenation of the following (uints are in big-endian byte order):
//
// 1. String "DEx2 Order: "(96)  (Note the trailing whitespace)
// 2. <market address>(160)      (For preventing cross-market replay attack)
// 3. <nonce>(64) <expireTimeSec>(64) <amountE8>(64) <priceE8>(64) <ioc>(8) <action>(8) <pairId>(32)
func getOrderBytesToSign(marketAddr common.Address, nonce uint64, order *models.Order) []byte {
	bs := make([]byte, 0, 98)
	bs = append(bs, []byte("\x19Ethereum Signed Message:\n70")...)
	bs = append(bs, []byte("DEx2 Order: ")...)
	bs = append(bs, marketAddr.Bytes()...)
	bs = append(bs, uint64ToBigEndianBytes(nonce)...)
	bs = append(bs, uint64ToBigEndianBytes(order.ExpireTimeSec)...)
	bs = append(bs, uint64ToBigEndianBytes(order.AmountE8)...)
	bs = append(bs, uint64ToBigEndianBytes(order.PriceE8)...)
	bs = append(bs, order.Ioc)
	bs = append(bs, order.Action)
	bs = append(bs, uint32ToBigEndianBytes(order.PairId)...)

	if len(bs) != 98 { // 784 bits
		panic("The byte length of signing an order must be 98")
	}
	return bs
}

func uint64ToBigEndianBytes(value uint64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, value)
	return bs
}

func uint32ToBigEndianBytes(value uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, value)
	return bs
}


func StringToUint64E8(x string) uint64 {
	r := new(big.Rat)
	if _, err := fmt.Sscan(x, r); err != nil {
		return 0
	}
	r.Mul(r, big.NewRat(1e8, 1))
	if !r.IsInt() {
		return 0
	}

	val := r.Num()
	if !val.IsUint64() {
		return 0
	}
	return val.Uint64()
}

func GetPairCode(pairId string) uint32 {
	tokens := strings.Split(pairId, "_")
	cashCode := TokenCodeBySymbol[tokens[0]]
	stockCode := TokenCodeBySymbol[tokens[1]]

	return (uint32(cashCode) << 16) | uint32(stockCode)
}