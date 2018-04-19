/** Account prepare

  1. register a account by email on https://kovan.dex.top
  2. binding trader address in account page.
  3. deposit eth or tokens in balance page.

**/
// orders example

package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dexDev/dexAPI/examples/models"
	// For sig
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var authenticateToken string
var marketAddr string
var TokenCodeBySymbol = map[string]uint16{}
var dextopHost = "https://kovan.dex.top/"
var testAccount = "flynn@dex.top"
var testAccountPwd = "12345678a"
var testTraderAddr = "0x6a83D834951F29924559B8146D11a70EaB8E328b"
var testTraderAddrPriKey = "121e1348709ca0f75ea8793bbac27886afe6eb272c9a5245890aa7e4c64a65b9"

// Http Request Util Fuction
func request(Method string, Url string, Params map[string]string, Auth bool) string {
	httpClient := &http.Client{}

	jsonParams := ""
	if nil != Params {
		bytesParams, _ := json.Marshal(Params)
		jsonParams = string(bytesParams)
	}

	request, err := http.NewRequest(Method, Url, strings.NewReader(jsonParams))
	if nil != err {
		return err.Error()
	}

	if Auth && authenticateToken != "" {
		request.Header.Add("Authorization", "Bearer "+authenticateToken)
	}

	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if nil != err {
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	fmt.Printf("%s\n%s\n", Url, body)
	return string(body)
}

// Signature Utils
func Uint64ToBigEndianBytes(value uint64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, value)
	return bs
}

func Uint32ToBigEndianBytes(value uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, value)
	return bs
}

// The returned bytes are to be hashed (using keccak256) for signing. The content are the
// concatenation of the following (uints are in big-endian byte order):
//
// 1. String "DEx2 Order: "(96)  (Note the trailing whitespace)
// 2. <market address>(160)      (For preventing cross-market replay attack)
// 3. <nonce>(64) <expireTimeSec>(64) <amountE8>(64) <priceE8>(64) <ioc>(8) <action>(8) <pairId>(32)
func GetOrderBytesToSign(marketAddr common.Address, nonce uint64, order *models.Order) []byte {
	bs := make([]byte, 0, 98)
	bs = append(bs, []byte("\x19Ethereum Signed Message:\n70")...)
	bs = append(bs, []byte("DEx2 Order: ")...)
	bs = append(bs, marketAddr.Bytes()...)
	bs = append(bs, Uint64ToBigEndianBytes(nonce)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.ExpireTimeSec)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.AmountE8)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.PriceE8)...)
	bs = append(bs, order.Ioc)
	bs = append(bs, order.Action)
	bs = append(bs, Uint32ToBigEndianBytes(order.PairId)...)

	if len(bs) != 98 { // 784 bits
		fmt.Printf("The byte length of signing an order must be 98, but got ", len(bs))
	}
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

//-------------------------------------
// Public Info API
//
// Get PairInfo By CashId like ETH
// return PairInfo Array
func GetPairsByCash(CashId string) models.GetPairsByCashResponse {

	GetPairsByCashResponse := models.GetPairsByCashResponse{}

	requestUrl := dextopHost + "v1/pairlist/" + CashId

	body := request("GET", requestUrl, nil, false)

	json.Unmarshal([]byte(body), &GetPairsByCashResponse)

	fmt.Printf("%+v", GetPairsByCashResponse)
	return GetPairsByCashResponse
}

// Get Depth
//
func GetPairDepth(PairId string, Count int) {

	GetPairDepthResponse := models.GetPairDepthResponse{}
	requestUrl := dextopHost + "v1/depth/" + PairId + "/" + strconv.Itoa(Count)

	body := request("GET", requestUrl, nil, false)
	json.Unmarshal([]byte(body), &GetPairDepthResponse)
}

// Get MarketInfo, Important!!! Including cashcode and tokencode for placeorder
//
func GetMarket() models.MarketInfo {

	MarketInfo := models.MarketInfo{}
	requestUrl := dextopHost + "v1/market"
	body := request("GET", requestUrl, nil, false)
	json.Unmarshal([]byte(body), &MarketInfo)

	marketAddr = MarketInfo.MarketAddr

	for _, v := range MarketInfo.Config.CashTokens {
		TokenCodeBySymbol[v.TokenId] = v.TokenCode
	}
	for _, v := range MarketInfo.Config.StockTokens {
		TokenCodeBySymbol[v.TokenId] = v.TokenCode
	}
	fmt.Printf("%+v", MarketInfo)
	return MarketInfo
}

//---------------------

// Account API
// Login Token Get Header token
func Login(Email string, Password string) models.LoginResponse {

	LoginResponse := models.LoginResponse{}

	requestUrl := dextopHost + "v1/authenticate"

	mapParams := make(map[string]string)
	mapParams["email"] = Email
	mapParams["password"] = Password

	body := request("POST", requestUrl, mapParams, false)

	json.Unmarshal([]byte(body), &LoginResponse)

	authenticateToken = LoginResponse.Token

	return LoginResponse
}

// Get Account Balance of traderAddr which has been binded
func GetBalance(TraderAddr string) {
	requestUrl := dextopHost + "v1/balances/" + TraderAddr
	request("GET", requestUrl, nil, true)
}

//---------------------
// Trade API

func PlaceOrder(traderAddr string, pairId string, action string, price string, amount string) {

	requestUrl := dextopHost + "v1/placeorder"

	// order action
	pbAction := 0
	switch action {
	case "Buy":
		pbAction = 1
	case "Sell":
		pbAction = 2
	}

	expireTimeSec := time.Now().Unix() + 3600
	nonce := time.Now().UnixNano() / 1e6

	// singuare order body
	order := &models.Order{
		PairId:        GetPairCode(pairId),
		Action:        byte(pbAction) - 1,
		Ioc:           0,
		PriceE8:       StringToUint64E8(price),
		AmountE8:      StringToUint64E8(amount),
		ExpireTimeSec: uint64(expireTimeSec),
	}

	// market address
	addr := common.HexToAddress(marketAddr)
	// trans signature info into bytes
	bytesToSign := GetOrderBytesToSign(addr, uint64(nonce), order)
	hash := crypto.Keccak256(bytesToSign)
	// trader addr private key
	traderPrivKey, err := crypto.HexToECDSA(testTraderAddrPriKey)
	if err != nil {
		return
	}
	// Signed
	sig, err := crypto.Sign(hash, traderPrivKey)
	if err != nil {
		return
	}

	// Place Order request Params
	mapParams := make(map[string]string)
	mapParams["pairId"] = pairId
	mapParams["traderAddr"] = traderAddr
	mapParams["action"] = action
	mapParams["price"] = price
	mapParams["amount"] = amount
	mapParams["expireTimeSec"] = strconv.FormatInt(expireTimeSec, 10)
	mapParams["nonce"] = strconv.FormatInt(nonce, 10)
	mapParams["sig"] = common.ToHex(sig)

	if err != nil {
		return
	}

	request("POST", requestUrl, mapParams, true)

}

func GetActiveOrders(TraderAddr string, PairId string, Size int, Page int) {
	requestUrl := dextopHost + "v1/activeorders/" + TraderAddr + "/" + PairId + "/" + strconv.Itoa(Size) + "/" + strconv.Itoa(Page)
	request("GET", requestUrl, nil, true)

}

func GetPastOrders(TraderAddr string, PairId string, Size int, Page int) {
	requestUrl := dextopHost + "v1/pastorders/" + TraderAddr + "/" + PairId + "/" + strconv.Itoa(Size) + "/" + strconv.Itoa(Page)
	request("GET", requestUrl, nil, true)

}

func main() {

	// Public API
	GetPairsByCash("ETH")
	GetPairDepth("ETH_BTM", 10)

	GetMarket()
	// Account API need token
	Login(testAccount, testAccountPwd)
	GetBalance(testTraderAddr)

	// PlaceOrder need token and sig
	PlaceOrder(testTraderAddr, "ETH_BTM", "Buy", "0.00001", "10000")

	// Account Trade Info API need token
	GetActiveOrders(testTraderAddr, "ETH_BTM", 10, 1)
	GetPastOrders(testTraderAddr, "ETH_BTM", 10, 1)

}
