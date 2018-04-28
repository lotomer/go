package crawler

import (
	"fmt"
	"log"
	"net/http"

	"../../http/router"
	"github.com/julienschmidt/httprouter"
)

func init() {
	crawlTaskURIPattern := "/crawl/task/:taskName"
	router.DefaultRouter.GET(crawlTaskURIPattern, crawlTaskHandle)
	log.Printf("Handle %s by %s", crawlTaskURIPattern, "crawlTaskHandle")
}

func crawlTaskHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("crawl task handle...")
	fmt.Fprintf(w, "start %s %s", "crawler", ps.ByName("taskName"))
}
