package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"log"
	"net/http"
	route "ryuzaki/app/route"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	go GetAllGamesInSteamEveryDay()
	handler := mux.NewRouter()
	handler.HandleFunc("/", Hello).Methods("GET")
	handler.HandleFunc("/api/steam", route.HelloSteam).Methods("POST")
	err := http.ListenAndServe("0.0.0.0:8080", handler)
	if err != nil {
		log.Fatalln("Listen and Serve problem..")
	}
}

func GetAllGamesInSteamEveryDay() {
	s := gocron.NewScheduler()
	s.Every(1).Day().Do(route.GetAllGamesInSteam)
	<-s.Start()
}
