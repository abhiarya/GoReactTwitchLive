package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var (
	twitchOauthConfig = oauth2.Config{
		ClientID:     "xxxxxxx",
		ClientSecret: "xxxxxxxx",
		Scopes:       []string{"user:edit"},
		RedirectURL:  "http://localhost:8080/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.twitch.tv/kraken/oauth2/authorize",
			TokenURL: "https://api.twitch.tv/kraken/oauth2/token",
		},
	}
)

//StreamFunc ...
func StreamFunc(w http.ResponseWriter, r *http.Request) {
	response := GetLiveStreams()
	json.NewEncoder(w).Encode(response)
}

//AuthorizeFunc ...
func AuthorizeFunc(w http.ResponseWriter, r *http.Request) {
	url := twitchOauthConfig.AuthCodeURL("twitch-live", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

//CallbackFunc ...
func CallbackFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.Form.Get("state")

	if state != "twitch-live" {
		http.Error(w, "invalid oauth state", http.StatusBadRequest)
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", "twitch-live", state)
		return
	}

	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := twitchOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Code exchange failed with '%s'\n", err)
		return
	}

	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken", nil)
	if err != nil {
		return
	}
	req.Header.Add("Client-ID", "xxxxx")
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Add("Authorization", "OAuth "+token.AccessToken)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   time.Second * 60,
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error %d %s", resp.StatusCode, body)
		return
	}

	var user TokenUser
	err = json.Unmarshal(body, &user)

	userName := User{user.TokenValid.UserName}
	json.NewEncoder(w).Encode(userName)
}

//User ...
type User struct {
	Name string
}
