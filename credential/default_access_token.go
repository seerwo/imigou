package credential

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/seerwo/imigou/cache"
	"github.com/seerwo/imigou/util"
	"strconv"
	"strings"
	"sync"
	"time"
)

const(
	//AccessTokenURL get access_token interface
	accessTokenURL = ""
	//CacheKeyWmsPrefix set cache key prefix
	CacheKeyWmsPrefix = "go_imigou_wms_"
)

//DefaultAccessToken default AccessToken to get
type DefaultAccessToken struct {
	appID string
	appSecret string
	cacheKeyPrefix string
	cache cache.Cache
	accessTokenLock *sync.Mutex
}

//NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(appID, appSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is inneed")
	}
	return &DefaultAccessToken{
		appID: appID,
		appSecret: appSecret,
		cache:cache,
		cacheKeyPrefix: cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

//ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
}

//GetAccessToken get access_token, get cache
func (ak *DefaultAccessToken) GetAccessToken()(accessToken string, err error){
	//lock
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.appID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	//cache invalid
	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(ak.appID, ak.appSecret)
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	if err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second); err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

func (ak *DefaultAccessToken) GetAccessParam(req string)(accessParam string, err error){

	ap := AccessParam{}
	ap.MsgType = ""
	ap.MsgId = util.GetId()
	ap.UserCode = ak.appID

	encodeText := base64.StdEncoding.EncodeToString([]byte(req))
	//fmt.Printf("Encoded text: %s\n", encodeText)

	resource := encodeText + ak.appID + ak.appSecret
	md5Text,_ := util.CalculateSign(resource, "", "")

	ap.MsgData = encodeText
	ap.MsgDigest = md5Text

	jsonBytes, err := xml.Marshal(ap)
	if err != nil {
		return
	}

	accessParam = strings.ReplaceAll(req, "<request>", "<request>" + string(jsonBytes))
	return
}

//GetTokenFromServer reforce from server
func GetTokenFromServer(appID,appSecret string) (resAccessToken ResAccessToken, err error){
	url := fmt.Sprintf("%s?grant_type=client_redential&appid-%s&secret=%s", accessTokenURL, appID, appSecret)
	var body []byte
	if body, err = util.HTTPGet(url); err != nil {
		return
	}
	if err = json.Unmarshal(body, &resAccessToken); err != nil {
		return
	}
	if resAccessToken.Message != ""{
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.Code, resAccessToken.Message)
	}
	return
}
