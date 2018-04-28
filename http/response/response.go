package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type resp struct {
	Status  uint8  `json:"status"`
	Message string `json:"message"`
}

// Fail 返回失败信息
func Fail(w http.ResponseWriter, errMessage string, formatType string) {
	if formatType != "xml" {
		FailJSON(w, errMessage)
	}
}

// FailJSON 以json返回失败信息
func FailJSON(w http.ResponseWriter, errMessage string) {
	re := resp{Status: 1, Message: errMessage}
	b, err := json.Marshal(re)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, string(b))
}
