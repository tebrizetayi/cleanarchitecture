package inmemory

import (
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorInmemoryRepo struct {
	Authors  map[int]model.Author
	sequence int
}

func NewAuthorInmemoryRepo() AuthorInmemoryRepo {
	return AuthorInmemoryRepo{sequence: 0, Authors: make(map[int]model.Author)}
}

func (a *AuthorInmemoryRepo) Create(Authors []model.Author) ([]model.Author, error) {

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

	result := []model.Author{}
	for _, v := range a.Authors {
		result = append(result, v)
	}
	return result, nil
}

func (a *AuthorInmemoryRepo) Delete(ids []int) error {

	for _, v := range ids {
		if _, ok := a.Authors[v]; ok {
			delete(a.Authors, v)
		}
	}

	return nil
}

func (a *AuthorInmemoryRepo) GetByIds(ids []int) ([]model.Author, error) {

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

func (a *AuthorInmemoryRepo) Update(authors []model.Author) ([]model.Author, error) {
	if authors == nil || len(authors) == 0 {
		return []model.Author{}, nil
	}

	result := []model.Author{}
	for _, author := range authors {
		if _, ok := a.Authors[author.ID]; ok {
			a.Authors[author.ID] = author
			result = append(result, a.Authors[author.ID])
		} else {
			result = append(result, model.Author{ID: author.ID})
		}
	}
	return result, nil
}
