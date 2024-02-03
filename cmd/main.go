package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alilachguer/share-link/internal/database"
	"github.com/alilachguer/share-link/internal/storage"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("could not load environment .env :", err)
	}

	var config EnvConfig
	loadEnvVariables(&config)

	conn, err := database.NewDB("mysql", config.dbUser+":"+config.dbPassword+"@/"+config.dbName)
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

	count, err := db.GetCount()
	if err != nil {
		panic(err)
	}

	for _, link := range allLinks {
		fmt.Println(link)
	}

	fmt.Println(count)
}

type EnvConfig struct {
	dbUser     string
	dbPassword string
	dbName     string
}

func loadEnvVariables(conf *EnvConfig) {
	conf.dbUser = os.Getenv("DB_USER_NAME")
	conf.dbPassword = os.Getenv("DB_PASSWORD")
	conf.dbName = os.Getenv("DB_DATABASE_NAME")
}
