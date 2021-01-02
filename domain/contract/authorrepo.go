package contract

import (
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorRepository interface {
	GetAll(skip int, limit int) ([]model.Author, error)
	Delete(authorId []int) error
}
