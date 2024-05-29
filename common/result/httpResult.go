package result

import (
	"fmt"
	"net/http"

	"microservices-go-zero/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "The server is down, please try again later"

		causeErr := errors.Cause(err)                
		if e, ok := causeErr.(*xerr.CodeError); ok { 
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { 
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { 
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
	}
}

func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "The server is down, please try again later"

		causeErr := errors.Cause(err)                
		if e, ok := causeErr.(*xerr.CodeError); ok { 
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { 
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { 
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
	}
}

func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REUQEST_PARAM_ERROR, errMsg))
}
