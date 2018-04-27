package crawler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lotomer/go/http/router"
)

func init() {
	fmt.Println("Start init crawler2...")
	router.DefaultRouter.GET("/crawl/task2/:taskName", crawlTask2Handle)
}

func crawlTask2Handle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("crawl task2 handle...")
	fmt.Fprintf(w, "start %s %s", "crawler2", ps.ByName("taskName"))
}
