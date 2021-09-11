package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		data, error := ioutil.ReadAll(r.Body)
		if error != nil {
			http.Error(responseWriter, "An error occurred", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(responseWriter, "Hello %s", data)

	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("GoodBye")
	})

	http.ListenAndServe(":9090", nil)
}
