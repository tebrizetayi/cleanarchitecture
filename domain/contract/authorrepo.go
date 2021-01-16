package contract

import (
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorRepository interface {
	GetAll() ([]model.Author, error)
	Delete(ids []int) error
	GetByIds(ids []int) ([]model.Author, error)
	Create(authors []model.Author) ([]model.Author, error)
	Reset()
}
