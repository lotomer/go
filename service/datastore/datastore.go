package datastore

import (
	"fmt"
	"net/http"

	"../../http/router"
	"github.com/julienschmidt/httprouter"
)

func init() {
	fmt.Println("Start init crawler2...")
	router.DefaultRouter.GET("/datastore/task2/:taskName", crawlTask2Handle)
}

func crawlTask2Handle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("crawl task2 handle...")
	fmt.Fprintf(w, "start %s %s", "crawler2", ps.ByName("taskName"))
}
