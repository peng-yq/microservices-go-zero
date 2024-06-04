package kqueue

// Third-party payment callback payment status change notification
type ThirdPaymentUpdatePayStatusNotifyMessage struct {
	PayStatus int64  `json:"payStatus"`
	OrderSn   string `json:"orderSn"`
}

