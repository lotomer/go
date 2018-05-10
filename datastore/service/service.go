package datastore

import (
	"log"
	"net/http"

	"../../common"
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

func dataStoreHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	dataID := ps.ByName("dataId")

	if config, ok := datastore.DataConfigPool[dataID]; ok {
		// 找到了该API，则继续
		key := req.URL.Query().Get("key")
		log.Printf("start %s %s, key=%s", "datastore", dataID, key)
		// 1、校验key
		user, err := privilege.GetUserByKey(key)
		if err != nil {
			response.FailJSON(w, err.Error())
			return
		}
		// 2、校验key是否有该API权限
		if err = privilege.CheckAPIPrivilege(user, thisServiceURIRoot+dataID); err != nil {
			response.FailJSON(w, err.Error())
			return
		}

		// 3、校验查询参数
		var queryParam = make(map[string]string)
		for _, param := range config.QueryParam {
			queryParam[param.Name] = req.URL.Query().Get(param.Name)
			if param.Required == "true" && queryParam[param.Name] == "" {
				response.FailJSON(w, "缺少必备参数："+param.Name)
				return
			}
		}
		// 4、执行查询
		sql, db := config.Options.SQL, config.DB
		cols, datas, err := common.LoadDatasFromDB(db, sql)
		if err != nil {
			response.FailJSON(w, "查询失败："+err.Error())
			return
		}
		log.Printf("Execute %s's SQL: %s", dataID, sql)

		// 5、校验返回结果
		returnCols := make(map[string]int)
		for _, v := range config.Returns {
			returnCols[v.Name] = -1
		}
		for i := range cols {
			if _, ok := returnCols[cols[i]]; ok {
				// 需要返回，则提取字段值，设置为索引号
				returnCols[cols[i]] = i
			}
		}
		resultDatas := []map[string]interface{}{}
		for i := range datas {
			data := make(map[string]interface{})
			for k, v := range returnCols {
				data[k] = datas[i][v]
			}
			resultDatas = append(resultDatas, data)
		}
		//var resultDatas = make([]map[string]interface{}, datas.Len())
		response.SuccessJSON(w, resultDatas)
	} else {
		log.Printf("API不存在：%s", dataID)
		response.FailJSON(w, "API不存在："+dataID)
	}

}
