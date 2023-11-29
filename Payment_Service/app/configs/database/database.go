package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var clientInstance *gorm.DB
var clientInstanceError error
var gormOnce sync.Once

func GetGormClient() (*gorm.DB, error) {
	gormOnce.Do(func() {
		dbHost := viper.Get("POSTGRES_HOST")
		dbUser := viper.Get("POSTGRES_USER")
		dbPassword := viper.Get("POSTGRES_PASSWORD")
		dbName := viper.Get("POSTGRES_DB")
		dbPort := viper.Get("POSTGRES_PORT")

		if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
			log.Fatal("Incomplete DB credentials on .env")
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

		client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
