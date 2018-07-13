#!/usr/bin/env python3

from decimal import Decimal

import json                  # python3 -m pip install simplejson
import requests              # python3 -m pip install requests
import pprint                # python3 -m pip install pprintpp
import arrow                 # python3 -m pip install arrow
import eth_utils             # python3 -m pip install eth_utils==1.0.3
from web3.auto import w3     # python3 -m pip install web3==4.2.0


# The following is a Python Script showing the use of current API's
# The USER_INFORMATION dictionary needs to filled with the appropriate information


USER_INFORMATION = {"token": "",        # This is returned by the user_login() function
                    "traderAddr": "",
                    "email": "",
                    "password": "",
                    "privateKey": ""}


# URL's for testnet and mainnet
HOST_URL = "https://dex.top"
TESTNET_URL = "https://testnet271828.dex.top"

# Market address for mainnet and testnet
MAINNET_MARKET_ADDR = "0x7600977Eb9eFFA627D6BD0DA2E5be35E11566341"
TESTNET_MARKET_ADDR = "0x752B76F59cdaE6925377569E3f9352A507535F99"

"""
For the full list of Token codes go to the following URL and populate dictionary accordingly
URL = "HOST_URL/v1/market"
"""

TOKEN_CODE_DICT = {}


def sign_order(order, private_key):
    """
    Takes the order and private key and returns the signature required for trades
    """

    market_addr_canon = eth_utils.to_canonical_address(MAINNET_MARKET_ADDR)

    symbol = order['pairId']
    pair_code = get_pair_code(symbol)
    pair_bytes = fixed_len_byte(pair_code, fixed_len=4)

    if order['action'] == "Buy":
        action_byte = 0
    else:
        action_byte = 1

    price = order['price']
    decimal_price = int(Decimal(price) * Decimal(1e8))
    price_bytes = fixed_len_byte(decimal_price)
    amount = order['amount']
    decimal_amount = int(Decimal(amount) * Decimal(1e8))
    amount_bytes = fixed_len_byte(decimal_amount)

    expire_time_sec = order['expireTimeSec']
    expire_time_sec_bytes = fixed_len_byte(expire_time_sec)

    nonce = order['nonce']
    nonce_bytes = fixed_len_byte(nonce)

    ioc = 0

    order_msg_bytes = get_orders_bytes_sign(
                    market_addr_canon, nonce_bytes, expire_time_sec_bytes, amount_bytes,
                    price_bytes, ioc, action_byte, pair_bytes)

    go_style_output = ''
    for b in order_msg_bytes:
        go_style_output += '{0:02x}'.format(b)

    hash_message = eth_utils.keccak(order_msg_bytes)
    signed_message = w3.eth.account.signHash(hash_message, private_key=private_key)

    return remove_sign_v_offset(signed_message['signature'])


def get_pair_code(pair_id):
    """
    Get pair code from pair id
    """
    sub_symbol = pair_id.split("_")
    cash_code = TOKEN_CODE_DICT[sub_symbol[0]]
    stock_code = TOKEN_CODE_DICT[sub_symbol[1]]
    return eth_utils.to_int(cash_code) << 16 | eth_utils.to_int(stock_code)


def fixed_len_byte(value, fixed_len=8):
    """
    Converts the value to big endian format and returns the appropriate bytes
    """
    int_bytes = eth_utils.int_to_big_endian(value)
    result_bytes = bytearray(0 for _ in range(fixed_len - len(int_bytes)))
    result_bytes.extend(int_bytes)
    return result_bytes


def remove_sign_v_offset(sign_msg):
    sign_bytes = bytearray(sign_msg)
    sign_bytes[len(sign_bytes) - 1] -= 27
    return eth_utils.to_hex(sign_bytes)


def get_orders_bytes_sign(market_addr, nonce, expire_time, amount, price, ioc, action, pairid):
    """
    Type of all params is bytes!
    """
    sign_bytes = bytearray()
    sign_bytes.extend(eth_utils.to_bytes(text="\x19Ethereum Signed Message:\n70"))
    sign_bytes.extend(eth_utils.to_bytes(text="DEx2 Order: "))
    sign_bytes.extend(market_addr)
    sign_bytes.extend(nonce)
    sign_bytes.extend(expire_time)
    sign_bytes.extend(amount)
    sign_bytes.extend(price)
    sign_bytes.append(ioc)
    sign_bytes.append(action)
    sign_bytes.extend(pairid)
    return sign_bytes


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
    The call returns a token that is placed in USER_INFORMATION["token"]
    """
    pay_load = {"email": USER_INFORMATION['email'], "password": USER_INFORMATION['password']}
    pay_load = json.dumps(pay_load)
    r = requests.post(HOST_URL + "/v1/authenticate", pay_load)
    if r.status_code == requests.codes.ok:
        resp = r.json()
        print("Login Successful")
    else:
        print("Failed user login ", r.json())
    return resp["token"]

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

def place_order(order_info):
    """
    Place a new order, with order params passed in
    """
    headers = {"Authorization": USER_INFORMATION["token"]}
    post_data = order_info
    signature = get_signature(post_data)
    post_data["sig"] = signature
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
                "nonce": 10000000000004}
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


def get_signature(order_data):
    """
    The get_signature() function returns the signature for order signing
    It takes in the order details for the input parameter
    """
    signature_hex = sign_order(order_data, USER_INFORMATION["privateKey"])
    return signature_hex


def populate_token_dictionary():
    """
    This function fills in the TOKEN_CODE_DICT with the 'tokenCode' and 'tokenId' for all current tokens on market
    """
    token_code_dict_temp = {}
    resp = requests.get(HOST_URL + "/v1/market")
    var = resp.json()
    for cash_token in var['config']['cashTokens']:
        token_code_dict_temp[cash_token['tokenId']] = cash_token['tokenCode']
    for token in var['config']['stockTokens']:
        token_code_dict_temp[token['tokenId']] = token['tokenCode']
    return token_code_dict_temp


if __name__ == '__main__':

    TOKEN_CODE_DICT = populate_token_dictionary()
    USER_INFORMATION["token"] = user_login()    # user_login returns a token which is stored in USER_INFORMATION[token]
    order_payload = {"pairId": "ETH_ZIL",
                     "action": "Buy",
                     "amount": "1000.00",
                     "price": "0.0001500",
                     "nonce": 10000000000003,
                     "expireTimeSec": 1531359841,
                     "traderAddr": USER_INFORMATION["traderAddr"]
                     }
    #place_order(order_payload)