package com.dextop.api;

import org.bouncycastle.util.encoders.Hex;
import org.web3j.abi.datatypes.generated.Uint32;
import org.web3j.abi.datatypes.generated.Uint64;
import org.web3j.abi.datatypes.generated.Uint8;
import org.web3j.crypto.ECKeyPair;
import org.web3j.crypto.Sign;
import org.web3j.utils.Numeric;

import java.math.BigDecimal;
import java.nio.ByteBuffer;
import java.util.Arrays;
import java.util.HashMap;

public class Models {

    public static class Token {
        public int tokenCode;
        public String tokenId;
    }

    public static class MarketInfo {
        public String marketAddr;
        public Token[] stockTokens;
        public Token[] cashTokens;

        @Override
        public String toString() {
            return "MarketInfo{" +
                    "marketAddr='" + marketAddr + '\'' +
                    ", stockTokens=" + Arrays.toString(stockTokens) +
                    ", cashTokens=" + Arrays.toString(cashTokens) +
                    '}';
        }
    }

    public static long stringToE8(String input) {
        return new BigDecimal(input).multiply(new BigDecimal(1e8)).longValue();
    }

    public static long GetNonce() {
        return System.currentTimeMillis();
    }

    public static int GetExpireTimeStamp() {
        return (int) (GetNonce() / 1000L + 3600); // set expire time to 1 hour later
    }

    public static String sigToString(Sign.SignatureData sig) {
        String R = Numeric.toHexString(sig.getR(), 0, 32, false);
        String S = Numeric.toHexString(sig.getS(), 0, 32, false);
        String V = sig.getV() - 27 == 0 ? "00" : "01";
        return "0x" + R + S + V;
    }

    public static class DexOrder {
        public String TraderAddr;
        public String MarketAddr;
        public String PairInfo;
        public Uint32 PairId;
        public String ActionName;
        public Uint8 Action;
        public Uint8 IoC;
        public String Price;
        public Uint64 PriceE8;
        public String Amount;
        public Uint64 AmountE8;
        public Uint64 ExpireTimeSec;
        public Uint64 Nonce;

        public DexOrder(String traderAddr, String marketAddr, String pairInfo, int pairId, String action, String price, String amount) throws Exception {
            this.TraderAddr = traderAddr;
            this.MarketAddr = marketAddr;
            this.PairInfo = pairInfo;
            this.PairId = new Uint32(pairId);
            if (action.equals("Buy")) {
                this.ActionName = action;
                this.Action = new Uint8(0);
            } else if (action.equals("Sell")) {
                this.ActionName = action;
                this.Action = new Uint8(1);
            } else {
                throw new Exception("invalid action");
            }
            this.IoC = new Uint8(0); // only support 0 for now
            this.Price = price;
            this.Amount = amount;
            this.PriceE8 = new Uint64(stringToE8(price));
            this.AmountE8 = new Uint64(stringToE8(amount));
            this.ExpireTimeSec = new Uint64(GetExpireTimeStamp());
            this.Nonce = new Uint64(GetNonce());
        }

        public byte[] ToBytes() {
            byte[] bytes = new byte[98];
            ByteBuffer buf = ByteBuffer.wrap(bytes);
            buf.put(new byte[]{0x19});
            buf.put("Ethereum Signed Message:\n70DEx2 Order: ".getBytes());
            buf.put(Numeric.hexStringToByteArray(this.MarketAddr));
            buf.put(Numeric.toBytesPadded(this.Nonce.getValue(), 8));
            buf.put(Numeric.toBytesPadded(this.ExpireTimeSec.getValue(), 8));
            buf.put(Numeric.toBytesPadded(this.AmountE8.getValue(), 8));
            buf.put(Numeric.toBytesPadded(this.PriceE8.getValue(), 8));
            buf.put(Numeric.toBytesPadded(this.IoC.getValue(), 1));
            buf.put(Numeric.toBytesPadded(this.Action.getValue(), 1));
            buf.put(Numeric.toBytesPadded(this.PairId.getValue(), 4));
            System.out.println(Hex.toHexString(bytes));
            return bytes;
        }

        public HashMap<String, String> ToBody(String privKey) {
            HashMap<String, String> ret = new HashMap<String, String>();
            ret.put("pairId", this.PairInfo);
            ret.put("traderAddr", this.TraderAddr);
            ret.put("action", this.ActionName);
            ret.put("price", this.Price);
            ret.put("amount", this.Amount);
            ret.put("expireTimeSec", this.ExpireTimeSec.getValue().toString());
            ret.put("nonce", this.Nonce.getValue().toString());

            ECKeyPair keyPairKovan = ECKeyPair.create(Numeric.hexStringToByteArray(privKey));
            Sign.SignatureData sig = Sign.signMessage(this.ToBytes(), keyPairKovan);
            ret.put("sig", sigToString(sig));
            return ret;
        }
    }
}
