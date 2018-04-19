/** Account prepare

  1. register a account by email on https://kovan.dex.top
  2. binding trader address in account page.
  3. deposit eth or tokens in balance page.

**/
// orders example

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/dexDev/dexAPI/examples/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)


// DEx.top testnet api host
var dextopTestnetHost = "https://kovan.dex.top/"

// Test account
const (
 userName = "flynn@dex.top"
 userPwd = "12345678a"
 userBindingTraderAddr = "0x6a83D834951F29924559B8146D11a70EaB8E328b"
 userBingdingTraderPriKey = "121e1348709ca0f75ea8793bbac27886afe6eb272c9a5245890aa7e4c64a65b9"
)

// Cache user auth token and market related information
var (
	authenticateToken string
	marketAddr string
	TokenCodeBySymbol = make(map[string]uint16)
)

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

// Market information contains the token's code, these are the necessary information when placing orders,
// where we cache them in this `TokenCodeBySymbol` map.
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
	bytesToSign := getOrderBytesToSign(addr, uint64(nonce), order)
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
	// Public apis do not require token and signature
	GetMarket()
	GetPairsByCash("ETH")
	GetPairDepth("ETH_BTM", 10)

	Login(testAccount, testAccountPwd)

	// Place order requires signature the order information
	PlaceOrder(testTraderAddr, "ETH_BTM", "Buy", "0.00001", "10000")

	// Account related information needs token
	GetBalance(testTraderAddr)
	GetActiveOrders(testTraderAddr, "ETH_BTM", 10, 1)
	GetPastOrders(testTraderAddr, "ETH_BTM", 10, 1)
}
