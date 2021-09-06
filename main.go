package main

import (
	"net/http"
	core_app "ryuzaki/app/core"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/hello", core_app.SayHello)
	err := http.ListenAndServe("0.0.0.0:8080", handler)
	if err != nil {
		return
	}
}
