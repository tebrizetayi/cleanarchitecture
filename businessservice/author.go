package businessservice

import (
	"github.com/tebrizetayi/cleanarchitecture/domain/contract"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorBS struct {
	AuthorRepo contract.AuthorRepository
}

func NewAuthorBS(AuthorRepo contract.AuthorRepository) AuthorBS {
	return AuthorBS{AuthorRepo: AuthorRepo}
}

func (a *AuthorBS) GetAll() (model.Author, error) {
	//Some Code
	a.AuthorRepo.GetAll(0, -1)
	return model.Author{}, nil
}

func (a *AuthorBS) Delete(AuthorIds []int) error {
	//Some Code
	a.AuthorRepo.Delete(AuthorIds)
	return nil
}
