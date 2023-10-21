package models

import "database/sql"

type ShareLink struct {
	ID       uint64
	Link     string
	Redirect string
	Visited  sql.NullInt64
}
