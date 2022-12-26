package setups

import (
	"ch/kirari/animeApi/models"
	"io"
	"log"
	"strings"

	b64 "encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

func SeedSearch() {
	var limit int64 = 200
	var index int64 = 0
	var itemLen int64
	DB.Model(&models.Anime{}).Count(&itemLen)

	log.Printf("Found %v animes to seed\n", itemLen)
	log.Println("Start seeding database")
	for i := int64(0); i < itemLen; i = i + limit {
		var animes []models.Anime
		DB.Limit(int(limit)).Offset(int(index * limit)).Find(&animes)
		ZincSearch_AddEntrys(animes)
		index++
	}
	log.Println("All animes had been added to the search")
}

func ZincSearch_AddEntrys(animes []models.Anime) bool {
	type ExpectedRequest struct {
		Index   string         `json:"index"`
		Records []models.Anime `json:"records"`
	}
	type ExpectedResponse struct {
		Message string `json:"message"`
	}

	url := os.Getenv("zinc_host") + "/api/_bulkv2"
	method := "POST"
	payload_string, err := json.Marshal(&ExpectedRequest{
		Index:   "animes",
		Records: animes,
	})
	if err != nil {
		log.Fatal(err)
		return false
	}
	payload := strings.NewReader(string(payload_string))

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

	body, err := io.ReadAll(res.Body)
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

	if data.Message != "v2 data inserted" {
		log.Fatal(string(body))
		return false
	}

	return true
}
