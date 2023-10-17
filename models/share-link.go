package models

type ShareLink struct {
	ID       uint64
	Link     string
	Redirect string
	Visited  uint64
}
