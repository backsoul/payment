package utils

import "encoding/json"

func ParseResponseJson(res interface{}) []byte {
	js, _ := json.Marshal(res)
	return js
}
