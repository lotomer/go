package main

import (

	//_ "net/http/pprof"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lotomer/go/application/common"
	_ "github.com/lotomer/go/datastore/service"
	_ "github.com/lotomer/go/privilege/service"
)

func main() {
	common.Main()
}
