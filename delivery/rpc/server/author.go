package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/businessservice/contract"
	author "github.com/tebrizetayi/cleanarchitecture/delivery/rpc/proto"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	/*db, err = testutils.MysqlDBConnection()
	if err != nil {
		panic(err)
	}
	*/
	//authorRepo := mysql.NewAuthorMysqlRepo(db)
	authorRepo := inmemory.NewAuthorInmemoryRepo()
	authorBS := businessservice.NewAuthorBS(&authorRepo)
	s := NewServer(&authorBS)
	grpcServer := grpc.NewServer()
	author.RegisterAuthorGrpcServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}

type Server struct {
	AuthorBS contract.AuthorBS
}

func NewServer(authorBs contract.AuthorBS) Server {

	return Server{
		AuthorBS: authorBs,
	}
}

//GetByIds(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*Author, error)
func (a *Server) GetByIds(ctx context.Context, authorID *author.AuthorIds) (*author.Authors, error) {

	ids := []uuid.UUID{}
	for i := 0; i < len(authorID.AuthorId); i++ {
		id, err := uuid.Parse(authorID.AuthorId[i].Id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	authors, err := a.AuthorBS.GetByIds(ids)

	if err != nil {
		return nil, err
	}

	result := author.Authors{}
	for i := 0; i < len(authors); i++ {
		result.Author = append(result.Author, &author.Author{
			Id:   authors[i].ID.String(),
			Name: authors[i].Name,
		})
	}

	return &result, nil
}

func (a *Server) Create(ctx context.Context, inAuthors *author.Authors) (*author.Authors, error) {
	authors := []model.Author{}
	for i := 0; i < len(inAuthors.Author); i++ {
		authors = append(authors, model.Author{Name: inAuthors.Author[i].Name})
	}

	authors, err := a.AuthorBS.Create(authors)
	if err != nil {
		return nil, err
	}

	result := author.Authors{}
	for i := 0; i < len(authors); i++ {
		result.Author = append(result.Author, &author.Author{
			Id:   authors[i].ID.String(),
			Name: authors[i].Name,
		})
	}

	return &result, nil
}
