package main

import (
	"log"
	"net/http"
	"os"

	//_ "../../crawler/service"
	"flag"

	"../../datastore"
	_ "../../datastore/service"
	"../../http/router"
	_ "github.com/go-sql-driver/mysql"
)

// NoneWriter 空
type NoneWriter struct {
}

func (w *NoneWriter) Write(p []byte) (int, error) {
	return 0, nil
}

var nolog = flag.Bool("nolog", false, "Without log")
var help = flag.Bool("h", false, "Help info")

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if *nolog {
		var noneWriter NoneWriter
		log.SetOutput(&noneWriter)
	} else {
		log.SetPrefix("[DataStore] ")
	}
	dbStr := `{
		"port":3306,
		"host":"vps2.tomstools.org",
		"dbname":"of",
		"username":"mysql",
		"password":"mysql123",
		"maxPoolSize":10,
		"maxIdleSize":2,
		"type":"mysql",
		"urlTemplate":"${username}:${password}@tcp(${host}:${port})/${dbname}?charset=utf8"
		}`

	db, err := datastore.GenerateDBWithJSONStr(dbStr)
	if err != nil {
		log.Fatal(err)
		os.Exit(-2)
	}
	defer db.Close()
	defer func() {
		for id, db := range datastore.DataSourcePool {
			log.Printf("close db %d", id)
			db.Close()
		}
	}()
	datastore.Use(db)
	log.Print(datastore.DataSources)
	log.Print(datastore.DataSourcePool)
	log.Print(datastore.DataConfigs)
	http.ListenAndServe(":8080", router.DefaultRouter)
}
