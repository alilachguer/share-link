package database

import (
	"database/sql"

	"github.com/alilachguer/share-link/models"
)

type DB struct {
	Conn *sql.DB
}

func NewDB(driver string, connectionString string) (*DB, error) {
	dbConnection, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: dbConnection}, nil
}

type User struct {
	id        uint64
	email     string
	firstName string
	lastName  string
	pass      string
}

func (mydb *DB) All() ([]models.ShareLink, error) {
	rows, err := mydb.Conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	links := []models.ShareLink{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.email, &user.firstName, &user.lastName, &user.pass)
		if err != nil {
			return nil, err
		}

		links = append(links, models.ShareLink{ID: user.id, Link: user.firstName, Redirect: user.email, Visited: 1})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

func (mydb *DB) Count() (int, error) {
	row := mydb.Conn.QueryRow("SELECT COUNT(*) FROM users")

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (mydb *DB) Create(sl models.ShareLink) ([]models.ShareLink, error) {
	return []models.ShareLink{}, nil
}
