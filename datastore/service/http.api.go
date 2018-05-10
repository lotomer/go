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

func init() {
	dataStoreURIPattern := thisServiceURIRoot + ":dataId"
	router.DefaultRouter.GET(dataStoreURIPattern, dataStoreHandle)
	log.Printf("Handle %s by %s", dataStoreURIPattern, "dataStoreHandle")
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
