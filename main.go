package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var c = cache.New(1*time.Minute, 5*time.Minute)

type user struct {
	Stream *Stream `json:"stream,omitempty"`
}

type Stream struct {
	Title string `json:"title"`
}

func main() {
	// get the port from the args used to call the program
	// if no port is provided, exit
	port := os.Args[1]
	http.HandleFunc("/", checkOnline)
	log.Println("Starting server on port " + port)
	http.ListenAndServe(":"+port, nil)
}

func checkOnline(w http.ResponseWriter, r *http.Request) {
	log.Println("checkOnline")
	channel := r.URL.Query().Get("channel")

	if x, found := c.Get(channel); found {
		fmt.Fprintf(w, "%v", x)
		return
	}

	requestUrl := "https://api.ivr.fi/v2/twitch/user?login=" + channel
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Println("Error getting channel data:", err)
		fmt.Fprintf(w, "Error getting channel data")
	}
	defer resp.Body.Close()

	var users []user
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&users)
	if err != nil {
		log.Println("Error decoding channel data:", err)
		fmt.Fprintf(w, "Error decoding channel data")
	}

	log.Println(users)

	if users[0].Stream != nil {
		c.Set(channel, true, cache.DefaultExpiration)
		fmt.Fprintf(w, "true")
		return
	}
	c.Set(channel, false, cache.DefaultExpiration)
	fmt.Fprintf(w, "false")
}
