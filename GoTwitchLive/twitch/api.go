package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func init() {
	fmt.Println("Initializing Twitch API...")
}

//GetLiveStreams gets all current livestreams from the Twitch API
func GetLiveStreams() []Response {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   time.Second * 60,
		Transport: tr,
	}

	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/streams/?stream_type=live&limit=100", nil)
	req.Header.Add("Client-ID", "<Put your client ID here>")
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%sError", err)
	}

	var streams Streams
	err = json.Unmarshal(body, &streams)

	responses := make([]Response, len(streams.Streams))
	for i := 0; i < len(streams.Streams); i++ {
		responses[i].ID = streams.Streams[i].ID
		responses[i].UserName = streams.Streams[i].Channel.DisplayName
		responses[i].PreviewImage = streams.Streams[i].Preview.Medium
		responses[i].LiveURL = streams.Streams[i].Channel.URL
	}

	return responses
}
