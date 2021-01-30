package main

import (
	"context"
	"log"

	author "github.com/tebrizetayi/cleanarchitecture/delivery/rpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to grpc server %s\n", err)
	}
	defer conn.Close()

	c := author.NewAuthorGrpcClient(conn)

	authors := []*author.Author{}
	authors = append(authors, &author.Author{Name: "Tabriz"})

	outAuthors, err := c.Create(context.Background(), &author.Authors{Author: authors})

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Println(outAuthors)
	/*authorIds := author.AuthorIds{}
	authorIds.AuthorId = append(authorIds.AuthorId, &author.AuthorId{Id: uuid.New().String()})
	response, err := c.GetByIds(context.Background(), &authorIds)

	if err != nil {
		log.Fatalf("Error when calling Author GetByIds")
	}

	log.Printf("Repsone %v", response)
	*/

}
