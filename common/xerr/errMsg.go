package xerr

var message map[uint32]string

// Init the errmsg corresponding to the errcode
func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "The server is out of service. Please try again later."
	message[REUQEST_PARAM_ERROR] = "Parameters error."
	message[TOKEN_EXPIRE_ERROR] = "Token is invalid, please log in again."
	message[TOKEN_GENERATE_ERROR] = "Failed to generate token."
	message[DB_ERROR] = "Database is busy, please try again later."
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "The number of rows affected by updating data is 0."
}

// Return the errmsg specified to the errcode
func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "The server is out of service. Please try again later."
	}
}

// Check the errcode
func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
