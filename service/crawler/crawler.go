package crawler

import (
	"fmt"
	"net/http"

	"../../http/router"
	"github.com/julienschmidt/httprouter"
)

func init() {
	fmt.Println("Start init crawler...")
	router.DefaultRouter.GET("/crawl/task/:taskName", crawlTaskHandle)
}

func crawlTaskHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("crawl task handle...")
	fmt.Fprintf(w, "start %s %s", "crawler", ps.ByName("taskName"))
}
