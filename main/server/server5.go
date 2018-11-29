package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string = ":8000"

func main() {
	http.HandleFunc("/", handle)
	if len(os.Args[1:]) > 0 {
		port = ":" + os.Args[1]
	}
	log.Fatal(http.ListenAndServe(port, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for key, val := range r.Form {
		fmt.Printf("%v = %v\n", key, val)
	}
	fmt.Fprintf(w, "%v\n", "Hello, world!")
}
