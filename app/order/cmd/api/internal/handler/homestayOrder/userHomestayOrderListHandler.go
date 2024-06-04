package homestayOrder

import (
	"net/http"

	"microservices-go-zero/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"microservices-go-zero/app/order/cmd/api/internal/logic/homestayOrder"
	"microservices-go-zero/app/order/cmd/api/internal/svc"
	"microservices-go-zero/app/order/cmd/api/internal/types"
)

func UserHomestayOrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserHomestayOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayOrder.NewUserHomestayOrderListLogic(r.Context(), svcCtx)
		resp, err := l.UserHomestayOrderList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
