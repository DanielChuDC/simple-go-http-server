package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)       // index function will be called when route /
	http.ListenAndServe(":8080", nil) //1
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from cloud native go") // diff with printf

}
