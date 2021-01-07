package inmemory

import (
	"errors"
	"sync/atomic"

	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type ArticleInmemoryRepo struct {
	Articles map[int]model.Article
	sequence int
}

func NewArticleInmemoryRepo() ArticleInmemoryRepo {
	return ArticleInmemoryRepo{sequence: 0}
}

func (a *ArticleInmemoryRepo) Create(articles []model.Article) ([]model.Article, error) {
	if a.Articles == nil {
		a.Articles = make(map[int]model.Article)
	}

	result := []model.Article{}
	for _, v := range articles {
		atomic.AddInt32(a.sequence, 1)
		v.ID = a.sequence
		a.Articles[v.ID] = v
		result = append(result, v)
	}
	return result, nil
}

func (a *ArticleInmemoryRepo) GetAll() ([]model.Article, error) {
	if a.Articles == nil {
		return nil, errors.New("No data")
	}

	result := []model.Article{}
	for _, v := range a.Articles {
		result = append(result, v)
	}
	return result, nil
}

func (a *ArticleInmemoryRepo) Delete(ids []int) error {
	if a.Articles == nil {
		return errors.New("No data")
	}

	for _, v := range ids {
		if _, ok := a.Articles[v]; ok {
			delete(a.Articles, v)
		}
	}

	return nil
}

func (a *ArticleInmemoryRepo) GetByIds(ids []int) ([]model.Article, error) {
	if a.Articles == nil {
		return nil, errors.New("No data")
	}

	result := []model.Article{}
	for _, v := range ids {
		if v, ok := a.Articles[v]; ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (a *ArticleInmemoryRepo) Reset() {
	a.Articles = make(map[int]model.Article)
	a.sequence = 0
}
