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

func main() {
	common.Main("datastore", func(db *sql.DB) {
		datastore.Use(db)
		privilege.Use(db)
	})
}
