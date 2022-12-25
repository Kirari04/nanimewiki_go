package setups

import (
	"ch/kirari/animeApi/models"
	"io/ioutil"
	"log"
	"strings"

	b64 "encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

func SeedZincSearch_AddEntry(anime models.Anime) bool {
	type ExpectedResponse struct {
		Message string `json:"message"`
	}

	url := os.Getenv("zinc_host") + "/api/animes/_doc/"
	method := "PUT"

	payload := strings.NewReader(`{
		"title": "` + anime.Title + `"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
		return false
	}
	authData := b64.StdEncoding.EncodeToString([]byte(os.Getenv("zinc_user") + ":" + os.Getenv("zinc_password")))
	req.Header.Add("Authorization", "Basic "+authData)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Println(string(body))

	var data ExpectedResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if data.Message != "ok" {
		log.Fatalf("Unknow response: %v", data.Message)
		log.Fatal(body)
		return false
	}

	return true
}
