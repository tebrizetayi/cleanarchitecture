package repository

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/domain/contract"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory"
	"github.com/tebrizetayi/cleanarchitecture/repository/mysql"
	"github.com/tebrizetayi/cleanarchitecture/testutils"
)

func TestMysqlAuthorRepo(t *testing.T) {
	db, err := sql.Open("mysql", testutils.MysqlDbConnString())
	if err != nil {
		t.Error(err)
	}
	authorRepo := mysql.NewAuthorMysqlRepo(db)
	AuthorRepositoryTestStructure(t, &authorRepo)
}

func TestInmemoryAuthorRepo(t *testing.T) {
	authorRepo := inmemory.NewAuthorInmemoryRepo()
	AuthorRepositoryTestStructure(t, &authorRepo)
}

func AuthorRepositoryTestStructure(t *testing.T, authorRepo contract.AuthorRepository) {
	Convey("Setup", t, func() {
		Authors := []model.Author{
			{
				Name: "John Doe",
			},
			{
				Name: "Jack London",
			},
		}
		created, err := authorRepo.Create(Authors)
		So(err, ShouldBeNil)

		Convey("When you add author", func() {
			newAuthor := model.Author{
				Name: "New Author",
			}
			newAuthors, err := authorRepo.Create([]model.Author{newAuthor})
			So(err, ShouldBeNil)
			So(len(newAuthors), ShouldEqual, 1)
			So(newAuthors[0].ID, ShouldNotBeNil)
			Convey("Then the added id should be in the database", func() {
				authors, err := authorRepo.GetByIds([]uuid.UUID{newAuthors[0].ID})
				So(err, ShouldBeNil)
				So(len(authors), ShouldEqual, 1)
				So(newAuthors, ShouldResemble, authors)
			})
		})

		Convey("When you update author where the id in the database", func() {
			author := created[0]
			author.Name = author.Name + " " + author.Name
			_, err := authorRepo.Update([]model.Author{author})
			So(err, ShouldBeNil)
			Convey("Then updated author can be get by id", func() {
				updateAuthors, err := authorRepo.GetByIds([]uuid.UUID{author.ID})
				So(err, ShouldBeNil)
				So(len(updateAuthors), ShouldEqual, 1)
				So(updateAuthors[0], ShouldResemble, author)
			})

		})

		Convey("When you delete author where the id is in the database", func() {
			author := created[0]
			err := authorRepo.Delete([]uuid.UUID{author.ID})
			So(err, ShouldBeNil)
			Convey("Then the deleted id should not be in the database", func() {
				deletedAuthors, err := authorRepo.GetByIds([]uuid.UUID{author.ID})
				So(err, ShouldBeNil)
				So(len(deletedAuthors), ShouldEqual, 0)
			})
		})

	})
}
