package inmemory

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

func TestArticleRepository(t *testing.T) {
	articlerepo := NewArticleInmemoryRepo()
	authorrepo := NewAuthorInmemoryRepo()

	author := model.Author{
		Name: "Jack London",
	}
	createdAuthor, _ := authorrepo.Create([]model.Author{author})

	Convey("Creating a article", t, func() {
		Convey("When you create a new article", func() {

			article := model.Article{
				Name:   "EFT from A to Z",
				Author: createdAuthor,
			}
			createdArticles, _ := articlerepo.Create([]model.Article{article})
			Convey("Then ID should be bigger than zero", func() {
				created := createdArticles[0]
				So(created.ID, ShouldNotBeZeroValue)

				Convey("Then it can get by id", func() {
					found, _ := articlerepo.GetByIds([]int{created.ID})
					for i, v := range found {
						So(v.ID, ShouldEqual, createdArticles[i].ID)
					}
				})

				Convey("Then it can be deleted", func() {
					articlerepo.Delete([]int{created.ID})
					found, err := articlerepo.GetByIds([]int{created.ID})
					So(err, ShouldBeEmpty)
					So(found, ShouldBeEmpty)

				})
			})

		})
		Reset(func() {
			articlerepo.Reset()
		})
	})

	Convey("Creating multiple  articles", t, func() {
		Convey("When you create new articles", func() {
			articles := []model.Article{
				{
					Name:   "EFT from A to Z",
					Author: createdAuthor,
				},
				{
					Name:   "Spanish with Lena",
					Author: createdAuthor,
				},
			}
			created, _ := articlerepo.Create(articles)
			Convey("Then ID should be bigger than zero", func() {

				ids := []int{}
				for _, v := range created {
					ids = append(ids, v.ID)
					So(v.ID, ShouldNotBeZeroValue)
				}

				Convey("Then it can get by id", func() {
					found, err := articlerepo.GetByIds(ids)
					So(err, ShouldBeNil)
					for i, _ := range found {
						So(found[i].ID, ShouldEqual, created[i].ID)
					}
				})

				Convey("Then it can be deleted", func() {
					articlerepo.Delete(ids)
					found, _ := articlerepo.GetByIds(ids)
					So(found, ShouldBeEmpty)
				})
			})

		})
		Reset(func() {
			articlerepo.Reset()
		})
	})
}
