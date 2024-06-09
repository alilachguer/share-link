package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alilachguer/share-link/internal/database"
	"github.com/alilachguer/share-link/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("could not load environment .env :", err)
	}

	var config EnvConfig
	loadEnvVariables(&config)

	fmt.Println("EnvConfig", config)
	fmt.Println("database url", config.dbUrl)

	conn, err := database.NewDB("sqlite3", config.dbUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Conn.Close()

	db := storage.NewStorageRepo(conn)
	// db := storage.NewStorageRepo(database.NewMemDB())

	allLinks, err := db.GetAll()
	if err != nil {
		panic(err)
	}
	for _, link := range allLinks {
		fmt.Println(link)
	}

	count, err := db.GetCount()
	if err != nil {
		panic(err)
	}
	fmt.Println("count", count)

	redirect, err := db.GetRedirectByLink("share.link/1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("redirect", redirect)
}

type EnvConfig struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbUrl      string
}

func loadEnvVariables(conf *EnvConfig) {
	conf.dbUrl = os.Getenv("DB_URL")
	conf.dbUser = os.Getenv("DB_USER_NAME")
	conf.dbPassword = os.Getenv("DB_PASSWORD")
	conf.dbName = os.Getenv("DB_DATABASE_NAME")
}
