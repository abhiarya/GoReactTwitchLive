package twitch

import (
	"time"
)

//Streams ...
type Streams struct {
	Total   int      `json:"_total"`
	Streams []Stream `json:"streams"`
}

//Stream ..
type Stream struct {
	ID          int       `json:"_id"`
	AverageFPS  int       `json:"average_fps"`
	Channel     Channel   `json:"channel"`
	CreatedAt   time.Time `json:"created_at"`
	Delay       int       `json:"delay"`
	Game        string    `json:"game"`
	IsPlaylist  bool      `json:"is_playlist"`
	Preview     Preview   `json:"preview"`
	VideoHeight int       `json:"video_height"`
	Viewers     int       `json:"viewers"`
}

//Channel ...
type Channel struct {
	BroadcasterLanguage   string    `json:"broadcaster_language"`
	CreatedAt             string    `json:"created_at"`
	DisplayName           string    `json:"display_name"`
	Followers             int       `json:"followers"`
	Game                  string    `json:"game"`
	Language              string    `json:"language"`
	Logo                  string    `json:"logo"`
	Mature                bool      `json:"mature"`
	Name                  string    `json:"name"`
	Partner               bool      `json:"partner"`
	ProfileBanner         string    `json:"profile_banner"`
	ProfileBannerBkgColor string    `json:"profile_banner_background_color"`
	Status                string    `json:"status"`
	UpdatedAt             time.Time `json:"updated_at"`
	URL                   string    `json:"url"`
	VideoBanner           string    `json:"video_banner"`
	Views                 int       `json:"views"`
}

//Preview ...
type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}
