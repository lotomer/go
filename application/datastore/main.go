package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"flag"

	"../../common"
	"../../privilege"

	"../../datastore"
	_ "../../datastore/service"
	"../../http/response"
	"../../http/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var nolog = flag.Bool("nolog", false, "Without log")
var help = flag.Bool("h", false, "Help info")
var dbfile = flag.String("dbfile", "", "The database config file(json)")

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if *nolog {
		var noneWriter common.NoneWriter
		log.SetOutput(&noneWriter)
	} else {
		log.SetPrefix("[DataStore] ")
	}
	var dbStr string
	if *dbfile == "" {
		for _, fname := range []string{"./db.json", "./etc/db.json", "/etc/db.json", "../etc/db.json", "../../etc/db.json"} {
			buff, err := ioutil.ReadFile(fname)
			if err != nil {
				log.Print(err)
			} else {
				log.Printf("Read database config success from %s", fname)
				dbStr = string(buff)
				break
			}
		}
	} else {
		buff, err := ioutil.ReadFile(*dbfile)
		if err != nil {
			panic(err)
		}
		log.Printf("Read database config success from %s", *dbfile)
		dbStr = string(buff)
	}
	if dbStr == "" {
		log.Fatal("Read database config failed!")
	}
	db, err := datastore.GenerateDBWithJSONStr(dbStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	defer func() {
		for id, db := range datastore.DataSourcePool {
			log.Printf("close db %d", id)
			db.Close()
		}
	}()
	initData(db)
	router.DefaultRouter.GET(uri4reload, reloadHandle)
	log.Printf("Handle %s by %s", uri4reload, "reloadHandle")
	http.ListenAndServe(":8080", router.DefaultRouter)
}

var uri4reload = "/reload"

func initData(db *sql.DB) {
	if db == nil {
		db = datastore.DataSourcePool[datastore.ThisDataSourceID]
	}
	// 首先初始化数据商店，以便后续模块使用
	datastore.Use(db)
	// 初始化权限，依赖数据商店
	privilege.Use(db)
	log.Print(datastore.DataSources)
	log.Print(datastore.DataSourcePool)
	log.Print(datastore.DataConfigs)
}

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
	initData(nil)
	response.SuccessJSON(w, "ok")
}
