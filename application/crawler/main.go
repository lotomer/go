package main

import (
	"github.com/lotomer/go/application/common"
	//_ "net/http/pprof"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lotomer/go/crawler/service"
	_ "github.com/lotomer/go/privilege/service"
)

func main() {
	common.Main("crawler")
}
