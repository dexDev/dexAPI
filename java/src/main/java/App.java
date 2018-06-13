import com.dextop.api.Models;
import com.dextop.api.Services;
import org.json.JSONObject;

public class App {
    public static void main(String[] args) {
        String dexTopUrl = "https://testnet271828.dex.top/v1";
        String privKeyKovan = "0xD3E463E0B7314C9943F0831DBE8C5C9B691A6E46ABCC72CB6A614B9425D246D3";
        String TraderAddr = "0xE89BB07333b3D792C29c744511bbCFAEE0DdD4b7";

        Services service = new Services(dexTopUrl);
        String authToken = service.Login("naituida@163.com", "dextoptest1");
        System.out.println(authToken);

        Models.MarketInfo marketInfo = service.GetMarketInfo();
        System.out.println(marketInfo);

        // ETH tokenCode = 0, YEE tokenCode = 300, ETH_YEE = int32(0) << 32 | int32(300) = 300

        try {
            Models.DexOrder order = new Models.DexOrder(
                    TraderAddr,
                    marketInfo.marketAddr,
                    "ETH_YEE",
                    300,
                    "Buy",
                    "0.0003",
                    "100"
            );
            JSONObject ret = service.PlaceOrder(order, privKeyKovan, authToken);
            System.out.println(ret.toString());
        } catch (Exception e) {
            System.out.println("failed to place order");
        }
    }
}
