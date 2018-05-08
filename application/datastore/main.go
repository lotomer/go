package main

import (
	"io/ioutil"
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
var dbfile = flag.String("dbfile", "", "The database config file(json)")

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
	datastore.Use(db)
	log.Print(datastore.DataSources)
	log.Print(datastore.DataSourcePool)
	log.Print(datastore.DataConfigs)
	http.ListenAndServe(":8080", router.DefaultRouter)
}
