package main

import (
	"net/http"

	_ "github.com/lotomer/go/crawler"
	"github.com/lotomer/go/http/router"
)

func main() {
	http.ListenAndServe(":8080", router.DefaultRouter)
}
