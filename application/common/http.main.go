package common

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/lotomer/go/common"
	"github.com/lotomer/go/config"
	"github.com/lotomer/go/http/response"
	"github.com/lotomer/go/http/router"
	//_ "../../privilege/service"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	mydb "github.com/lotomer/go/db"
)

var httpPort = flag.Int("port", 8080, "Http port")
var nolog = flag.Bool("nolog", false, "Without log")
var help = flag.Bool("h", false, "Help info")
var configFile = flag.String("config", "", "The database config file(json)")
var accessControlAllowOrigin = flag.String("Access-Control-Allow-Origin", "", "The http header Access-Control-Allow-Origin")

//Main 提供给通用http服务入口
func Main(name string, fun func(*sql.DB)) {
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
		log.SetPrefix("[" + name + "] ")
	}
	// 设置全局变量
	common.GlobalConfig.AccessControlAllowOrigin = *accessControlAllowOrigin
	fileName := name + ".json"
	// 从配置文件读取数据库配置
	var configStr []byte
	if *configFile == "" {
		for _, fname := range []string{"./" + fileName, "./etc/" + fileName, "/etc/" + fileName, "../etc/" + fileName, "../../etc/" + fileName,
			"./config.json", "./etc/config.json", "/etc/config.json", "../etc/config.json", "../../etc/config.json"} {
			buff, err := ioutil.ReadFile(fname)
			if err != nil {
				log.Print(err)
			} else {
				log.Printf("Read config success from %s", fname)
				configStr = buff
				break
			}
		}
	} else {
		buff, err := ioutil.ReadFile(*configFile)
		if err != nil {
			panic(err)
		}
		log.Printf("Read config success from %s", *configFile)
		configStr = buff
	}
	if len(configStr) == 0 {
		log.Fatal("Read config failed!")
	}

	err := config.Config.Use(configStr)
	if err != nil {
		log.Fatal(err)
	}
	// 获取数据库连接
	db, err := mydb.GenerateDB()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	// defer func() {
	// 	for id, db := range datastore.DataSourcePool {
	// 		log.Printf("close db %d", id)
	// 		db.Close()
	// 	}
	// }()

	mydb.Use(db) // 保存主库

	if fun != nil {
		fun(db)
	}

	// 首先初始化数据商店，以便后续模块使用
	//datastore.Use(db)
	// 初始化权限，依赖数据商店
	//privilege.Use(db)

	router.DefaultRouter.GET("/", notFoundHandle)

	common.ProgramSignalHandle(func() {
		fmt.Println("开始退出...")
		fmt.Println("执行清理...")
		// for id, db := range mydb.DataSourcePool {
		// 	fmt.Printf("close db %d\n", id)
		// 	db.Close()
		// }
		fmt.Println("结束退出...")
		os.Exit(0)
	}, nil, nil)
	log.Printf("listening at: 0.0.0.0:%d", *httpPort)
	http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), router.DefaultRouter)
}
func notFoundHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 执行预处理
	if !response.BeforeProcessHandle(w, req) {
		return
	}

	response.SuccessJSON(w, "Not found haha")

}
