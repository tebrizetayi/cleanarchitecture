package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/tebrizetayi/cleanarchitecture/businessservice/contract"
	"github.com/tebrizetayi/cleanarchitecture/delivery/rpc/proto"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorHandler struct {
	AuthorBS contract.AuthorBS
}

func NewAuthorHandler(authorBs contract.AuthorBS) AuthorHandler {

	return AuthorHandler{
		AuthorBS: authorBs,
	}
}
func (a *AuthorHandler) GetByIds(ctx context.Context, authorID *proto.AuthorIds) (*proto.Authors, error) {

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

	result := proto.Authors{}
	for i := 0; i < len(authors); i++ {
		result.Author = append(result.Author, &proto.Author{
			Id:   authors[i].ID.String(),
			Name: authors[i].Name,
		})
	}

	return &result, nil
}

func (a *AuthorHandler) Create(ctx context.Context, inAuthors *proto.Authors) (*proto.Authors, error) {
	authors := []model.Author{}
	for i := 0; i < len(inAuthors.Author); i++ {
		authors = append(authors, model.Author{Name: inAuthors.Author[i].Name})
	}

	authors, err := a.AuthorBS.Create(authors)
	if err != nil {
		return nil, err
	}

	result := proto.Authors{}
	for i := 0; i < len(authors); i++ {
		result.Author = append(result.Author, &proto.Author{
			Id:   authors[i].ID.String(),
			Name: authors[i].Name,
		})
	}

	return &result, nil
}
