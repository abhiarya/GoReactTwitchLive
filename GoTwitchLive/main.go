package main

import (
	"GoTwitchLive/twitch"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Booting the server...")

	http.HandleFunc("/authorize", twitch.AuthorizeFunc)
	http.HandleFunc("/oauth2", twitch.CallbackFunc)
	http.HandleFunc("/go-twitch-live", twitch.StreamFunc)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
