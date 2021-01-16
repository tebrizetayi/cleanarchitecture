package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory"
)

func TestAuthorHandler(t *testing.T) {
	Convey("Setup", t, func() {
		authorRepo := inmemory.NewAuthorInmemoryRepo()
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
			authorJSON := fmt.Sprintf(`{"authors":[{"name":"%s","id":%d}]}`, author.Name, author.ID)
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/author/%d", author.ID), bytes.NewBuffer([]byte(authorJSON)))
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then updated author can be get by id", func() {
				So(rr.Code, ShouldEqual, http.StatusOK)
				var aRes AuthorResponse
				err := json.NewDecoder(rr.Body).Decode(&aRes)
				So(err, ShouldBeNil)
				So(author, ShouldResemble, aRes.Authors[0])
			})
		})
		Convey("When you update author where the id is not in the database", func() {
			author := model.Author{
				Name: "John Doe",
				ID:   math.MaxInt32,
			}
			authorJSON := fmt.Sprintf(`{"authors":[{"name":"%s","id":%d}]}`, author.Name, author.ID)
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/author/%d", author.ID), bytes.NewBuffer([]byte(authorJSON)))
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then updated author can be get by id", func() {
				So(rr.Code, ShouldEqual, http.StatusOK)
				var aRes AuthorResponse
				err := json.NewDecoder(rr.Body).Decode(&aRes)
				So(err, ShouldBeNil)
				So(author.Name, ShouldEqual, aRes.Authors[0].Name)
				//Because id is autoincremented property
				So(author.ID, ShouldNotEqual, aRes.Authors[0].ID)
			})
		})

		Convey("When you update author with wrong json", func() {

			authorJSON := `{"authors":[{"name"}]}`
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/author/%d", author.ID), bytes.NewBuffer([]byte(authorJSON)))
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then updated author can be get by id", func() {
				So(rr.Code, ShouldEqual, http.StatusBadRequest)

			})
		})
		Convey("When you delete author where the id is the database", func() {
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/author/%d", author.ID), nil)
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then the response should be no content", func() {
				So(rr.Code, ShouldEqual, http.StatusNoContent)
			})
		})

	})
}
