package main

import (
	"database/sql"

	//_ "net/http/pprof"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lotomer/go/application/common"
	"github.com/lotomer/go/datastore"
	_ "github.com/lotomer/go/datastore/service"
	"github.com/lotomer/go/privilege"
	_ "github.com/lotomer/go/privilege/service"
)

const (
	// VERSION 版本号
	VERSION = "unknown"
)

func main() {
	common.Main("datastore", VERSION, func(db *sql.DB) {
		datastore.Use(db)
		privilege.Use(db)
	})
}
