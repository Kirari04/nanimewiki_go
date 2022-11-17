package models

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"encoding/json"
	"io"
	"net/http"
	"os"
)

var DB *gorm.DB

type offline_database struct {
	Repository string  `json:"repository"`
	LastUpdate string  `json:"lastUpdate"`
	Data       []Anime `json:"data"`
}

func ConnectDatabase(databaseFile string) {
	database, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Anime{})
	if err != nil {
		return
	}

	DB = database
}

func SeedDatabase() {
	fmt.Println("Downloading Seed")
	success := downloadFile(os.Getenv("database_seed_file"), os.Getenv("database_seed_data"))
	if !success {
		fmt.Println("Failed to download file from [database_seed_data] and seed the database")
		return
	}

	jsonFile, err := os.Open(os.Getenv("database_seed_file"))
	if err != nil {
		fmt.Println("Failed to open [database_seed_file]")
		fmt.Printf("err: %v\n", err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Failed to read all from [database_seed_file]")
		fmt.Printf("err: %v\n", err)
		return
	}

	// we initialize our off_db_val
	var off_db_val offline_database

	// we unmarshal our byteArray which contains our
	err = json.Unmarshal(byteValue, &off_db_val)
	if err != nil {
		fmt.Println("Failed to parse [database_seed_file]")
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println("Seed has been downloaded")
	fmt.Println("animes: " + strconv.Itoa(len(off_db_val.Data)))
	fmt.Println("lastUpdate: " + off_db_val.LastUpdate)
	fmt.Println("Start seeding database")

	seedStart := time.Now()
	var deleteCurrent, _ = strconv.ParseBool(os.Getenv("database_seed_overwrite"))
	if deleteCurrent == true {
		DB.Where("1 = 1").Delete(&Anime{})
	}

	for i := 0; i < len(off_db_val.Data); i++ {
		DB.Create(&off_db_val.Data[i])
	}
	seedElapsed := time.Since(seedStart)
	fmt.Printf("Time to Seed: %s\n", seedElapsed)
}

func downloadFile(filepath string, url string) bool {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("bad status: %s", resp.Status)
		return false
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	}

	return true
}
