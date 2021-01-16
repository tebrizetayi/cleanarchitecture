package inmemory

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

func TestArticleRepository(t *testing.T) {
	Convey("Setup", t, func() {
		articlerepo := NewArticleInmemoryRepo()
		Articles := []model.Article{
			{
				Name:   "Spanish",
				Author: []model.Author{{Name: "John Doe"}},
			},
			{
				Name:   "Jack London",
				Author: []model.Author{{Name: "William"}},
			},
		}
		created, err := articlerepo.Create(Articles)
		So(err, ShouldBeNil)

		ids := []int{}
		for _, v := range created {
			So(v.ID, ShouldNotBeZeroValue)
			ids = append(ids, v.ID)
		}

		//First article is taken for testing
		article := created[0]
		Convey("When you update article where the id in the database", func() {
			article.Name = article.Name + " " + article.Name
			_, err := articlerepo.Update([]model.Article{article})
			So(err, ShouldBeNil)
			Convey("Then updated article can be get by id", func() {
				updateArticles, err := articlerepo.GetByIds([]int{article.ID})
				So(err, ShouldBeNil)
				So(len(updateArticles), ShouldEqual, 1)
				So(updateArticles[0], ShouldResemble, article)
			})

		})

		Convey("When you delete article where the id is in the database", func() {
			err := articlerepo.Delete([]int{article.ID})
			So(err, ShouldBeNil)
			Convey("Then the deleted id should not be in the database", func() {
				deletedArticles, err := articlerepo.GetByIds([]int{article.ID})
				So(err, ShouldBeNil)
				So(len(deletedArticles), ShouldEqual, 0)
			})
		})

		Convey("When you add article", func() {
			newArticle := model.Article{
				Name: "New Article",
			}
			newArticles, err := articlerepo.Create([]model.Article{newArticle})
			So(err, ShouldBeNil)
			So(len(newArticles), ShouldEqual, 1)
			So(newArticles[0].ID, ShouldBeGreaterThan, 0)
			Convey("Then the added id should be in the database", func() {
				articles, err := articlerepo.GetByIds([]int{newArticles[0].ID})
				So(err, ShouldBeNil)
				So(len(articles), ShouldEqual, 1)
				So(newArticles, ShouldResemble, articles)
			})
		})
	})
}
