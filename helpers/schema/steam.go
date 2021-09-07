package schema

type SteamMapping struct {
	Response struct {
		GameCount int `json:"game_count"`
		Games     []struct {
			Appid                  int `json:"appid"`
			PlaytimeForever        int `json:"playtime_forever"`
			PlaytimeWindowsForever int `json:"playtime_windows_forever"`
			PlaytimeMacForever     int `json:"playtime_mac_forever"`
			PlaytimeLinuxForever   int `json:"playtime_linux_forever"`
		} `json:"games"`
	} `json:"response"`
}

type SteamReqBody struct {
	UserId string `json:"userId"`
}
