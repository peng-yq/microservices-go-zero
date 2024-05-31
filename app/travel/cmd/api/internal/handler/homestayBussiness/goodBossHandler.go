package homestayBussiness

import (
	"net/http"

	"microservices-go-zero/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"microservices-go-zero/app/travel/cmd/api/internal/logic/homestayBussiness"
	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
)

func GoodBossHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodBossReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBussiness.NewGoodBossLogic(r.Context(), svcCtx)
		resp, err := l.GoodBoss(&req)
		result.HttpResult(r, w, resp, err)
	}
}
