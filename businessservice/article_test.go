package businessservice

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory"
)

var (
	articleRepo = inmemory.NewArticleInmemoryRepo()
	articleBS   = NewArticleBS(&articleRepo)
)

func TestArticle(t *testing.T) {
	Convey("Testing getting article", t, func() {
		Convey("WHEN an article is created in the system", func() {
			articles := []model.Article{
				{
					Name:   "Beginning to programming",
					Author: nil,
				},
				{
					Name:   "Beginning to spanish",
					Author: nil,
				}}
			created, _ := articleBS.Create(articles)
			ids := []int{}
			for i := 0; i < len(created); i++ {
				ids = append(ids, created[i].ID)
			}
			Convey("THEN created should be in the database", func() {
				found, err := articleBS.GetByIds(ids)
				So(err, ShouldBeNil)
				for i := 0; i < len(found); i++ {
					So(found[i].ID, ShouldEqual, created[i].ID)
				}
				Convey("THEN can be deleted from the database", func() {
					err := articleBS.Delete(ids)
					So(err, ShouldBeNil)

					found, err := articleBS.GetByIds(ids)
					So(err, ShouldBeNil)
					So(found, ShouldBeEmpty)
				})
			})
		})
	})
}
