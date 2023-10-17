package main

import (
	"fmt"
	"os"

	"github.com/alilachguer/share-link/database"
	"github.com/alilachguer/share-link/storage"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// mem := database.NewMemDB()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	var (
		dbUser = os.Getenv("DB_USER_NAME")
		dbPass = os.Getenv("DB_PASSWORD")
		dbName = os.Getenv("DB_DATABASE_NAME")
	)

	conn, err := database.NewDB("mysql", dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	defer conn.Conn.Close()

	db := storage.NewStorageRepo(conn)

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
