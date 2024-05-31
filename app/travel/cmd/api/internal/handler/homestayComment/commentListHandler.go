package homestayComment

import (
	"net/http"

	"microservices-go-zero/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"microservices-go-zero/app/travel/cmd/api/internal/logic/homestayComment"
	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayComment.NewCommentListLogic(r.Context(), svcCtx)
		resp, err := l.CommentList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
