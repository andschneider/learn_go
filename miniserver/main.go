package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld hit")
	fmt.Fprintf(w, "hello world! %s", r.URL.Path[1:])
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Print("ping hit")
	w.Write([]byte("pong"))
}

func main() {
	log.Print("running a little server!")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", helloworld)
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":3456", nil))
}
