package user

import (
	"net/http"

	"microservices-go-zero/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"microservices-go-zero/app/usercenter/cmd/api/internal/logic/user"
	"microservices-go-zero/app/usercenter/cmd/api/internal/svc"
	"microservices-go-zero/app/usercenter/cmd/api/internal/types"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
