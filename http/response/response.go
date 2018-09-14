package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lotomer/go/common"
)

// ResultDatas 查询结果数据，包含页码及总数信息
type ResultDatas struct {
	Total    int64       `json:"total"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Datas    interface{} `json:"datas"`
}
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

// BeforeProcessHandle 执行之前调用，用户处理公共情况
func BeforeProcessHandle(w http.ResponseWriter, req *http.Request) bool {
	// 如果指定了具体的域名，则允许跨域
	if origin := common.GlobalConfig.AccessControlAllowOrigin; origin != "" && origin != "*" {
		w.Header().Set("Access-Control-Allow-Origin", origin) // 支持跨域
	}
	w.Header().Set("Access-Control-Allow-Origin", "*") // 支持跨域 测试
	return true
}
