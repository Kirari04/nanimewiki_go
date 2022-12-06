package setups

import (
	"ch/kirari/animeApi/models"
	"io/ioutil"
	"log"
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
	Repository string         `json:"repository"`
	LastUpdate string         `json:"lastUpdate"`
	Data       []models.Anime `json:"data"`
}

func ConnectDatabase(databaseFile string) {
	database, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	// migrate Animes
	err = database.AutoMigrate(&models.Anime{})
	if err != nil {
		log.Panic("Failed to migrate Anime{}")
		return
	}
	// migrate Users
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Panic("Failed to migrate User{}")
		return
	}
	// migrate EmailVerificationKey
	err = database.AutoMigrate(&models.EmailVerificationCode{})
	if err != nil {
		log.Panic("Failed to migrate EmailVerificationKey{}")
		return
	}

	DB = database
}

func SeedDatabase() {
	log.Println("Downloading Seed")
	success := downloadFile(os.Getenv("database_seed_file"), os.Getenv("database_seed_data"))
	if !success {
		log.Println("Failed to download file from [database_seed_data] and seed the database")
		return
	}

	jsonFile, err := os.Open(os.Getenv("database_seed_file"))
	if err != nil {
		log.Println("Failed to open [database_seed_file]")
		log.Printf("err: %v\n", err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("Failed to read all from [database_seed_file]")
		log.Printf("err: %v\n", err)
		return
	}

	// we initialize our off_db_val
	var off_db_val offline_database

	// we unmarshal our byteArray which contains our
	err = json.Unmarshal(byteValue, &off_db_val)
	if err != nil {
		log.Println("Failed to parse [database_seed_file]")
		log.Printf("err: %v\n", err)
		return
	}
	log.Println("Seed has been downloaded")
	log.Println("animes: " + strconv.Itoa(len(off_db_val.Data)))
	log.Println("lastUpdate: " + off_db_val.LastUpdate)
	log.Println("Start seeding database")

	seedStart := time.Now()
	var deleteCurrent, _ = strconv.ParseBool(os.Getenv("database_seed_overwrite"))
	if deleteCurrent == true {
		DB.Where("1 = 1").Delete(&models.Anime{})
	}

	for i := 0; i < len(off_db_val.Data); i++ {
		DB.Create(&off_db_val.Data[i])
	}
	seedElapsed := time.Since(seedStart)
	log.Printf("Time to Seed: %s\n", seedElapsed)
}

func downloadFile(filepath string, url string) bool {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Printf("err: %v\n", err)
		return false
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("err: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		log.Printf("bad status: %s", resp.Status)
		return false
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Printf("err: %v\n", err)
		return false
	}

	return true
}
