package contract

import (
	"github.com/google/uuid"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorRepository interface {
	GetAll() ([]model.Author, error)
	Delete(ids []uuid.UUID) error
	GetByIds(ids []uuid.UUID) ([]model.Author, error)
	Create(authors []model.Author) ([]model.Author, error)
	Reset()
	Update(authors []model.Author) ([]model.Author, error)
}
