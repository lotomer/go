package main

import (

	//_ "net/http/pprof"
	_ "../../datastore/service"
	_ "../../privilege/service"
	"../common"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	common.Main()
}
