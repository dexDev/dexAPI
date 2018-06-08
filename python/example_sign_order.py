#!/usr/bin/env python3

from decimal import Decimal

import arrow              # python3 -m pip install arrow
import eth_utils          # python3 -m pip install eth_utils==1.0.3
from web3.auto import w3  # python3 -m pip install web3==4.2.0

TOKEN_CODE_DICT = {
    'ETH': 0,
    'LOOM': 100,
    'KNC': 101,
    'ZIL': 102,
    'CTXC': 103
}

MARKET_ADDRESS_DICT = {
    'ETH_LOOM': '0x7600977Eb9eFFA627D6BD0DA2E5be35E11566341',
    'ETH_KNC': '0x7600977Eb9eFFA627D6BD0DA2E5be35E11566341',
    'ETH_ZIL': '0x7600977Eb9eFFA627D6BD0DA2E5be35E11566341',
    'ETH_CTXC': '0x7600977Eb9eFFA627D6BD0DA2E5be35E11566341'
}

def get_nonce():
    return int(arrow.now().float_timestamp * 1000)


def get_expire_time_sec(shift=3600):
    return arrow.now().timestamp + shift


def hex_to_address(addr):
    a = eth_utils.to_canonical_address(addr)
    return a


def get_pair_code(symbol):
    sub_symbol = symbol.split("_")
    cash_code = TOKEN_CODE_DICT[sub_symbol[0]]
    stock_code = TOKEN_CODE_DICT[sub_symbol[1]]
    return eth_utils.to_int(cash_code) << 16 | eth_utils.to_int(stock_code)


def fixed_len_byte(value, fixed_len=8):
    int_bytes = eth_utils.int_to_big_endian(value)
    # print(value, list(int_bytes))
    result_bytes = bytearray(0 for _ in range(fixed_len - len(int_bytes)))
    result_bytes.extend(int_bytes)
    return result_bytes


def get_orders_bytes_sign(market_addr, nonce, expire_time, amount, price, ioc, action, pairid):
    """
    Type of all params is bytes !
    :param market_addr:
    :param nonce:
    :param expire_time:
    :param amount:
    :param price:
    :param ioc:
    :param action:
    :param pairid:
    :return:
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


def remove_sign_v_offset(sign_msg):
    sign_bytes = bytearray(sign_msg)
    sign_bytes[len(sign_bytes) - 1] -= 27
    return eth_utils.to_hex(sign_bytes)


def sign_order_msg(order, private_key):
    symbol = order['pairId']
    market_addr = MARKET_ADDRESS_DICT[symbol]
    market_addr_bytes = hex_to_address(market_addr)
    # print('market_addr', list(market_addr_bytes))

    pair_code = get_pair_code(symbol)

    pair_bytes = fixed_len_byte(pair_code, fixed_len=4)
    # print('pair', pair_code, list(pair_bytes))

    bs = order['action']
    bs_code = 1 if bs == 'Buy' else 2
    action_byte = bs_code - 1
    # print('action', action_byte)

    price = order['price']
    decimal_price = int(Decimal(price) * Decimal(1e8))
    price_bytes = fixed_len_byte(decimal_price)
    # print('price', list(price_bytes))
    amount = order['amount']
    decimal_amount = int(Decimal(amount) * Decimal(1e8))
    amount_bytes = fixed_len_byte(decimal_amount)
    # print('amount', list(amount_bytes))

    expire_time_sec = order['expireTimeSec']
    expire_time_sec_bytes = fixed_len_byte(expire_time_sec)
    # print('amount', list(expire_time_sec_bytes))

    nonce = order['nonce']
    nonce_bytes = fixed_len_byte(nonce)
    # print('nonce', list(nonce_bytes))

    ioc = 0

    order_msg_bytes = get_orders_bytes_sign(market_addr_bytes, nonce_bytes, expire_time_sec_bytes, amount_bytes,
                                            price_bytes, ioc, action_byte, pair_bytes)
    # print('order_msg_bytes', list(order_msg_bytes))

    go_style_output = ''
    for b in order_msg_bytes:
        go_style_output += '{0:02x}'.format(b)
    print('order_msg_bytes:', '0x'+go_style_output, '\n')

    hash_message = eth_utils.keccak(order_msg_bytes)
    signed_message = w3.eth.account.signHash(hash_message, private_key=private_key)
    # logging.info(symbol, market_addr[:6], pair_code, bs_code, price, amount, expire_time_sec, nonce, ioc,
    #              private_key[:6])
    return remove_sign_v_offset(signed_message['signature'])


def play():
    post_data = {'pairId': 'ETH_KNC', 'action': 'Buy', 'amount': '277.39167387', 'price': '0.00254032',
                 'nonce': 1526871035141,
                 'expireTimeSec': 1526874635, 'traderAddr': ''}
    private_key = '0xd5b39ee354b4e06f0829b55ce2af3a49693126593e6ed8c8f6c9e1f8a1781a0d'
    signature_hex = sign_order_msg(post_data, private_key)
    print('signature_hex:', signature_hex)


if __name__ == '__main__':
    play()
