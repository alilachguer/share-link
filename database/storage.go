package database

import (
	"github.com/alilachguer/share-link/models"
)

type storage interface {
	All() ([]models.ShareLink, error)
	Count() (int, error)
	Create(models.ShareLink) ([]models.ShareLink, error)
}

type Storage struct {
	db storage
}

func NewStorage(storage storage) *Storage {
	return &Storage{db: storage}
}

func (s *Storage) GetAll() ([]models.ShareLink, error) {
	return s.db.All()
}

func (s *Storage) GetCount() (int, error) {
	return s.db.Count()
}

func (s *Storage) CreateNew(sl models.ShareLink) ([]models.ShareLink, error) {
	return s.db.Create(sl)
}
