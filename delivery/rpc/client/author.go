package main

import (
	"context"
	"log"

	"github.com/tebrizetayi/cleanarchitecture/delivery/rpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to grpc server %s\n", err)
	}
	defer conn.Close()

	c := proto.NewAuthorGrpcClient(conn)

	authors := []*proto.Author{}
	authors = append(authors, &proto.Author{Name: "Tabriz"})

	outAuthors, err := c.Create(context.Background(), &proto.Authors{Author: authors})

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Println(outAuthors)
}
