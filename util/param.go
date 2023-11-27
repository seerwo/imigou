package util

import (
	"bytes"
	"sort"
)

func UrlParam(p map[string]string, bizKey string) (returnStr string) {
	keys := make([]string, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if p[k] == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(p[k])
	}
	if buf.Len() > 0 {
		returnStr += bizKey + "?"
	}
	returnStr += buf.String()
	return
}

//OrderParam trade params
func OrderParam(p map[string]string, bizKey string) (returnStr string){
	keys := make([]string, 0, len(p))
	for k := range p {
		if k == "signed"{
			continue
		}
		keys = append(keys, k)
	}
	//sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if p[k] == ""{
			continue
		}
		if buf.Len() > 0 {
			//buf.WriteByte('&')
		}
		buf.WriteString(k)
		//buf.WriteByte('=')
		buf.WriteString(p[k])
	}
	if buf.Len() > 0 {
		returnStr += bizKey
	}

	buf.WriteString(bizKey)
	returnStr += buf.String()

	return
}