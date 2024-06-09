package models

type ShareLink struct {
	ID       int
	Link     string
	Redirect string
	Visited  bool
}
