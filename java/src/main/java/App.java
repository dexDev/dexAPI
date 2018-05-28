import org.web3j.abi.datatypes.generated.Uint32;
import org.web3j.abi.datatypes.generated.Uint64;
import org.web3j.abi.datatypes.generated.Uint8;
import org.web3j.crypto.ECKeyPair;
import org.web3j.crypto.Sign;
import org.web3j.utils.Numeric;
import java.nio.ByteBuffer;

public class App {
    public static String sigToString(Sign.SignatureData sig) {
        String R = Numeric.toHexString(sig.getR(),0,32,false);
        String S = Numeric.toHexString(sig.getS(),0, 32,false);
        String V = sig.getV()-27==0?"00":"01";
        return "0x"+R+S+V;
    }

//    (*dex2.Order)(0xc420222d60)({
//        PairId: (uint32) 100,
//                Action: (uint8) 0,
//                Ioc: (uint8) 0,
//                PriceE8: (uint64) 10000000,
//                AmountE8: (uint64) 1500000000,
//                ExpireTimeSec: (uint64) 1527323455
//    })
    public static class DexOrder {
        public Uint32 PairId;
        public Uint8 Action;
        public Uint8 IoC;
        public Uint64 PriceE8;
        public Uint64 AmountE8;
        public Uint64 ExpireTimeSec;
        public Uint64 Nonce;
        public byte[] Market;


        public DexOrder(int pairId, int action, int ioc, long priceE8, long amountE8, long expireTimeSec, long nonce, String market) {
            this.PairId = new Uint32(pairId);
            this.Action = new Uint8(action);
            this.IoC = new Uint8(ioc);
            this.PriceE8 = new Uint64(priceE8);
            this.AmountE8 = new Uint64(amountE8);
            this.ExpireTimeSec = new Uint64(expireTimeSec);
            this.Nonce = new Uint64(nonce);
            this.Market = Numeric.hexStringToByteArray(market);
        }

        public byte[] ToBytes() {
            byte[] bytes = new byte[98];
            ByteBuffer buf = ByteBuffer.wrap(bytes);
            buf.put(new byte[]{0x19});
            buf.put("Ethereum Signed Message:\n70DEx2 Order: ".getBytes());
            buf.put(this.Market);
            buf.put(Numeric.toBytesPadded(this.Nonce.getValue(),8));
            buf.put(Numeric.toBytesPadded(this.ExpireTimeSec.getValue(),8));
            buf.put(Numeric.toBytesPadded(this.AmountE8.getValue(),8));
            buf.put(Numeric.toBytesPadded(this.PriceE8.getValue(),8));
            buf.put(Numeric.toBytesPadded(this.IoC.getValue(),1));
            buf.put(Numeric.toBytesPadded(this.Action.getValue(),1));
            buf.put(Numeric.toBytesPadded(this.PairId.getValue(),4));
            return bytes;
        }
    }

    public static void main(String[] args) {
        // create string to sign
        DexOrder order = new DexOrder(100,0,0,10000000,1500000000,1527323455, 1527319855888L,"0x4afF2E056D2fE5CeBed442Dc924eAeBe25909550");
        String stringToSign = Numeric.toHexString(order.ToBytes());
        System.out.println(stringToSign);
        // 0x19457468657265756d205369676e6564204d6573736167653a0a373044457832204f726465723a204aff2e056d2fe5cebed442dc924eaebe25909550000001639b5b8310000000005b091b3f0000000059682f000000000000989680000000000064

        // create ecdsa key pair
        String privKey = "0xd5b39ee354b4e06f0829b55ce2af3a49693126593e6ed8c8f6c9e1f8a1781a0d";
        ECKeyPair keyPair = ECKeyPair.create(Numeric.hexStringToByteArray(privKey));

        // create sig
        Sign.SignatureData sig = Sign.signMessage(Numeric.hexStringToByteArray(stringToSign), keyPair);
        System.out.println(sigToString(sig));
        // 0x16944b324f3a9b3b14d5eaee5a96b55b68da2eaf09b831f177ce98374b4e098f5729f3e774b7331422e6b40f420d588df37fd2f8279f80d3d37eaca83bdf26d600
    }
}
