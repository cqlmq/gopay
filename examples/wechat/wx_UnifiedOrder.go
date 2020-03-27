package wechat

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cqlmq/sycpay"
	"github.com/cqlmq/sycpay/wechat"
)

func UnifiedOrder() {
	//初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	//设置国家
	client.SetCountry(wechat.China)

	number := sycpay.GetRandomString(32)
	fmt.Println("out_trade_no:", number)

	//初始化参数Map
	bm := make(sycpay.BodyMap)
	bm.Set("nonce_str", sycpay.GetRandomString(32))
	bm.Set("body", "H5支付")
	bm.Set("out_trade_no", number)
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "http://www.sycpay.ink")
	bm.Set("trade_type", wechat.TradeType_H5)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", wechat.SignType_MD5)

	sceneInfo := make(map[string]map[string]string)
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "http://www.sycpay.ink"
	h5Info["wap_name"] = "H5测试支付"
	sceneInfo["h5_info"] = h5Info
	bm.Set("scene_info", sceneInfo)

	//body.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	// 正式
	//sign := sycpay.GetWeChatParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	// 沙箱
	//sign, _ := sycpay.GetWeChatSanBoxParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	//body.Set("sign", sign)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(bm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp:", *wxRsp)
	//fmt.Println("wxRsp.MwebUrl:", wxRsp.MwebUrl)

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	//获取小程序支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := wechat.GetMiniPaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, wechat.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	//获取H5支付需要的paySign
	pac := "prepay_id=" + wxRsp.PrepayId
	paySign := wechat.GetH5PaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, wechat.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	fmt.Println("paySign:", paySign)

	//获取小程序需要的paySign
	//paySign := wechat.GetAppPaySign("wxdaa2ab9ef87b5497","", wxRsp.NonceStr, wxRsp.PrepayId, wechat.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)
}
