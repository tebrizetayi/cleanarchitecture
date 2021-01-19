package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
	"github.com/tebrizetayi/cleanarchitecture/repository/mysql"
)

func TestAuthorHandler(t *testing.T) {
	Convey("Setup", t, func() {
		//authorRepo := inmemory.NewAuthorInmemoryRepo()
		authorRepo, err := mysql.NewAuthorMysqlRepo("root:secret@tcp(127.0.0.1:3306)/Academia")
		So(err, ShouldBeNil)
		authors := []model.Author{
			{
				Name: "John Doe",
			},
			{
				Name: "Lev Tolstoy",
			},
		}
		createdAuthors, err := authorRepo.Create(authors)
		So(err, ShouldBeNil)

		authorBS := businessservice.NewAuthorBS(&authorRepo)
		authorHandler := NewAuthorHandler(&authorBS)
		handler := InitRoutes(&authorHandler)

		//First author is taken for testing
		author := createdAuthors[0]
		Convey("When you update author where the id in the database", func() {
			author.Name = author.Name + " " + author.Name
			authorJSON := fmt.Sprintf(`{"authors":[{"name":"%s","id":"%s"}]}`, author.Name, author.ID)
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/author/%s", author.ID), bytes.NewBuffer([]byte(authorJSON)))
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then updated author can be get by id", func() {
				So(rr.Code, ShouldEqual, http.StatusOK)
				var aRes AuthorResponse
				err := json.NewDecoder(rr.Body).Decode(&aRes)
				So(err, ShouldBeNil)
				So(len(aRes.Authors), ShouldBeGreaterThan, 0)
				So(author, ShouldResemble, aRes.Authors[0])
			})
		})
		Convey("When you update author where the id is not in the database", func() {
			author := model.Author{
				Name: "John Doe",
				ID:   uuid.New(),
			}
			authorJSON := fmt.Sprintf(`{"authors":[{"name":"%s","id":"%s"}]}`, author.Name, author.ID)
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/author/%s", author.ID), bytes.NewBuffer([]byte(authorJSON)))
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then updated author can be get by id", func() {
				So(rr.Code, ShouldEqual, http.StatusOK)
				var aRes AuthorResponse
				err := json.NewDecoder(rr.Body).Decode(&aRes)
				So(err, ShouldBeNil)
				So(author.Name, ShouldEqual, aRes.Authors[0].Name)
				//Because id is autogenereted property
				So(author.ID, ShouldNotEqual, aRes.Authors[0].ID)
			})
		})

		Convey("When you update author with wrong json", func() {

			authorJSON := `{"authors":[{"name"}]}`
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/author/%s", author.ID), bytes.NewBuffer([]byte(authorJSON)))
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then updated author can be get by id", func() {
				So(rr.Code, ShouldEqual, http.StatusBadRequest)

			})
		})
		Convey("When you delete author where the id is the database", func() {
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/author/%s", author.ID), nil)
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then the response should be no content", func() {
				So(rr.Code, ShouldEqual, http.StatusNoContent)
			})
		})

		Convey("When you create a new author", func() {
			authorJSON := fmt.Sprintf(`{"authors":[{"name":"%s"}]}`, author.Name)
			req, err := http.NewRequest(http.MethodPost, "/api/author/", bytes.NewBuffer([]byte(authorJSON)))
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then the response should be the new created author", func() {
				So(rr.Code, ShouldEqual, http.StatusOK)
				var aResp AuthorResponse
				err = json.NewDecoder(rr.Body).Decode(&aResp)
				So(err, ShouldBeNil)
				So(len(aResp.Authors), ShouldEqual, 1)
				author := aResp.Authors[0]
				So(author.ID, ShouldNotBeNil)

				Convey("And the a new created author should be get", func() {
					req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/author/%s", author.ID), nil)
					rr := httptest.NewRecorder()
					handler.ServeHTTP(rr, req)
					So(rr.Code, ShouldEqual, http.StatusOK)
					var aResp AuthorResponse
					err = json.NewDecoder(rr.Body).Decode(&aResp)
					So(err, ShouldBeNil)
					So(len(aResp.Authors), ShouldEqual, 1)
					So(aResp.Authors[0], ShouldResemble, author)
				})
			})
		})
	})
}
