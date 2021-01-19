package mysql

import (
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

func TestAuthorRepository(t *testing.T) {
	Convey("Setup", t, func() {
		conn := "root:secret@tcp(127.0.0.1:3306)/Academia"
		authorrepo, err := NewAuthorMysqlRepo(conn)
		So(err, ShouldBeNil)
		Authors := []model.Author{
			{
				Name: "John Doe",
			},
			{
				Name: "Jack London",
			},
		}
		created, err := authorrepo.Create(Authors)
		So(err, ShouldBeNil)

		ids := []uuid.UUID{}
		for _, v := range created {
			So(v.ID, ShouldNotBeZeroValue)
			ids = append(ids, v.ID)
		}

		//First author is taken for testing
		author := created[0]
		Convey("When you update author where the id in the database", func() {
			author.Name = author.Name + " " + author.Name
			_, err := authorrepo.Update([]model.Author{author})
			So(err, ShouldBeNil)
			Convey("Then updated author can be get by id", func() {
				updateAuthors, err := authorrepo.GetByIds([]uuid.UUID{author.ID})
				So(err, ShouldBeNil)
				So(len(updateAuthors), ShouldEqual, 1)
				So(updateAuthors[0], ShouldResemble, author)
			})

		})

		Convey("When you delete author where the id is in the database", func() {
			err := authorrepo.Delete([]uuid.UUID{author.ID})
			So(err, ShouldBeNil)
			Convey("Then the deleted id should not be in the database", func() {
				deletedAuthors, err := authorrepo.GetByIds([]uuid.UUID{author.ID})
				So(err, ShouldBeNil)
				So(len(deletedAuthors), ShouldEqual, 0)
			})
		})

		Convey("When you add author", func() {
			newAuthor := model.Author{
				Name: "New Author",
			}
			newAuthors, err := authorrepo.Create([]model.Author{newAuthor})
			So(err, ShouldBeNil)
			So(len(newAuthors), ShouldEqual, 1)
			So(newAuthors[0].ID, ShouldNotBeNil)
			Convey("Then the added id should be in the database", func() {
				authors, err := authorrepo.GetByIds([]uuid.UUID{newAuthors[0].ID})
				So(err, ShouldBeNil)
				So(len(authors), ShouldEqual, 1)
				So(newAuthors, ShouldResemble, authors)
			})
		})
	})
}
