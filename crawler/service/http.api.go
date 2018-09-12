package crawler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lotomer/go/http/router"
)

var thisServiceURIRoot = "/crawl/task/"

func init() {
	crawlTaskURIPattern := thisServiceURIRoot + ":taskName"
	router.DefaultRouter.GET(crawlTaskURIPattern, crawlTaskHandle)
	log.Printf("Handle %s by %s", crawlTaskURIPattern, "crawlTaskHandle")
}

func crawlTaskHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("crawl task handle...")
	fmt.Fprintf(w, "start %s %s", "crawler", ps.ByName("taskName"))
}
