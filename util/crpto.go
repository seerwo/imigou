package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
	"strings"
)

const(
	SignTypeMD5 = `MD5`
	SignTypeHMACSHA256 = `HMAC-SHA256`
)

//CalculateSign calculate sign
func CalculateSign(content, signType, key string) (string, error){
	var h hash.Hash
	if signType == SignTypeHMACSHA256 {
		h = hmac.New(sha256.New, []byte(key))
	} else {
		h = md5.New()
	}

	if _, err := h.Write([]byte(content)); err != nil {
		return ``, err
	}
	return strings.ToLower(hex.EncodeToString(h.Sum(nil))),  nil
}

//ParamSign calculate param sign
func ParamSign(p map[string]string, key string) (string, error){
	str := OrderParam(p, key)

	var signType string
	switch p["sign_type"] {
	case SignTypeMD5, SignTypeHMACSHA256:
		signType = p["sign_type"]
	case ``:
		signType = SignTypeMD5
	default:
		return ``, errors.New(`invalid sign_type`)
	}
	return CalculateSign(str, signType, key)
}