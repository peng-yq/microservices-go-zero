package homestay

import (
	"net/http"

	"microservices-go-zero/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"microservices-go-zero/app/travel/cmd/api/internal/logic/homestay"
	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
)

func GuessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GuessListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestay.NewGuessListLogic(r.Context(), svcCtx)
		resp, err := l.GuessList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
