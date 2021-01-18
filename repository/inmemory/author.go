package inmemory

import (
	"github.com/google/uuid"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorInmemoryRepo struct {
	Authors map[uuid.UUID]model.Author
}

func NewAuthorInmemoryRepo() AuthorInmemoryRepo {
	return AuthorInmemoryRepo{Authors: make(map[uuid.UUID]model.Author)}
}

func (a *AuthorInmemoryRepo) Create(Authors []model.Author) ([]model.Author, error) {

	result := []model.Author{}
	for _, v := range Authors {
		v.ID = uuid.New()
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

func (a *AuthorInmemoryRepo) Delete(ids []uuid.UUID) error {

	for _, v := range ids {
		if _, ok := a.Authors[v]; ok {
			delete(a.Authors, v)
		}
	}

	return nil
}

func (a *AuthorInmemoryRepo) GetByIds(ids []uuid.UUID) ([]model.Author, error) {

	result := []model.Author{}
	for _, v := range ids {
		if v, ok := a.Authors[v]; ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (a *AuthorInmemoryRepo) Reset() {
	a.Authors = make(map[uuid.UUID]model.Author)
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
