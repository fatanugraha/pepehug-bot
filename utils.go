package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func getProfileImage(userID string) (*ProfileAPIResult, error) {
	apiURL := fmt.Sprintf("https://slack.com/api/users.profile.get?token=%s&user=%s", tokenAPI, userID)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, _ := ioutil.ReadAll(resp.Body)
	result := &ProfileAPIResult{}
	if err := json.Unmarshal(res, &result); err != nil {
		panic(err)
	}

	return result, err

}

func postResponseImage(responseURL string, title string, imageURL string) error {
	payload := ResponsePayload{
		Attachments: &[]Attachment{
			Attachment{
				Title:    "Here's a hug for you",
				Pretext:  title,
				ImageURL: imageURL,
				ThumbURL: imageURL,
			},
		},
		ResponseType: "in_channel",
	}

	serialized, err := json.Marshal(payload)
	_, err = http.Post(responseURL, "application/json", bytes.NewBuffer(serialized))
	return err
}

func postResponseText(responseURL string, text string) error {
	payload := ResponsePayload{Text: text}
	serialized, err := json.Marshal(payload)
	_, err = http.Post(responseURL, "application/json", bytes.NewBuffer(serialized))
	return err
}

func downloadImage(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	written, err = io.Copy(out, resp.Body)
	fmt.Println("written", written)
	return err
}

func processImage(userID string) {
	dir, _ := os.Getwd()
	photoPath := fmt.Sprintf("/app/tmp/%s.jpg", userID)
	photoResizedPath := fmt.Sprintf("/app/tmp/%s_200.png", userID)
	backgroundPath := fmt.Sprintf("%s/assets/back.png", dir)
	handPath := fmt.Sprintf("%s/assets/top.png", dir)
	resultPath := fmt.Sprintf("%s/static/%s.png", dir, userID)
	fmt.Println(resultPath, handPath, backgroundPath, photoPath, photoResizedPath)
	exec.Command("convert", photoPath, "-resize", "160", photoResizedPath).Run()
	exec.Command("convert", "-background", "rgba(0,0,0,0)", "-rotate", "335", photoResizedPath, photoResizedPath).Run()
	exec.Command("composite", "-geometry", "+40+260", photoResizedPath, backgroundPath, resultPath).Run()
	exec.Command("composite", "-gravity", "center", handPath, resultPath, resultPath).Run()
}
