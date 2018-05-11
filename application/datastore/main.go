package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//_ "net/http/pprof"
	"../../common"
	"../../datastore"
	_ "../../datastore/service"
	"../../http/router"
	"../../privilege"
	_ "../../privilege/service"
	_ "github.com/go-sql-driver/mysql"
)

var nolog = flag.Bool("nolog", false, "Without log")
var help = flag.Bool("h", false, "Help info")
var dbfile = flag.String("dbfile", "", "The database config file(json)")

func main() {
	// 解析命令行参数
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// 根据命令行参数判断是否启用log
	if *nolog {
		var noneWriter common.NoneWriter
		log.SetOutput(&noneWriter)
	} else {
		log.SetPrefix("[DataStore] ")
	}
	// 从配置文件读取数据库配置
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
	// 获取数据库连接
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

	// 首先初始化数据商店，以便后续模块使用
	datastore.Use(db)
	// 初始化权限，依赖数据商店
	privilege.Use(db)

	http.ListenAndServe(":8080", router.DefaultRouter)
}
