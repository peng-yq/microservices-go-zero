### 1. "login"

1. route definition

- Url: /usercenter/v1/user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. request definition



```golang
type LoginReq struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 2. "register"

1. route definition

- Url: /usercenter/v1/user/register
- Method: POST
- Request: `RegisterReq`
- Response: `RegisterResp`

2. request definition



```golang
type RegisterReq struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type RegisterResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 3. "get user info"

1. route definition

- Url: /usercenter/v1/user/detail
- Method: POST
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. request definition



```golang
type UserInfoReq struct {
}
```


3. response definition



```golang
type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type User struct {
	Id int64 `json:"id"`
	Mobile string `json:"mobile"`
	Nickname string `json:"nickname"`
	Sex int64 `json:"sex"`
	Avatar string `json:"avatar"`
	Info string `json:"info"`
}
```

### 4. "wechat mini program auth"

1. route definition

- Url: /usercenter/v1/user/wxMiniAuth
- Method: POST
- Request: `WXMiniAuthReq`
- Response: `WXMiniAuthResp`

2. request definition



```golang
type WXMiniAuthReq struct {
	Code string `json:"code"`
	IV string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}
```


3. response definition



```golang
type WXMiniAuthResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

