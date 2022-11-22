package controllers__test

import (
	"ch/kirari/animeApi/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var host string

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
	code := m.Run()
	os.Exit(code)
}

// TESTS

func TestAnimeListCorrect(t *testing.T) {
	res := getReq(t, "/api/v1/anime/list/0", 200)
	AssertEqualInt(t, res.Success, 1)
	AssertEqualBool(t, (res.Len > 0), true)
}

func TestAnimeListNoParameter(t *testing.T) {
	res := getReq(t, "/api/v1/anime/list", 200)
	AssertEqualInt(t, res.Success, 1)
	AssertEqualBool(t, (res.Len > 0), true)
}

func TestAnimeListNegativIndex(t *testing.T) {
	getReq(t, "/api/v1/anime/list/-1", 400)
}

func TestAnimeListIncorrectParameterValue(t *testing.T) {
	getReq(t, "/api/v1/anime/list/onlyIntegerHere", 400)
}

func TestAnimeListBigIntParameter(t *testing.T) {
	getReq(t, "/api/v1/anime/list/999999999999999999999999999999999999999999999", 400)
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
		t.Errorf("Status Code %v - %v", res.StatusCode, res.Body)
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
