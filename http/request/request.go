package request

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Page 分页信息
type Page struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

// Meta 元数据信息
type Meta struct {
	Token string `json:"token"` // 令牌
	//UserName string `json:"userName"` // 用户名
}

// RequestParameter 请求参数
type RequestParameter struct {
	Meta Meta `json:"meta"`
	Page
	Params json.RawMessage `json:"params"`
}

// ContentType4json 用于JSON请求的HTTP请求头
const ContentType4json = "application/json"

// ContentType4post 用户POST/FORM请求的常用请求头
const ContentType4post = "application/x-www-form-urlencoded"

// Get http get页面
func Get(url string) (string, http.Header, error) {
	return HTTPDo("GET", url, "", nil)
}

// ParseReaderContent2map 将reader中的内容（json格式）解析成map
func ParseReaderContent2map(reader io.Reader) (*map[string]interface{}, error) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	param := make(map[string]interface{})
	if err = json.Unmarshal(body, &param); err != nil {
		return nil, err
	}

	return &param, nil

}

// Post http post请求
func Post(url string, body string) (string, http.Header, error) {
	headers := map[string]string{"Content-Type": "application/json"}
	return HTTPDo("POST", url, body, headers)
}

// PostForm http post请求
func PostForm(url string, body string) (string, http.Header, error) {
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	return HTTPDo("POST", url, body, headers)
}

//HTTPDo 根据指定参数执行http请求并返回结果
func HTTPDo(method string, url string, body string, headers map[string]string) (string, http.Header, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return "", nil, err
	}
	// 设置header
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	// 读取内容
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	return string(content), resp.Header, nil
}
