package inmemory

import (
	"errors"

	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type ArticleInmemory struct {
	Articles map[int]model.Article
	sequence int
}

func NewArticleInmemory() ArticleInmemory {
	return ArticleInmemory{sequence: 0}
}

func (a *ArticleInmemory) Create(articles []model.Article) ([]model.Article, error) {
	if a.Articles == nil {
		a.Articles = make(map[int]model.Article)
	}

	result := []model.Article{}
	for _, v := range articles {
		a.sequence++
		v.ID = a.sequence
		a.Articles[v.ID] = v
		result = append(result, v)
	}
	return result, nil
}

func (a *ArticleInmemory) GetAll() ([]model.Article, error) {
	if a.Articles == nil {
		return nil, errors.New("No data")
	}

	result := []model.Article{}
	for _, v := range a.Articles {
		result = append(result, v)
	}
	return result, nil
}

func (a *ArticleInmemory) Delete(ids []int) error {
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

func (a *ArticleInmemory) GetByIds(ids []int) ([]model.Article, error) {
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

func (a *ArticleInmemory) Reset() {
	a.Articles = make(map[int]model.Article)
	a.sequence = 0
}
