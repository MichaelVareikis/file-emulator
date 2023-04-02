package main

import (
	helper "filewafe/helpers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/inv/api/v1/query/151", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		fmt.Fprintf(w, helper.Query)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		log.Println(r.URL.Query())
		log.Println("Token", r.Header.Get("Authorization"))
		fmt.Fprintf(w, helper.Query_result)
	})

	log.Println("Start FileWave Api on port 8080")
	http.ListenAndServe(":8080", nil)
}
