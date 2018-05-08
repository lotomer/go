package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type failResp struct {
	Status  uint8  `json:"status"`
	Message string `json:"message"`
}
type successResp struct {
	Status uint8       `json:"status"`
	Data   interface{} `json:"data"`
}

// Fail 返回失败信息
func Fail(w http.ResponseWriter, errMessage string, formatType string) {
	if formatType != "xml" {
		FailJSON(w, errMessage)
	}
}

// FailJSON 以json返回失败信息
func FailJSON(w http.ResponseWriter, errMessage string) {
	re := failResp{Status: 1, Message: errMessage}
	b, err := json.Marshal(re)
	if err != nil {
		FailJSON(w, err.Error())
		return
	}
	fmt.Fprint(w, string(b))
}

// SuccessJSON 以json返回结果数据
func SuccessJSON(w http.ResponseWriter, data interface{}) {
	re := successResp{Status: 0, Data: data}
	b, err := json.Marshal(re)
	if err != nil {
		FailJSON(w, err.Error())
		return
	}
	fmt.Fprint(w, string(b))
}
