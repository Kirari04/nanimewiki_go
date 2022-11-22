package controllers__test

import (
	"ch/kirari/animeApi/models"
	"ch/kirari/animeApi/setups"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var host string

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

type ExpectedResponse struct {
	Data    []models.Anime `json:"data"`
	Error   string         `json:"error"`
	Len     int            `json:"len"`
	Success int            `json:"success"`
}

// PREPARATION

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("m: %v\n", err)
		return
	}
	host = os.Getenv("test_host")
	log.Printf("Host: %v\n", host)

	setups.ConnectDatabase(os.Getenv("database"))

	code := m.Run()
	os.Exit(code)
}

// HELPERS

func AssertEqualString(t *testing.T, val1 string, val2 string) {
	if val1 != val2 {
		t.Fatalf("The data doesn't match as expected: %v vs: %v", val1, val2)
	}
}

func AssertEqualInt(t *testing.T, val1 int, val2 int) {
	if val1 != val2 {
		t.Fatalf("The data doesn't match as expected: %v vs: %v", val1, val2)
	}
}

func AssertEqualBool(t *testing.T, val1 bool, val2 bool) {
	if val1 != val2 {
		t.Fatalf("The data doesn't match as expected: %v vs: %v", val1, val2)
	}
}

func getReq(t *testing.T, route string, expectedStatusCode int) ExpectedResponse {
	res, err := http.Get(host + route)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != expectedStatusCode {
		t.Fatalf("Status Code %v - %v", res.StatusCode, res.Body)
	}

	defer res.Body.Close()
	byteValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var data ExpectedResponse
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		t.Fatal(err)
	}

	return data
}

func postReq(t *testing.T, route string, expectedStatusCode int, formData url.Values) ExpectedResponse {
	res, err := http.PostForm(host+route, formData)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != expectedStatusCode {
		t.Error(formData)
		t.Fatalf("Status Code %v - %v", res.StatusCode, res.Body)
	}

	defer res.Body.Close()
	byteValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var data ExpectedResponse
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		t.Fatal(err)
	}

	return data
}

func getTmpEmails(t *testing.T, route string, expectedStatusCode int, mapper *[]string) {
	res, err := http.Get(route)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != expectedStatusCode {
		t.Fatalf("Status Code %v - %v", res.StatusCode, res.Body)
	}

	defer res.Body.Close()
	byteValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &mapper)
	if err != nil {
		t.Fatal(err)
	}
}

func getTmpMessages(t *testing.T, route string, expectedStatusCode int, mapper *[]tmpEmailListResponse) {
	res, err := http.Get(route)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != expectedStatusCode {
		t.Fatalf("Status Code %v - %v", res.StatusCode, res.Body)
	}

	defer res.Body.Close()
	byteValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &mapper)
	if err != nil {
		t.Fatal(err)
	}
}
