package main

import (
	"github.com/gorilla/mux"
	"net/http"
	route "ryuzaki/app/route"
)

func main() {
	handler := mux.NewRouter()
	handler.HandleFunc("/api/steam", route.HelloSteam).Methods("POST")
	err := http.ListenAndServe("0.0.0.0:8080", handler)
	if err != nil {
		return
	}
}
