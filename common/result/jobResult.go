package result

import (
	"context"

	"microservices-go-zero/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

func JobResult(ctx context.Context, resp interface{}, err error) {
	if err == nil {
		if resp != nil {
			logx.Infof("resp: %+v", resp)
		}
		return
	} else {
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := "The server is down, please try again later"

		causeErr := errors.Cause(err)                
		if e, ok := causeErr.(*xerr.CodeError); ok { 
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { 
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { 
					errCode = grpcCode
					errMsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(ctx).Errorf("【JOB-ERR】 : %+v ,errCode:%d , errMsg:%s ", err, errCode, errMsg)
		return
	}
}
