package businessservice

import (
	"github.com/tebrizetayi/cleanarchitecture/domain/model"

	"github.com/tebrizetayi/cleanarchitecture/domain/contract"
)

type ArticleBS struct {
	ArticleRepo contract.ArticleRepository
	AuthorRepo  contract.AuthorRepository
}

func NewArticleBS(articleRepo contract.ArticleRepository,
	authorRepo contract.AuthorRepository) ArticleBS {
	return ArticleBS{
		ArticleRepo: articleRepo,
		AuthorRepo:  authorRepo,
	}
}

func (a *ArticleBS) GetAll() (model.Article, error) {
	a.ArticleRepo.GetAll()
	return model.Article{}, nil
}

func (a *ArticleBS) Delete(articleIds []int) error {
	a.ArticleRepo.Delete(articleIds)
	return nil
}
