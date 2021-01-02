package businessservice

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory"
)

var (
	authorRepo = inmemory.NewAuthorInmemoryRepo()
	authorBS   = NewAuthorBS(&authorRepo)
)

func TestAuthor(t *testing.T) {
	Convey("Testing getting Author", t, func() {
		Convey("WHEN an Author is created in the system", func() {
			authors := []model.Author{
				{
					Name: "John Doe",
				},
				{
					Name: "Jack London",
				}}
			created, _ := authorBS.Create(authors)
			ids := []int{}
			for i := 0; i < len(created); i++ {
				ids = append(ids, created[i].ID)
			}
			Convey("THEN created should be in the database", func() {
				found, err := authorBS.GetByIds(ids)
				So(err, ShouldBeNil)
				for i := 0; i < len(found); i++ {
					So(found[i].ID, ShouldEqual, created[i].ID)
				}
				Convey("THEN can be deleted from the database", func() {
					err := authorBS.Delete(ids)
					So(err, ShouldBeNil)

					found, err := authorBS.GetByIds(ids)
					So(err, ShouldBeNil)
					So(found, ShouldBeEmpty)
				})
			})
		})
	})
}
