package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	steam_app "ryuzaki/app/core/steam"
	schema "ryuzaki/helpers/schema"
	kafka "ryuzaki/platform/streaming"
	"strconv"
)

func HelloSteam(w http.ResponseWriter, r *http.Request) {
	var steamSchema schema.SteamReqBody
	err := json.NewDecoder(r.Body).Decode(&steamSchema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	app := steam_app.Parser(steamSchema.UserId)
	kafka.Producer("jungle", app, []byte(steamSchema.UserId))
	fmt.Fprintf(w, string(app))
}

func GetAllGamesInSteam() {
	app := steam_app.GetAllGames()
	for _, game := range app {
		user := &schema.SteamGames{Appid: game.Appid, Name: game.Name}
		bytes, _ := json.Marshal(user)
		kafka.Producer("steamgames", bytes, []byte(strconv.Itoa(game.Appid)))
	}
}
