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

func main() {
	common.Main("crawler", func(db *sql.DB) {
		datastore.Use(db)
		privilege.Use(db)
	})
}
