package main

import "fmt"

func main() {
	// Public apis do not require token and signature
	market, err := GetMarket()
	if err != nil {
		fmt.Println("Get market failed. ", err)
		return
	}
	fmt.Printf("Market: \n%+v\n", *market)

	pairs, err := GetPairsByCash("ETH")
	if err != nil {
		fmt.Println("Get pairs failed. ", err)
		return
	}
	fmt.Printf("Pairs information: \n%+v\n", *pairs)

	//GetPairDepth("ETH_BTM", 10)

	if _, err := Login(userName, userPwd); err != nil {
		fmt.Println("Login failed. ", err)
		return
	}

	// Place order requires signature the order information
	err = PlaceOrder(userBindingTraderAddr, "ETH_BTM", "Buy", "0.00001", "10000")
	if err != nil {
		fmt.Println("Place order failed. ", err)
		return
	}
	fmt.Println("Order placed")

	// Account related information needs token
	GetBalance(userBindingTraderAddr)
	GetActiveOrders(userBindingTraderAddr, "ETH_BTM", 10, 1)
	GetPastOrders(userBindingTraderAddr, "ETH_BTM", 10, 1)
}
