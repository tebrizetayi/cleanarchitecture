package businessservice

import (
	"github.com/tebrizetayi/cleanarchitecture/domain/contract"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type ArticleBS struct {
	articleRepo contract.ArticleRepository
}

func NewArticleBS(articleRepo contract.ArticleRepository) ArticleBS {
	return ArticleBS{
		articleRepo: articleRepo,
	}
}

func (a *ArticleBS) GetAll() ([]model.Article, error) {
	articles, err := a.articleRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleBS) Delete(articleIds []int) error {
	err := a.articleRepo.Delete(articleIds)
	return err
}

func (a *ArticleBS) Create(articles []model.Article) ([]model.Article, error) {
	articles, err := a.articleRepo.Create(articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleBS) GetByIds(ids []int) ([]model.Article, error) {
	articles, err := a.articleRepo.GetByIds(ids)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
