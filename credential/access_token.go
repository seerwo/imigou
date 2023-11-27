package credential

//AccessTokenHandle AccessToken interface
type AccessTokenHandle interface {
	GetAccessToken()(accessToken string, err error)
	GetAccessParam(string)(string, error)
}
