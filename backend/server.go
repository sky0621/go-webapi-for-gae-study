package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handleRoot)
	appengine.Main()
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[environment] is app engine: %v, is dev: %v\n", appengine.IsAppEngine(), appengine.IsDevAppServer())
}
