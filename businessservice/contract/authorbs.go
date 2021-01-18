package contract

import (
	"github.com/google/uuid"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorBS interface {
	GetAll() ([]model.Author, error)
	Delete(ids []uuid.UUID) error
	Create(authors []model.Author) ([]model.Author, error)
	GetByIds(ids []uuid.UUID) ([]model.Author, error)
	Update(authors []model.Author) ([]model.Author, error)
}
