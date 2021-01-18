package businessservice

import (
	"github.com/google/uuid"
	"github.com/tebrizetayi/cleanarchitecture/domain/contract"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorBS struct {
	AuthorRepo contract.AuthorRepository
}

func NewAuthorBS(AuthorRepo contract.AuthorRepository) AuthorBS {
	return AuthorBS{
		AuthorRepo: AuthorRepo,
	}
}

func (a *AuthorBS) GetAll() ([]model.Author, error) {
	authors, err := a.AuthorRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (a *AuthorBS) Delete(ids []uuid.UUID) error {
	err := a.AuthorRepo.Delete(ids)
	return err
}

func (a *AuthorBS) Create(authors []model.Author) ([]model.Author, error) {
	authors, err := a.AuthorRepo.Create(authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (a *AuthorBS) GetByIds(ids []uuid.UUID) ([]model.Author, error) {
	authors, err := a.AuthorRepo.GetByIds(ids)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (a *AuthorBS) Update(authors []model.Author) ([]model.Author, error) {
	if authors == nil || len(authors) == 0 {
		return []model.Author{}, nil
	}

	return a.AuthorRepo.Update(authors)

}
