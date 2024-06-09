package storage

import (
	"github.com/alilachguer/share-link/internal/models"
)

type storage interface {
	All() ([]models.ShareLink, error)
	Count() (int, error)
	Create(models.ShareLink) ([]models.ShareLink, error)

	GetRedirect(string) (string, error)
}

type StorageRepo struct {
	db storage
}

func NewStorageRepo(storage storage) *StorageRepo {
	return &StorageRepo{db: storage}
}

func (s *StorageRepo) GetAll() ([]models.ShareLink, error) {
	return s.db.All()
}

func (s *StorageRepo) GetRedirectByLink(link string) (string, error) {
	return s.db.GetRedirect(link)
}

func (s *StorageRepo) GetCount() (int, error) {
	return s.db.Count()
}

func (s *StorageRepo) CreateNew(sl models.ShareLink) ([]models.ShareLink, error) {
	return s.db.Create(sl)
}
