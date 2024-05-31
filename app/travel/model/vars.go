package model

import (
    "errors"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrNoRowsUpdate = errors.New("update db no rows change")

// homestay activity type
var HomestayActivityPreferredType = "preferredHomestay" //优选民宿
var HomestayActivityGoodBusiType = "goodBusiness"       //最佳房东

// homestay activities on and off shelves
var HomestayActivityDownStatus int64 = 0 // activity close
var HomestayActivityUpStatus int64 = 1   // activity open