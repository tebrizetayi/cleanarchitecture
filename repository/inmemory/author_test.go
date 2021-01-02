package inmemory

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

func TestAuthorRepository(t *testing.T) {
	authorrepo := NewAuthorInmemoryRepo()
	Convey("Creating a Author", t, func() {
		Convey("When you create a new Author", func() {
			author := model.Author{
				Name: "John Doe",
			}
			createdAuthors, _ := authorrepo.Create([]model.Author{author})
			Convey("Then ID should be bigger than zero", func() {
				created := createdAuthors[0]
				So(created.ID, ShouldNotBeZeroValue)

				Convey("Then it can get by id", func() {
					found, _ := authorrepo.GetByIds([]int{created.ID})
					for i, v := range found {
						So(v.ID, ShouldEqual, createdAuthors[i].ID)
					}
				})

				Convey("Then it can be deleted", func() {
					authorrepo.Delete([]int{created.ID})
					found, err := authorrepo.GetByIds([]int{created.ID})
					So(err, ShouldBeEmpty)
					So(found, ShouldBeEmpty)

				})
			})

		})
		Reset(func() {
			authorrepo.Reset()
		})
	})

	Convey("Creating multiple  Authors", t, func() {
		Convey("When you create new Authors", func() {
			Authors := []model.Author{
				{
					Name: "John Doe",
				},
				{
					Name: "Jack London",
				},
			}
			created, _ := authorrepo.Create(Authors)
			Convey("Then ID should be bigger than zero", func() {

				ids := []int{}
				for _, v := range created {
					ids = append(ids, v.ID)
					So(v.ID, ShouldNotBeZeroValue)
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
		Reset(func() {
			authorrepo.Reset()
		})
	})
}
