package inmemory

import (
	"errors"

	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorInmemoryRepo struct {
	Authors  map[int]model.Author
	sequence int
}

func NewAuthorInmemoryRepo() AuthorInmemoryRepo {
	return AuthorInmemoryRepo{sequence: 0}
}

func (a *AuthorInmemoryRepo) Create(Authors []model.Author) ([]model.Author, error) {
	if a.Authors == nil {
		a.Authors = make(map[int]model.Author)
	}

	result := []model.Author{}
	for _, v := range Authors {
		a.sequence++
		v.ID = a.sequence
		a.Authors[v.ID] = v
		result = append(result, v)
	}
	return result, nil
}

func (a *AuthorInmemoryRepo) GetAll() ([]model.Author, error) {
	if a.Authors == nil {
		return nil, errors.New("No data")
	}

	result := []model.Author{}
	for _, v := range a.Authors {
		result = append(result, v)
	}
	return result, nil
}

func (a *AuthorInmemoryRepo) Delete(ids []int) error {
	if a.Authors == nil {
		return errors.New("No data")
	}

	for _, v := range ids {
		if _, ok := a.Authors[v]; ok {
			delete(a.Authors, v)
		}
	}

	return nil
}

func (a *AuthorInmemoryRepo) GetByIds(ids []int) ([]model.Author, error) {
	if a.Authors == nil {
		return nil, errors.New("No data")
	}

	result := []model.Author{}
	for _, v := range ids {
		if v, ok := a.Authors[v]; ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (a *AuthorInmemoryRepo) Reset() {
	a.Authors = make(map[int]model.Author)
	a.sequence = 0
}
