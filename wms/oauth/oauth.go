package oauth

import (
	"fmt"
	"github.com/seerwo/imigou/util"
	"github.com/seerwo/imigou/wms/context"
	"net/http"
	"net/url"
)

const(
	redirectOauthURL = "%s%s%s%s"
	WebAppRedirectOauthURL = "%s%s%s%s"
	accessTokenURL = ""
	refreshAccessTokenURL = ""
	userInfoURL = ""
	checkAccessTokenURL = ""
)

//Oauth save use oauth message
type Oauth struct {
	*context.Context
}

//NewOauth instance message
func NewOauth(context *context.Context) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	return auth
}

//GetRedirectURL get jump url
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, state string)(string, error){
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(redirectOauthURL, oauth.AppID, urlStr, scope, state), nil
}

//GetWebAppRedirectURL get web jump url
func(oauth *Oauth) GetWebAppRedirectURL(redirectURI, scope, state string) (string, error){
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(WebAppRedirectOauthURL, oauth.AppID, urlStr, scope, state), nil
}

//Redirect to jump web oauth
func (oauth *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, redirectURI, scope, state string) error {
	location, err := oauth.GetRedirectURL(redirectOauthURL, scope, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, http.StatusFound)
	return nil
}

//ResAccessToken get oauth access_token message
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID string `json:"openid"`
	Scope string `json:"scope"`
}

//GetUserAccessToken
func (oauth *Oauth) GetUserAccessToken(code string)(result ResAccessToken, err error){
	return
}

//RefreshAccessToken refresh access_token
func (oauth *Oauth) RefreshAcessToken(refreshToken string)(result ResAccessToken, err error){
	return
}

//CheckAccessToken check access_token
func (oauth *Oauth) CheckAccessToken(accessToken, openID string) (b bool, err error){
	return
}