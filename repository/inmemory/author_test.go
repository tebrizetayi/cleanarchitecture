package inmemory

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

func TestAuthorRepository(t *testing.T) {
	Convey("Creating multiple  Authors", t, func() {
		authorrepo := NewAuthorInmemoryRepo()
		Authors := []model.Author{
			{
				Name: "John Doe",
			},
			{
				Name: "Jack London",
			},
		}
		Convey("When you create new Authors", func() {
			created, err := authorrepo.Create(Authors)
			So(err, ShouldBeNil)
			Convey("Then ID should be bigger than zero", func() {
				ids := []int{}
				for _, v := range created {
					So(v.ID, ShouldNotBeZeroValue)
					ids = append(ids, v.ID)
				}

				Convey("Then it can get by id", func() {
					found, err := authorrepo.GetByIds(ids)
					So(err, ShouldBeNil)
					for i, _ := range found {
						So(found[i].ID, ShouldEqual, created[i].ID)
					}
				})

				Convey("Then it can be deleted", func() {
					authorrepo.Delete(ids)
					found, _ := authorrepo.GetByIds(ids)
					So(found, ShouldBeEmpty)
				})
			})
		})
	})
}
