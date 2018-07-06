import json     # python3 -m pip install simplejson
import requests # python3 -m pip install requests
import pprint   # python3 -m pip install pprintpp
import example_sign_order


"""
The following is a Python Script showing the use of current API's
The following information must be inputted before using
user_information 
    - traderAddr 
    - email
    - password
example_sign_order.py
- post_data
    - traderAddr
- Private Key

"""
testnet_url = "https://testnet271828.dex.top"
mainnet_url = "https://dex.top"
user_information = {"token": "",
                    "traderAddr": "",
                    "email": "",
                    "password": ""}

def get_market_info():
    """
    Get relevant market information such as market contract address and token codes for signature.
    """
    resp = requests.get(mainnet_url + "/v1/market")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("ERROR: Get Market Information request did not work. ")

def get_pairs_by_cash():
    """
    Get the real-time trading information of all available trading pairs of the specified cash token (e.g. "ETH")
    """
    resp = requests.get(mainnet_url + "/v1/pairlist/ETH")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("ERROR: Get Pairs by Cash request did not work. ")

def get_pair_info():
    """
    Get the real-time information of a trading pair.
    """
    resp = requests.get(mainnet_url + "/v1/pairinfo/ETH_CHSB")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("ERROR: Get Pair Information did not work. ")

def get_trade_history():
    """
    Get recent trades by pairId sort by trade time.
    """
    resp = requests.get(mainnet_url + "/v1/tradehistory/ETH_CHSB/9")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("ERROR: Get Trade History request did not work. ")

def get_pair_Depth():
    """""
    # Get the depth data of a trading pair.
    """
    resp = requests.get(mainnet_url + "/v1/depth/ETH_YEE/10")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("ERROR: Get Pair Depth request did not work. ")

def user_login():
    """
    API used to login to exchange - Returns a token
    """
    pay_load = {"email": user_information['email'], "password": user_information['password']}
    pay_load = json.dumps(pay_load)
    r = requests.post(mainnet_url + "/v1/authenticate", pay_load)
    if r.status_code == requests.codes.ok:
        resp = r.json()
        user_information["token"] = resp['token']
        print("Login Successful")
    else:
        print("ERROR: User login post did not work. ")

def get_balance():
    """
    Get the balances of all tokens of a trader.
    """
    headers = {"Authorization": user_information["token"]}
    resp = requests.get(mainnet_url + "/v1/balances/" + user_information["traderAddr"], headers=headers)
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("ERROR: Get Balance request did not work. ")

def place_order():
    """
    Place a new order
    """
    headers = {"Authorization": user_information["token"]}
    signature = example_sign_order.play()
    #print("Signature: ", signature)
    post_data = {'pairId': 'ETH_ZIL',
                 'action': 'Buy',
                 'amount': '660.00',
                 'price': '0.0001600',
                 'nonce': 3640819234038,
                 'expireTimeSec': 1531251449,
                 'traderAddr': user_information["traderAddr"],
                 'sig': signature}
    post_data = json.dumps(post_data)
    order_response = requests.post(mainnet_url + "/v1/placeorder", post_data, headers=headers)
    if order_response.status_code == requests.codes.ok:
        pprint.pprint(order_response.json())
    else:
        print(order_response.content)
        print("ERROR: Place Order post did not work. ")

def cancel_order():
    """
    cancel a single specific order
    """
    headers = {"Authorization": user_information["token"]}
    pay_load = {"traderAddr": user_information["traderAddr"],
                "orderId": "10827893",
                "pairId": "ETH_ZIL",
                "nonce": "2640819234038"}
    pay_load = json.dumps(pay_load)
    cancel_order_response = requests.post(mainnet_url + "/v1/cancelorder", pay_load, headers=headers)
    if cancel_order_response.status_code == requests.codes.ok:
        print("ok")
    else:
        print("ERROR: Cancel Order post did not work. ")

def cancel_all_orders():
    """
    # Cancel all existing orders using given parameters
    """
    headers = {"Authorization": user_information["token"]}
    pay_load = {"traderAddr": user_information["traderAddr"],
                "pairId": 'ETH_ZIL',
                "nonce": '3640819234038'}
    pay_load = json.dumps(pay_load)
    response = requests.post(mainnet_url + "/v1/cancelallorders", pay_load, headers=headers)
    print(response.content)

def get_active_orders():
    """
    Get unfilled or partially filled orders that have not been cancelled or expired of a trader.
    """
    headers = {"Authorization": user_information["token"]}
    pay_load = {"pairId": 'ETH_ZIL',
                "size": '100',
                "page": '1',
                "traderAddr": user_information["traderAddr"]}
    url_params = "/v1/activeorders/" + pay_load["traderAddr"] + "/" + pay_load['pairId'] + "/" + pay_load['size'] + "/" + pay_load["page"]
    pay_load = json.dumps(pay_load)
    response = requests.get(mainnet_url + url_params, pay_load, headers=headers)
    if response.status_code == requests.codes.ok:
        pprint.pprint(response.json())
    else:
        print("ERROR: Get Active Orders request did not work. ")

def get_trades():
    """
    Get recent trades on current wallet address.
    """
    headers = {"Authorization": user_information["token"]}
    pay_load = {'pairId': 'ETH_ZIL',
                'size': '100',
                'traderAddr': user_information["traderAddr"]}
    url_params = "/v1/trades/" + pay_load["traderAddr"] + "/" + pay_load["pairId"] + "/" + pay_load["size"]
    pay_load = json.dumps(pay_load)
    response = requests.get(mainnet_url + url_params, pay_load, headers=headers)
    if response.status_code == requests.codes.ok:
        pprint.pprint(response.json())
    else:
        print("ERROR: Get trades request did not work. ")


if __name__ == '__main__':
    user_login()
    #get_active_orders()
    #get_trades()
    #cancel_all_orders()
    #place_order()
    #user_login()
    #get_balance()
    #get_market_info()
    #get_pairs_by_cash()
    #get_pair_info()
    #get_trade_history()
    #get_pair_Depth()

