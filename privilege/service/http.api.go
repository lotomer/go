package service

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mydb "github.com/lotomer/go/db"
	"github.com/lotomer/go/http/response"
	"github.com/lotomer/go/http/router"
	"github.com/lotomer/go/privilege"
)

var uri4reload = "/privilege/reload"

func init() {
	// 注册重新加载数据源、数据服务及权限数据功能
	router.DefaultRouter.GET(uri4reload, reloadHandle)
	log.Printf("Handle %s", uri4reload)
}

// 重新加载
func reloadHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
	case "user":
		err = privilege.InitUserAndRole(mydb.GetDB(mydb.ThisDataSourceID))
		if err != nil {
			response.FailJSON(w, err.Error())
		} else {
			response.SuccessJSON(w, "ok")
		}

	case "uri":
		err = privilege.InitURIPrivileges(mydb.GetDB(mydb.ThisDataSourceID))
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
