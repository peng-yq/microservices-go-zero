// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	thirdPayment "microservices-go-zero/app/payment/cmd/api/internal/handler/thirdPayment"
	"microservices-go-zero/app/payment/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// third payment：wechat pay callback
				Method:  http.MethodPost,
				Path:    "/thirdPayment/thirdPaymentWxPayCallback",
				Handler: thirdPayment.ThirdPaymentWxPayCallbackHandler(serverCtx),
			},
		},
		rest.WithPrefix("/payment/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// third payment：wechat pay
				Method:  http.MethodPost,
				Path:    "/thirdPayment/thirdPaymentWxPay",
				Handler: thirdPayment.ThirdPaymentwxPayHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/payment/v1"),
	)
}
