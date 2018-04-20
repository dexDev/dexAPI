// Api response models
//
package models

// TODO: more models and more detail descriptions

type PlaceOrderRequest struct {
	Amount        string `json:"amount"`
	Price         string `json:"price"`
	Action        string `json:"action"`
	PairId        string `json:"pairId"`
	Nonce         int    `json:"amount"`
	TraderAddr    string `json:"traderAddr"`
	Sig           string `json:"sig"`
	ExpireTimeSec int    `json:"expireTimeSec"`
}

type GetPairDepthResponse struct {
	DepthInfo DepthInfo `json: "depth"`
}

type DepthInfo struct {
	PairId string  `json: "pairId"`
	TimeMs string  `json: "timeMs"`
	Asks   []Depth `json: "asks"`
	Bids   []Depth `json: "bids"`
}

type Depth struct {
	Price  string `json: "price"`
	Amount string `json: "Amount"`
}

type PairInfo struct {
	PairId          string `json:"pairId"`
	LastPrice       string `json:"lastPrice"`
	Volume24        string `json:"volume24"`
	Change24        string `json:"change24"`
	ChangePercent24 string `json:"changePercent24"`
	TimeMs          string `json:"timeMs"`
	High24          string `json:"high24"`
	Low24           string `json:"low24"`
}

type GetPairsByCashResponse struct {
	Pairs []PairInfo `json:"pairs"`
}

type User struct {
	Id            string `json: "id"`
	Email         string `json: "email"`
	EmailVerified bool   `json: "emailVerified"`
}

type LoginResponse struct {
	Token string `json: "token"`
	User  User   `json: "user"`
}

type Order struct {
	PairId        uint32
	Action        uint8
	Ioc           uint8
	PriceE8       uint64
	AmountE8      uint64
	ExpireTimeSec uint64
}

type TokenInfo struct {
	TokenId     string `json: "tokenId"`
	TokenCode   uint16 `json: "tokenCode"`
	TokenAddr   string `json: "tokenAddr"`
	scaleFactor string `json: "scaleFactor"`
}

type MarketConfig struct {
	CashTokens  []TokenInfo `json: "cashTokens"`
	StockTokens []TokenInfo `json: "stockTokens"`
}

type Market struct {
	MarketAddr string       `json: "marketAddr"`
	Config     MarketConfig `json: "config"`
}
