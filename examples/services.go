package main

import (
	"encoding/json"
	"strconv"
	"time"

	"fmt"
	"github.com/dexDev/dexAPI/examples/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// NOTE: Preparation
//- 1. Register an account using email on https://kovan.dex.top
//- 2. Binding trader address in the account page
//- 3. Deposit eth or stock tokens in the balance page

// DEx.top testnet api host
var dextopTestnetHost = "https://kovan.dex.top"

// Test account
const (
	userName                 = "" // test@dex.top
	userPwd                  = "" // 12345678a
	userBindingTraderAddr    = "" // 0x6a83D834951F29924559B8146D11a70EaB8E328b
	userBingdingTraderPriKey = "" // string privatekey
)

// Cache user auth token and market related information
var (
	authenticateToken string
	marketAddr        string
	tokenCodesById    = make(map[string]uint16)
)

// Market information contains the token's code, these are the necessary information when placing orders,
// where we cache them in this `tokenCodesById` map.
func GetMarket() (*models.Market, error) {
	market := models.Market{}
	resp, err := httpRequest("GET", dextopTestnetHost+"/v1/market", nil, false)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp, &market); err != nil {
		return nil, err
	}

	marketAddr = market.MarketAddr
	for _, v := range market.Config.CashTokens {
		tokenCodesById[v.TokenId] = v.TokenCode
	}
	for _, v := range market.Config.StockTokens {
		tokenCodesById[v.TokenId] = v.TokenCode
	}
	return &market, nil
}

// TODO: do we need to define a struct for the corresponding normal json response

// Get all pair information for the specified cash token.
func GetPairsByCash(cashTokenId string) (*models.GetPairsByCashResponse, error) {
	getPairsByCashResponse := models.GetPairsByCashResponse{}
	resp, err := httpRequest("GET", dextopTestnetHost+"/v1/pairlist/"+cashTokenId, nil, false)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp, &getPairsByCashResponse); err != nil {
		return nil, err
	}

	return &getPairsByCashResponse, nil
}

//Get the depth data of a certain transaction pair
//TODO: define depth response
func GetPairDepth(pairId string, size int) (*models.GetPairDepthResponse, error) {
	getPairDepthResponse := models.GetPairDepthResponse{}
	url := fmt.Sprintf("%s/%s/%d", dextopTestnetHost+"/v1/depth", pairId, size)
	resp, err := httpRequest("GET", url, nil, false)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp, &getPairDepthResponse); err != nil {
		return nil, err
	}

	return &getPairDepthResponse, nil
}

// Login and get auth token
func Login(email string, password string) (*models.LoginResponse, error) {
	loginResponse := models.LoginResponse{}

	mapParams := make(map[string]string)
	mapParams["email"] = email
	mapParams["password"] = password

	resp, err := httpRequest("POST", dextopTestnetHost+"/v1/authenticate", mapParams, false)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp, &loginResponse); err != nil {
		return nil, err
	}

	authenticateToken = loginResponse.Token
	return &loginResponse, nil
}

// Get account balances of a trader
// TODO: more implementation
func GetBalance(traderAddr string) {
	url := dextopTestnetHost + "v1/balances/" + traderAddr
	httpRequest("GET", url, nil, true)
}

// Place an order
// TODO: more implementation, return the placed order
func PlaceOrder(traderAddr string, pairId string, action string, price string, amount string) (resp string, err error) {
	pbAction := 0
	switch action {
	case "Buy":
		pbAction = 1
	case "Sell":
		pbAction = 2
	}
	expireTimeSec := time.Now().Unix() + 3600
	nonce := time.Now().UnixNano() / 1e6

	// Signature the order body content
	order := &models.Order{
		PairId:        getPairCode(pairId),
		Action:        byte(pbAction) - 1,
		Ioc:           0,
		PriceE8:       stringToUint64E8(price),
		AmountE8:      stringToUint64E8(amount),
		ExpireTimeSec: uint64(expireTimeSec),
	}
	addr := common.HexToAddress(marketAddr) // market addr
	bytesToSign := getOrderBytesToSign(addr, uint64(nonce), order)
	hash := crypto.Keccak256(bytesToSign)
	traderPrivKey, err := crypto.HexToECDSA(userBingdingTraderPriKey) // user private key
	if err != nil {
		return resp, err
	}
	sig, err := crypto.Sign(hash, traderPrivKey) // signature the hash
	if err != nil {
		return resp, err
	}

	// Place order request params
	mapParams := make(map[string]string)
	mapParams["pairId"] = pairId
	mapParams["traderAddr"] = traderAddr
	mapParams["action"] = action
	mapParams["price"] = price
	mapParams["amount"] = amount
	mapParams["expireTimeSec"] = strconv.FormatInt(expireTimeSec, 10)
	mapParams["nonce"] = strconv.FormatInt(nonce, 10)
	mapParams["sig"] = common.ToHex(sig)

	respBytes, err := httpRequest("POST", dextopTestnetHost+"/v1/placeorder", mapParams, true)
	return string(respBytes), err
}

// Get active orders of a specified trader address
// TODO: more implementation
func GetActiveOrders(traderAddr string, pairId string, size int, page int) error {
	url := fmt.Sprintf("%s/%s/%s/%d/%d", /* apiPath/traderAddr/pairId/size/page */
		dextopTestnetHost+"/v1/activeorders", traderAddr, pairId, size, page)
	_, err := httpRequest("GET", url, nil, true)
	return err
}

// Get past orders of a specified trader address
// TODO: more implementation
func GetPastOrders(traderAddr string, pairId string, size int, page int) error {
	url := fmt.Sprintf("%s/%s/%s/%d/%d", /* apiPath/traderAddr/pairId/size/page */
		dextopTestnetHost+"/v1/pastorders", traderAddr, pairId, size, page)
	_, err := httpRequest("GET", url, nil, true)
	return err
}
