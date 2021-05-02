
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/",index) // index function will be called when route //
	http.ListenAndServe(port(),nil) 
}

func port() string {
	// read from env
	port := os.Getenv("PORT")
	if len(port)==0 {
		port = "8080"
	}
	return  ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.StatusOK)
	 fmt.Fprintf(w, "Hello from cloud native go") // diff with printf

}