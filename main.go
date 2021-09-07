package main

import (
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"net/http"
	route "ryuzaki/app/route"
)

func main() {
	go GetAllGamesInSteamEveryDay()
	handler := mux.NewRouter()
	handler.HandleFunc("/api/steam", route.HelloSteam).Methods("POST")
	err := http.ListenAndServe("0.0.0.0:8080", handler)
	if err != nil {
		return
	}
}

func GetAllGamesInSteamEveryDay() {
	s := gocron.NewScheduler()
	s.Every(1).Day().Do(route.GetAllGamesInSteam)
	<-s.Start()
}
