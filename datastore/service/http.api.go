package service

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lotomer/go/datastore"
	mydb "github.com/lotomer/go/db"
	"github.com/lotomer/go/http/response"
	"github.com/lotomer/go/http/router"
	"github.com/lotomer/go/privilege"
)

var thisServiceURIRoot = "/datastore/service/"
var uri4reload = "/datastore/reload"

func init() {
	// 注册API服务
	dataStoreURIPattern := thisServiceURIRoot + ":dataId"
	router.DefaultRouter.GET(dataStoreURIPattern, dataStoreGetHandle)
	router.DefaultRouter.POST(dataStoreURIPattern, dataStorePostHandle)
	log.Printf("Handle %s", dataStoreURIPattern)

	// 注册重新加载数据源、数据服务及权限数据功能
	router.DefaultRouter.GET(uri4reload, dataStoreReloadHandle)
	log.Printf("Handle %s", uri4reload)
}

// 重新加载
func dataStoreReloadHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
		err = datastore.InitDataSource(mydb.GetDB(mydb.ThisDataSourceID))
		if err != nil {
			response.FailJSON(w, err.Error())
		} else {
			response.SuccessJSON(w, "ok")
		}
	case "dataconfig":
		err = datastore.InitDataConfig(mydb.GetDB(mydb.ThisDataSourceID))
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

// http://localhost:8080/datastore/service/getUserByKey?KEY=ed14cf6d-d41d-48e4-806d-a9431baa9b46&key=ed14cf6d-d41d-48e4-806d-a9431baa9b46
func dataStoreGetHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 执行预处理
	if !response.BeforeProcessHandle(w, req) {
		return
	}
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
	inputParams := make(map[string][]string)
	for k, v := range req.URL.Query() {
		inputParams[k] = v
	}
	resultDatas, err := datastore.GetAPIDatas(dataID, inputParams)
	if err != nil {
		response.FailJSON(w, err.Error())
		return
	}
	response.SuccessJSON(w, resultDatas)
}

func dataStorePostHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 执行预处理
	if !response.BeforeProcessHandle(w, req) {
		return
	}
	dataID := ps.ByName("dataId")
	key := req.URL.Query().Get("key")
	req.ParseForm()
	inputParams := make(map[string][]string)
	for k, v := range req.PostForm {
		inputParams[k] = v
	}
	if key == "" {
		// url参数中没有key，则从post参数中提取
		if vs, ok := inputParams["key"]; ok && len(vs) > 0 {
			key = vs[0]
		}

	}
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

	resultDatas, err := datastore.GetAPIDatas(dataID, inputParams)
	if err != nil {
		response.FailJSON(w, err.Error())
		return
	}
	response.SuccessJSON(w, resultDatas)
}
