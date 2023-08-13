package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
