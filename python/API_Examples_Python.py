#!/usr/bin/env python3

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

HOST_URL = "https://dex.top"
"""
For testnet use "https://testnet271828.dex.top"
"""
USER_INFORMATION = {"token": "",
                    "traderAddr": "",
                    "email": "",
                    "password": ""}

def get_market_info():
    """
    Get relevant market information such as market contract address and token codes for signature.
    """
    resp = requests.get(HOST_URL + "/v1/market")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("Failed to get market data: ", resp.json())

def get_pairs_by_cash():
    """
    Get the real-time trading information of all available trading pairs of the specified cash token (e.g. "ETH")
    """
    resp = requests.get(HOST_URL + "/v1/pairlist/ETH")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("Failed to get pairs by cash: ", resp.json())

def get_pair_info():
    """
    Get the real-time information of a trading pair.
    """
    resp = requests.get(HOST_URL + "/v1/pairinfo/ETH_CHSB")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("Failed to get pair info: ", resp.json())

def get_trade_history():
    """
    Get recent trades by pairId sort by trade time.
    """
    resp = requests.get(HOST_URL + "/v1/tradehistory/ETH_CHSB/9")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("Failed to get trade history: ", resp.json())

def get_pair_Depth():
    """""
    Get the depth data of a trading pair.
    """
    resp = requests.get(HOST_URL + "/v1/depth/ETH_YEE/10")
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("Failed to get pair depth: ", resp.json())

def user_login():
    """
    The login function that takes 'email' and 'password' from 'USER_INFORMATION'
    The call returns a token that is placed in USER_INFORMATION
    """
    pay_load = {"email": USER_INFORMATION['email'], "password": USER_INFORMATION['password']}
    pay_load = json.dumps(pay_load)
    r = requests.post(HOST_URL + "/v1/authenticate", pay_load)
    if r.status_code == requests.codes.ok:
        resp = r.json()
        USER_INFORMATION["token"] = resp['token']
        print("Login Successful")
    else:
        print("Failed user login ", r.json())

def get_balance():
    """
    Get the balances of all tokens of a trader.
    """
    headers = {"Authorization": USER_INFORMATION["token"]}
    resp = requests.get(HOST_URL + "/v1/balances/" + USER_INFORMATION["traderAddr"], headers=headers)
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("Failed to get balance: ", resp.json())

def place_order():
    """
    Place a new order
    """
    headers = {"Authorization": USER_INFORMATION["token"]}
    signature = example_sign_order.play()
    #print("Signature: ", signature)
    post_data = {'pairId': 'ETH_ZIL',
                 'action': 'Buy',
                 'amount': '660.00',
                 'price': '0.0001600',
                 'nonce': 3640819234038,
                 'expireTimeSec': 1531251449,
                 'traderAddr': USER_INFORMATION["traderAddr"],
                 'sig': signature}
    post_data = json.dumps(post_data)
    order_response = requests.post(HOST_URL + "/v1/placeorder", post_data, headers=headers)
    if order_response.status_code == requests.codes.ok:
        pprint.pprint(order_response.json())
    else:
        print(order_response.content)
        print("Failed to place order: ", order_response.json())

def cancel_order():
    """
    cancel a single specific order
    """
    headers = {"Authorization": USER_INFORMATION["token"]}
    pay_load = {"traderAddr": USER_INFORMATION["traderAddr"],
                "orderId": "10827893",
                "pairId": "ETH_ZIL",
                "nonce": "2640819234038"}
    pay_load = json.dumps(pay_load)
    cancel_order_response = requests.post(HOST_URL + "/v1/cancelorder", pay_load, headers=headers)
    if cancel_order_response.status_code == requests.codes.ok:
        print("ok")
    else:
        print("Failed to cancel order: ", cancel_order_response.json())

def cancel_all_orders():
    """
    Cancel all existing orders using given parameters
    """
    headers = {"Authorization": USER_INFORMATION["token"]}
    pay_load = {"traderAddr": USER_INFORMATION["traderAddr"],
                "pairId": 'ETH_ZIL',
                "nonce": '3640819234038'}
    pay_load = json.dumps(pay_load)
    resp = requests.post(HOST_URL + "/v1/cancelallorders", pay_load, headers=headers)
    if resp.status_code == requests.codes.ok:
        print(resp.content)
    else:
        print("Failed to cancel all orders: ", resp.json())

def get_active_orders():
    """
    Get unfilled or partially filled orders that have not been cancelled or expired of a trader.
    """
    headers = {"Authorization": USER_INFORMATION["token"]}
    url_params = "/v1/activeorders/" + USER_INFORMATION["traderAddr"] + "/ETH_ZIL/100/1"
    resp = requests.get(HOST_URL + url_params, headers=headers)
    if resp.status_code == requests.codes.ok:
        pprint.pprint(resp.json())
    else:
        print("Failed to get active orders: ", resp.json())

def get_trades():
    """
    Get recent trades on current wallet address.
    """
    headers = {"Authorization": USER_INFORMATION["token"]}
    url_params = "/v1/trades/" + USER_INFORMATION["traderAddr"] + "/ETH_CHSB/100"
    response = requests.get(HOST_URL + url_params, headers=headers)
    if response.status_code == requests.codes.ok:
        pprint.pprint(response.json())
    else:
        print("Failed to get trades: ", response.json())


if __name__ == '__main__':
    user_login()        #This call returns a token, which should be placed in USER_INFORMATION["token"]



