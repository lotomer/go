package crawler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lotomer/go/http/router"
)

func init() {
	fmt.Println("Start init crawler...")
	router.DefaultRouter.GET("/crawl/task/:taskName", crawlTaskHandle)
}

func crawlTaskHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("crawl task handle...")
	fmt.Fprintf(w, "start %s %s", "crawler", ps.ByName("taskName"))
}
