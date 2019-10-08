package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var slackSecret = os.Getenv("SLACK_SECRET")
var tokenAPI = os.Getenv("SLACK_API_TOKEN")
var domain = os.Getenv("DOMAIN")

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func doHug(user string, responseURL string) {
	userParts := strings.Split(user, "|")
	userID := userParts[0][2:]

	if !IsLetter(userID) || len(userID) != 9 {
		postResponseText(responseURL, "usage: /pepe hug @person")
		return
	}

	profile, _ := getProfileImage(userID)
	downloadImage(fmt.Sprintf("/app/tmp/%s.jpg", userID), profile.Profile.Image512)
	processImage(userID)
	imageURL := fmt.Sprintf("%s/static/%s.png", domain, userID)
	postResponseImage(responseURL, fmt.Sprintf("relax %s, everything will be alright", user), imageURL)
}

func webhookClientHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	parts := strings.Split(text, " ")

	responseURL := r.FormValue("response_url")
	if parts[0] == "hug" {
		if len(parts) == 1 {
			postResponseText(responseURL, "you need to tag someone")
			return
		}

		go doHug(parts[1], responseURL)
	} else {
		postResponseText(responseURL, "command unrecognized")
	}
}

func main() {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8000"
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc(fmt.Sprintf("/slack-webhook/%s", slackSecret), webhookClientHandler)

	connStr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Listening on %s", connStr)
	if err := http.ListenAndServe(connStr, nil); err != nil {
		log.Fatalf("Can't listen in port %s", port)
	}
}
