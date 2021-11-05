package tools

import (
	"encoding/json"
	"fmt"
)

func InList(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func GetFmtStr(data interface{}) string {
	resp, _ := json.Marshal(data)
	respStr := string(resp)
	if respStr == "" {
		respStr = fmt.Sprintf("%+v", data)
	}
	return respStr
}
