package main

import (
	"github.com/gorilla/mux"
	"github.com/zemirco/route-template-testing/login"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login.GetLogin).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":9000", nil)
}
