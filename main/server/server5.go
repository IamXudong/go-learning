package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	handle := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		for key, val := range r.Form {
			fmt.Printf("%v = %v\n", key, val)
		}
	}

	http.HandleFunc("/", handle)

	port := "8000"
	if len(os.Args[1:]) > 0 {
		port = os.Args[1]
	}
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
