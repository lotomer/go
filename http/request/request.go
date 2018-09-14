package request

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// ContentType4json 用于JSON请求的HTTP请求头
const ContentType4json = "application/json"

// ContentType4post 用户POST/FORM请求的常用请求头
const ContentType4post = "application/x-www-form-urlencoded"

// Get http get页面
func Get(url string) (string, error) {
	return HTTPDo("GET", url, "", nil)
}

//HTTPDo 根据指定参数执行http请求并返回结果
func HTTPDo(method string, url string, body string, headers map[string]string) (string, error) {
	var bodyReader *strings.Reader
	if body != "" {
		bodyReader = strings.NewReader(body)
	} else {
		bodyReader = nil
	}
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return "", err
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
		return "", nil
	}
	defer resp.Body.Close()
	// 读取内容
	content, err := ioutil.ReadAll(strings.NewReader(body))
	if err != nil {
		return "", nil
	}
	return string(content), nil
}
