package svc

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"microservices-go-zero/app/mqueue/cmd/job/internal/config"
)

// wechat miniprogram 
func newMiniprogramClient(c config.Config) *miniprogram.MiniProgram {
	return  wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     c.WxMiniConf.AppId,
		AppSecret: c.WxMiniConf.Secret,
		Cache:     cache.NewMemory(),
	})
}

