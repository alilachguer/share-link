package database

import (
	"github.com/alilachguer/share-link/models"
)

type Mem struct {
	shareLinks []models.ShareLink
}

func NewMemDB() *Mem {
	return &Mem{shareLinks: []models.ShareLink{}}
}

func (mem *Mem) All() ([]models.ShareLink, error) {
	return mem.shareLinks, nil
}

func (mem *Mem) Count() (int, error) {
	return len(mem.shareLinks), nil
}

func (mem *Mem) Create(sl models.ShareLink) ([]models.ShareLink, error) {
	mem.shareLinks = append(mem.shareLinks, sl)

	return mem.shareLinks, nil
}
