package main

import (
	"database/sql"

	"github.com/lotomer/go/application/common"
	"github.com/lotomer/go/datastore"
	"github.com/lotomer/go/privilege"
	//_ "net/http/pprof"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lotomer/go/crawler/service"
	_ "github.com/lotomer/go/privilege/service"
)

const (
	// VERSION 版本号
	VERSION = "unknown"
)

func main() {
	common.Main("crawler", VERSION, func(db *sql.DB) {
		datastore.Use(db)
		privilege.Use(db)
	})
}
