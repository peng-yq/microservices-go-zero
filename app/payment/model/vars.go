package model

import (
    "errors"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrNoRowsUpdate = errors.New("update db no rows change")

// service type need to be paied
var ThirdPaymentServiceTypeHomestayOrder string = "homestayOrder" 

// thirdpayment model
var ThirdPaymentPayModelWechatPay = "WECHAT_PAY" 

// payment status
var ThirdPaymentPayTradeStateFAIL int64 = -1   // pay failed
var ThirdPaymentPayTradeStateWait int64 = 0    // need to be paied
var ThirdPaymentPayTradeStateSuccess int64 = 1 // pay success
var ThirdPaymentPayTradeStateRefund int64 = 2  // refunded