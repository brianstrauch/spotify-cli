package pkg

import (
	"fmt"
	"net/http"
)

const BaseAPIURL = "https://api.spotify.com/v1"

func Play(token string) error {
	url := BaseAPIURL + "/me/player/play"
	req, _ := http.NewRequest("PUT", url, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := http.Client{}
	_, err := client.Do(req)
	return err
}

func Pause(token string) error {
	url := BaseAPIURL + "/me/player/pause"
	req, _ := http.NewRequest("PUT", url, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := http.Client{}
	_, err := client.Do(req)
	return err
}
