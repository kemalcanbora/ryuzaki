package steam

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	helpers "ryuzaki/helpers"
	schema "ryuzaki/helpers/schema"
)

func Parser(steamId string) []byte {
	helpers.GetEnv()
	url := fmt.Sprintf("https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json",
		os.Getenv("STEAM_API_KEY"),
		steamId)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to URL:", err)
	}
	var steamSchema schema.SteamMapping
	steamErr := json.NewDecoder(resp.Body).Decode(&steamSchema)
	if steamErr != nil {
		return []byte("Steam Mapping Error")
	}

	bytes, _ := json.Marshal(steamSchema.Response)
	return bytes
}

func GetAllGames() []byte {
	resp, err := http.Get("http://api.steampowered.com/ISteamApps/GetAppList/v0002/")
	if err != nil {
		log.Fatal("Game Parser Problem")
	}
	var steamGames schema.SteamGetAllGames
	steamErr := json.NewDecoder(resp.Body).Decode(&steamGames)

	if steamErr != nil {
		log.Fatal("Steam Mapping Error")
	}

	bytes, _ := json.Marshal(steamGames.Applist)
	return bytes
}
