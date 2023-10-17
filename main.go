package main

import (
	"fmt"

	"github.com/alilachguer/share-link/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// mem := database.NewMemDB()
	conn, err := database.NewDB("mysql", "root:root@/user_management")
	if err != nil {
		panic(err)
	}
	defer conn.Conn.Close()

	db := database.NewStorage(conn)

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
