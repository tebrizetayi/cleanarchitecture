package contract

import "github.com/tebrizetayi/cleanarchitecture/domain/model"

type ArticleBS interface {
	GetAll() ([]model.Article, error)
	Delete(ids []int) error
	Create(articles []model.Article) ([]model.Article, error)
	GetById(ids []int) (model.Article, error)
}
