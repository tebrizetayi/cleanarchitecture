package inmemory

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

func TestArticleRepository(t *testing.T) {
	Convey("Creating multiple  articles", t, func() {
		articlerepo := NewArticleInmemoryRepo()
		authorrepo := NewAuthorInmemoryRepo()

		author := model.Author{
			Name: "Jack London",
		}
		createdAuthor, _ := authorrepo.Create([]model.Author{author})
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
		Convey("When you create new articles", func() {
			created, err := articlerepo.Create(articles)
			So(err, ShouldBeNil)
			Convey("Then ID should be bigger than zero", func() {
				ids := []int{}
				for _, v := range created {
					So(v.ID, ShouldNotBeZeroValue)
					ids = append(ids, v.ID)
				}

				Convey("And it can get by id", func() {
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

				Convey("")
			})
		})
	})
}
