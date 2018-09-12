package main

import (
	"../common"
	//_ "net/http/pprof"
	_ "../../crawler/service"
	_ "../../privilege/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	common.Main()
}
