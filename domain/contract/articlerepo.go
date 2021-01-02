package contract

import (
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type ArticleRepository interface {
	GetAll() ([]model.Article, error)
	Delete(articleId []int) error
	Reset()
	GetByIds(ids []int) ([]model.Article, error)
	Create(articles []model.Article) ([]model.Article, error)
}
