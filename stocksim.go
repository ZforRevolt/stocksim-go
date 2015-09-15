package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := "assets/index.html"
	if _, err := os.Stat(p); err == nil {
		http.ServeFile(w, r, p)
		return
	}
	http.NotFound(w, r)
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:8800", router))
}
