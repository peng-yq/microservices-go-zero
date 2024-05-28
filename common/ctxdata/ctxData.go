package ctxdata

import (
	"context"
	"encoding/json"
	
	"github.com/zeromicro/go-zero/core/logx"
)

// Keys in context (context.Context). 
// This key is used to access the user ID associated with JWT in the context
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}
