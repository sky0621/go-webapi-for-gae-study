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
	fmt.Println(w, "This is a root path.")
}
