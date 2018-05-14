package datastore

import (
	"log"
	"net/http"

	"../../datastore"
	"../../http/response"
	"../../http/router"
	"../../privilege"
	"github.com/julienschmidt/httprouter"
)

var thisServiceURIRoot = "/datastore/service/"
var uri4reload = "/datastore/reload"

func init() {
	// 注册API服务
	dataStoreURIPattern := thisServiceURIRoot + ":dataId"
	router.DefaultRouter.GET(dataStoreURIPattern, dataStoreHandle)
	log.Printf("Handle %s", dataStoreURIPattern)

	// 注册重新加载数据源、数据服务及权限数据功能
	router.DefaultRouter.GET(uri4reload, dataStoreReloadHandle)
	log.Printf("Handle %s", uri4reload)
}

// 重新加载
func dataStoreReloadHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 执行预处理
	if !response.BeforeProcessHandle(w, req) {
		return
	}
	key := req.URL.Query().Get("key")
	// 1、校验key
	user, err := privilege.GetUserByKey(key)
	if err != nil {
		response.FailJSON(w, err.Error())
		return
	}
	// 2、校验key是否有该API权限
	if err = privilege.CheckURIPrivilege(user, uri4reload); err != nil {
		response.FailJSON(w, err.Error())
		return
	}

	// 重新加载数据
	switch module := req.URL.Query().Get("module"); module {
	case "datasource":
		err = datastore.InitDataSource(datastore.DataSourcePool[datastore.ThisDataSourceID])
		if err != nil {
			response.FailJSON(w, err.Error())
		} else {
			response.SuccessJSON(w, "ok")
		}
	case "dataconfig":
		err = datastore.InitDataConfig(datastore.DataSourcePool[datastore.ThisDataSourceID])
		if err != nil {
			response.FailJSON(w, err.Error())
		} else {
			response.SuccessJSON(w, "ok")
		}
	case "":
		response.FailJSON(w, "Need parameter 'module'")
	default:
		response.FailJSON(w, "Invalid module '"+module+"'")
	}

}

// http://localhost:8080/datastore/service/getUserByKey?KEY=111&key=ed14cf6d-d41d-48e4-806d-a9431baa9b46
func dataStoreHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	dataID := ps.ByName("dataId")
	key := req.URL.Query().Get("key")
	// 找到了该API，则继续
	log.Printf("start %s %s, key=%s", "datastore", dataID, key)
	// 1、校验key
	user, err := privilege.GetUserByKey(key)
	if err != nil {
		response.FailJSON(w, err.Error())
		return
	}
	// 2、校验key是否有该API权限
	if err = privilege.CheckURIPrivilege(user, thisServiceURIRoot+dataID); err != nil {
		response.FailJSON(w, err.Error())
		return
	}
	inputParams := make(map[string]string)
	for k, v := range req.URL.Query() {
		inputParams[k] = v[0]
	}
	resultDatas, err := datastore.GetAPIDatas(dataID, inputParams)
	if err != nil {
		response.FailJSON(w, err.Error())
		return
	}
	response.SuccessJSON(w, resultDatas)
}
