package inmemory

import (
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory/db"
)

type AuthorInmemoryRepo struct {
	*db.Inmemorydb
}

func NewAuthorInmemoryRepo() AuthorInmemoryRepo {
	return AuthorInmemoryRepo{
		Inmemorydb: db.GetInmemorydb(),
	}
}

func (a *AuthorInmemoryRepo) Create(Authors []model.Author) ([]model.Author, error) {
	result := []model.Author{}
	for _, v := range Authors {
		a.Inmemorydb.Authors.sequence++
		v.ID = a.Inmemorydb.Authors.sequence
		a.Inmemorydb.Authors[v.ID] = v
		result = append(result, v)
	}
	return result, nil
}

func (a *AuthorInmemoryRepo) GetAll() ([]model.Author, error) {
	result := []model.Author{}
	for _, v := range a.Inmemorydb.Authors {
		result = append(result, v)
	}
	return result, nil
}

func (a *AuthorInmemoryRepo) Delete(ids []int) error {
	for _, v := range ids {
		if _, ok := a.Inmemorydb.Authors[v]; ok {
			delete(a.Inmemorydb.Authors, v)
		}
	}

	return nil
}

func (a *AuthorInmemoryRepo) GetByIds(ids []int) ([]model.Author, error) {

	result := []model.Author{}
	for _, v := range ids {
		if v, ok := a.Inmemorydb.Authors[v]; ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (a *AuthorInmemoryRepo) Reset() {
	a.Inmemorydb.Authors = make(map[int]model.Author)
	a.Inmemorydb.sequence = 0
}
