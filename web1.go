package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	//var http = http
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>INDEX PAGE</h1>")
	}

	hiHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>HELLO WORLD</h1>\n")
	}

	//http.Serve(8080, nil)
	http.HandleFunc("/", hiHandler)
	http.HandleFunc("/**", hiHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server started on 8080")
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
