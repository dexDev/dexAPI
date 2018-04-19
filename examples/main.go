package main

func main() {
	// Public apis do not require token and signature
	GetMarket()
	GetPairsByCash("ETH")
	GetPairDepth("ETH_BTM", 10)

	Login(userName, userPwd)

	// Place order requires signature the order information
	PlaceOrder(userBindingTraderAddr, "ETH_BTM", "Buy", "0.00001", "10000")

	// Account related information needs token
	GetBalance(userBindingTraderAddr)
	GetActiveOrders(userBindingTraderAddr, "ETH_BTM", 10, 1)
	GetPastOrders(userBindingTraderAddr, "ETH_BTM", 10, 1)
}
