package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory"
)

func TestAuthor(t *testing.T) {
	Convey("Sending Get Request", t, func() {
		req, err := http.NewRequest(http.MethodGet, "/author", nil)
		So(err, ShouldBeNil)
		rr := httptest.NewRecorder()
		authorRepo := inmemory.NewAuthorInmemoryRepo()
		authorBS := businessservice.NewAuthorBS(&authorRepo)
		handler := InitRoute(NewAuthorHandler(&authorBS))
		Convey("When sending request", func() {
			handler.ServeHTTP(rr, req)
			So(rr.Code, ShouldEqual, http.StatusOK)
		})

	})
}
