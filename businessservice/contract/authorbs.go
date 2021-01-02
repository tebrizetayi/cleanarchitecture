package contract

import "github.com/tebrizetayi/cleanarchitecture/domain/model"

type AuthorBS interface {
	GetAll() ([]model.Author, error)
	Delete(ids []int) error
	Create(authors []model.Author) ([]model.Author, error)
	GetByIds(ids []int) (model.Author, error)
}
