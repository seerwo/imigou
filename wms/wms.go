package wms

import (
	"github.com/seerwo/imigou/credential"
	"github.com/seerwo/imigou/wms/config"
	"github.com/seerwo/imigou/wms/context"
	"github.com/seerwo/imigou/wms/oauth"
	"github.com/seerwo/imigou/wms/order"
)

//Wms wms relate api
type Wms struct {
	ctx *context.Context
}

//NewWms instance api
func NewWms(cfg *config.Config) *Wms {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyWmsPrefix, cfg.Cache)
	ctx := &context.Context{
		Config: cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Wms{ctx:ctx}
}

//SetAccessTokenHandle custom access_token get method
func(wms *Wms) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle){
	wms.ctx.AccessTokenHandle = accessTokenHandle
}

//GetContext get Context
func (wms *Wms) GetContext() *context.Context {
	return wms.ctx
}

//GetAcessToken get access_token
func (wms *Wms) GetAccessToken()(string, error){
	return wms.ctx.GetAccessToken()
}


//GetWms
func (wms *Wms) GetOrder() *order.Order {
	return order.NewOrder(wms.ctx)
}
