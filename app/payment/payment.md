### 1. "third payment：wechat pay callback"

1. route definition

- Url: /payment/v1/thirdPayment/thirdPaymentWxPayCallback
- Method: POST
- Request: `ThirdPaymentWxPayCallbackReq`
- Response: `ThirdPaymentWxPayCallbackResp`

2. request definition



```golang
type ThirdPaymentWxPayCallbackReq struct {
}
```


3. response definition



```golang
type ThirdPaymentWxPayCallbackResp struct {
	ReturnCode string `json:"return_code"`
}
```

### 2. "third payment：wechat pay"

1. route definition

- Url: /payment/v1/thirdPayment/thirdPaymentWxPay
- Method: POST
- Request: `ThirdPaymentWxPayReq`
- Response: `ThirdPaymentWxPayResp`

2. request definition



```golang
type ThirdPaymentWxPayReq struct {
	OrderSn string `json:"orderSn"`
	ServiceType string `json:"serviceType"`
}
```


3. response definition



```golang
type ThirdPaymentWxPayResp struct {
	Appid string `json:"appid"`
	NonceStr string `json:"nonceStr"`
	PaySign string `json:"paySign"`
	Package string `json:"package"`
	Timestamp string `json:"timestamp"`
	SignType string `json:"signType"`
}
```

