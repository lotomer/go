package main

import (
	"net/http"

	_ "../../crawler"
	"../../http/router"
)

func main() {
	http.ListenAndServe(":8080", router.DefaultRouter)
}
