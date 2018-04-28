package datastore

import (
	"fmt"
	"log"
	"net/http"

	"../../datastore"
	"../../http/response"
	"../../http/router"
	"github.com/julienschmidt/httprouter"
)

func init() {
	dataStoreURIPattern := "/datastore/:dataId"
	router.DefaultRouter.GET(dataStoreURIPattern, dataStoreHandle)
	log.Printf("Handle %s by %s", dataStoreURIPattern, "dataStoreHandle")
}

func dataStoreHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("datastore handle...")
	dataID := ps.ByName("dataId")

	if config, ok := datastore.DataConfigPool[dataID]; ok {
		// 找到了该API，则继续
		key := req.URL.Query().Get("key")
		// 1、校验key
		// 2、校验key是否有该API权限
		// 3、校验查询参数
		// 4、执行查询
		// 5、校验返回结果

		// 3、校验查询参数
		var queryParam = make(map[string]string)
		for _, param := range config.QueryParam {
			queryParam[param.Name] = req.URL.Query().Get(param.Name)
			if param.Required == "true" && queryParam[param.Name] == "" {
				response.FailJSON(w, "缺少必备参数："+param.Name)
				return
			}
		}
		fmt.Fprintln(w, config)
		fmt.Fprintf(w, "start %s %s, key=%s", "datastore", dataID, key)
	}

	fmt.Fprintf(w, "start %s %s", "datastore", dataID)
}
