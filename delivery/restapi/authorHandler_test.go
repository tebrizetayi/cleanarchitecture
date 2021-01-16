package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory"
)

func TestAuthor(t *testing.T) {
	Convey("Sending Get Request", t, func() {
		authorRepo := inmemory.NewAuthorInmemoryRepo()
		authorBS := businessservice.NewAuthorBS(&authorRepo)
		authorHandler := NewAuthorHandler(&authorBS)
		handler := InitRoutes(&authorHandler)

		Convey("When adding multiple authors", func() {
			aReq := []byte(`{"authors":[{"name":"Jack London"},{"name":"William"},{"name":"Arnold"},{"name":"Victor Hugo"}]}`)
			req, err := http.NewRequest(http.MethodPost, "/author", bytes.NewBuffer(aReq))
			So(err, ShouldBeNil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			Convey("Then the status code should be ok", func() {
				So(rr.Code, ShouldEqual, http.StatusOK)
				Convey("Then the returned count should be more than 0", func() {
					var aResp AuthorResponse
					data, err := ioutil.ReadAll(rr.Body)
					err = json.Unmarshal(data, &aResp)
					So(err, ShouldBeNil)
					So(len(aResp.Authors), ShouldBeGreaterThan, 0)
					Convey("When deleting existing author", func() {
						author := aResp.Authors[0]
						req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/author/%d", author.ID), nil)
						So(err, ShouldBeNil)
						handler.ServeHTTP(rr, req)
						Convey("Then it the result should be empty", func() {
							So(rr.Code, ShouldEqual, http.StatusOK)
						})
					})

				})
			})
		})
	})
}
