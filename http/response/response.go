package response

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lotomer/go/common"
	"github.com/lotomer/go/http/request"
)

// Result 返回结果
type Result struct {
	Success bool        `json:"success"` // 状态
	Message string      `json:"message"` // 状态信息
	Data    interface{} `json:"data"`    // 结果数据
}

// ResultDatas 查询结果数据，包含页码及总数信息
type ResultDatas struct {
	request.Page `json:"page"`
	Total        int64       `json:"total"`
	Datas        interface{} `json:"datas"`
}

// type failResp struct {
// 	Status  uint8  `json:"status"`
// 	Message string `json:"message"`
// }
// type successResp struct {
// 	Status uint8       `json:"status"`
// 	Data   interface{} `json:"data"`
// }

// Output2json 将结果以json方式输出
func Output2json(w io.Writer, result Result) {
	b, err := json.Marshal(result)
	if err != nil {
		FailJSON(w, err.Error())
		return
	}
	fmt.Fprint(w, string(b))
}

// FailJSON 以json返回失败信息
func FailJSON(w io.Writer, errMessage string) {
	re := Result{Success: false, Message: errMessage}
	Output2json(w, re)
}

// SuccessJSON 以json返回结果数据
func SuccessJSON(w io.Writer, data interface{}) {
	re := Result{Success: true, Data: data}
	Output2json(w, re)
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

// ProcessResultWithPageInfo 处理带分页信息的结果数据
func ProcessResultWithPageInfo(w io.Writer, datas *ResultDatas, err error, pageNum, pageSize int) {
	if err != nil {
		FailJSON(w, err.Error())
		return
	}
	// 回填分页信息
	(*datas).PageNum = pageNum
	(*datas).PageSize = pageSize

	SuccessJSON(w, datas)
}

// ProcessResult 处理结果数据（不带分页）
func ProcessResult(w io.Writer, datas *ResultDatas, err error) {
	if err != nil {
		FailJSON(w, err.Error())
		return
	}

	SuccessJSON(w, datas)
}
