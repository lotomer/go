package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "../../crawler/service"
	"../../datastore"
	_ "../../datastore/service"
	"../../http/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
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
	fmt.Println(datastore.DataSources)
	fmt.Println(datastore.DataSourcePool)
	fmt.Println(datastore.DataConfigs)
	http.ListenAndServe(":8080", router.DefaultRouter)
}
