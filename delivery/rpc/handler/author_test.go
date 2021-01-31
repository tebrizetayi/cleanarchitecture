package handler

import (
	"context"
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/delivery/rpc/proto"
	"github.com/tebrizetayi/cleanarchitecture/repository/mysql"
	"github.com/tebrizetayi/cleanarchitecture/testutils"
)

func TestAuthorProto(t *testing.T) {
	db, err := testutils.MysqlDBConnection()
	if err != nil {
		t.Error(err)
		return
	}
	authorRepo := mysql.NewAuthorMysqlRepo(db)
	authorBS := businessservice.NewAuthorBS(&authorRepo)
	srv := NewAuthorHandler(&authorBS)
	AuthorProtoTestStructure(t, &srv)
}

func AuthorProtoTestStructure(t *testing.T, srv *AuthorHandler) {
	Convey("When you create a new author", t, func() {
		author := proto.Author{
			Name: "TestName",
		}
		pbAuthors, err := srv.Create(context.Background(), &proto.Authors{Author: []*proto.Author{&author}})
		Convey("Then the response should be the new created author", func() {
			So(err, ShouldBeNil)
			So(pbAuthors, ShouldNotBeNil)
			So(len(pbAuthors.Author), ShouldEqual, 1)
			pbAuthor := pbAuthors.Author[0]
			So(pbAuthor.Name, ShouldEqual, author.Name)
			So(pbAuthor.Id, ShouldNotEqual, uuid.Nil)
			Convey("And the a new created author should be get", func() {
				pbAuthorID := proto.AuthorId{
					Id: pbAuthor.Id,
				}
				pbAuthorIds := proto.AuthorIds{
					AuthorId: []*proto.AuthorId{&pbAuthorID},
				}
				pbAuthors, err := srv.GetByIds(context.Background(), &pbAuthorIds)
				So(err, ShouldBeNil)
				So(len(pbAuthors.Author), ShouldEqual, 1)
				So(pbAuthors.Author[0], ShouldResemble, pbAuthor)
			})
		})
	})
}
