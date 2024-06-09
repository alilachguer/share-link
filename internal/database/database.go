package database

import (
	"database/sql"

	"github.com/alilachguer/share-link/internal/models"
	_ "github.com/mattn/go-sqlite3"
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

func (mydb *DB) All() ([]models.ShareLink, error) {
	rows, err := mydb.Conn.Query("SELECT rowid, * FROM sharelinks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := []models.ShareLink{}
	for rows.Next() {
		var sl models.ShareLink
		var v sql.NullInt16
		err := rows.Scan(&sl.ID, &sl.Link, &sl.Redirect, &v)
		if err != nil {
			return nil, err
		}

		sl.Visited = isVisited(v)

		links = append(links, sl)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

func (mydb *DB) GetRedirect(link string) (string, error) {
	row := mydb.Conn.QueryRow("SELECT redirect FROM sharelinks WHERE link = ?", link)

	var redirect string
	err := row.Scan(&redirect)
	if err != nil {
		return "", err
	}

	return redirect, nil
}

func (mydb *DB) Count() (int, error) {
	row := mydb.Conn.QueryRow("SELECT COUNT(*) FROM sharelinks")

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

func isVisited(i sql.NullInt16) bool {
	return i.Valid && i.Int16 != 0
}
