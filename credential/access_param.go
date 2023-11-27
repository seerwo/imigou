package credential

type AccessParam struct {
	MsgType string `json:"msgType"` //
	UserCode string `json:"userCode"`
	MsgId string `json:"msgId"`
	MsgData string `json:"msgData"`
	MsgDigest string `json:"msgDigest"`
}

func NewAccessParam() AccessParam {
	accessParam := *new(AccessParam)
	return accessParam
}