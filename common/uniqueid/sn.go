package uniqueid

import (
	"fmt"
	"time"

	"microservices-go-zero/common/tool"
)

type SnPrefix string

const (
	SN_PREFIX_HOMESTAY_ORDER SnPrefix = "HSO" // looklook_order/homestay_order
	SN_PREFIX_THIRD_PAYMENT  SnPrefix = "PMT" // looklook_payment/third_payment
)

// generate order number
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), tool.Krand(8, tool.KC_RAND_KIND_NUM))
}
