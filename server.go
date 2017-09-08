package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Hander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Grrrrlrah!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hander)
	log.Fatal(http.ListenAndServe(":8000", r))
}
