package contract

import "github.com/tebrizetayi/cleanarchitecture/domain/model"

type AuthorBS interface {
	GetAll() (model.Author, error)
	Delete(AuthorIds []int) error
}
