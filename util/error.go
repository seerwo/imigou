package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//CommonError return common json
type CommonError struct {
	Flag string  `json:"flag"`
	Code int64 `json:"code"`
	Message string `json:"message"`

}

//DecodeWithCommonError return value commonError
func DecodeWithCommonError(response []byte, apiName string) (err error){
	var commError CommonError
	if err = json.Unmarshal(response, &commError); err != nil {
		return
	}
	if commError.Flag != "success" {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, commError.Code, commError.Message)
	}
	return nil
}

//DecodeWithError return value from json
func DecodeWithError(response []byte, obj interface{}, apiName string) error {
	if err := json.Unmarshal(response, obj); err != nil {
		return fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		return fmt.Errorf("obj is invalid")
	}
	commonError := responseObj.Elem().FieldByName("CommonError")
	if !commonError.IsValid() || commonError.Kind() != reflect.Struct {
		return fmt.Errorf("commonError is invalid or not struct")
	}
	errCode := commonError.FieldByName("ErrCode")
	errMsg := commonError.FieldByName("ErrMsg")
	if !errCode.IsValid() || !errMsg.IsValid() {
		return fmt.Errorf("errcode or errmsg is invalid")
	}
	if errCode.Int() != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, errCode.Int(), errMsg.String())
	}
	return nil
}