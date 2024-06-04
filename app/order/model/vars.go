package model

import (
    "errors"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrNoRowsUpdate = errors.New("update db no rows change")

// HomestayOrder Transaction Status: -1: Cancelled 0: Waiting for payment 1: Unused 2: Used 3: Refunded 4: Expired
// note: order already be paied is unused; but i think there should has a paid state
var HomestayOrderTradeStateCancel int64 = -1
var HomestayOrderTradeStateWaitPay int64 = 0
var HomestayOrderTradeStateWaitUse int64 = 1
var HomestayOrderTradeStateUsed int64 = 2
var HomestayOrderTradeStateRefund int64 = 3
var HomestayOrderTradeStateExpire int64 = 4

// food
var HomestayOrderNeedFoodNo int64 = 0
var HomestayOrderNeedFoodYes int64 = 1