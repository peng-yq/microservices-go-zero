package jobtype

import "microservices-go-zero/app/order/model"

// DeferCloseHomestayOrderPayload defer close homestay order
type DeferCloseHomestayOrderPayload struct {
	Sn string
}

// PaySuccessNotifyUserPayload pay success notify user
type PaySuccessNotifyUserPayload struct {
	Order *model.HomestayOrder
}
