package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var listenOn = "localhost:8800"

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
	log.Printf("Listening on %q ...", listenOn)
	log.Fatal(http.ListenAndServe(listenOn, router))
}
