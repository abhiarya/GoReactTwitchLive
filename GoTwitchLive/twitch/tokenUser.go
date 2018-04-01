package twitch

import (
	"time"
)

//TokenUser ...
type TokenUser struct {
	TokenValid TokenValid `json:"token"`
}

//TokenValid ...
type TokenValid struct {
	Valid         bool          `json:"valid"`
	Authorization Authorization `json:"authorization"`
	UserName      string        `json:"user_name"`
	UserID        string        `json:"user_id"`
	ClientID      string        `json:"client_id"`
	ExpiresIn     int           `json:"expires_in"`
}

//Authorization ...
type Authorization struct {
	Scopes    []string  `json:"scopes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
