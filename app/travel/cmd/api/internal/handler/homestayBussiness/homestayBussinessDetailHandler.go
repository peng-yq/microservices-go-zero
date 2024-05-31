package homestayBussiness

import (
	"net/http"

	"microservices-go-zero/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"microservices-go-zero/app/travel/cmd/api/internal/logic/homestayBussiness"
	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
)

func HomestayBussinessDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBussiness.NewHomestayBussinessDetailLogic(r.Context(), svcCtx)
		resp, err := l.HomestayBussinessDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
