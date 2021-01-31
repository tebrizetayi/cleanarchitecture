package main

import (
	"log"
	"net"

	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/delivery/rpc/handler"
	"github.com/tebrizetayi/cleanarchitecture/delivery/rpc/proto"
	"github.com/tebrizetayi/cleanarchitecture/repository/mysql"
	"github.com/tebrizetayi/cleanarchitecture/testutils"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	db, err := testutils.MysqlDBConnection()
	if err != nil {
		panic(err)
	}

	authorRepo := mysql.NewAuthorMysqlRepo(db)
	authorBS := businessservice.NewAuthorBS(&authorRepo)
	authorHandler := handler.NewAuthorHandler(&authorBS)
	grpcServer := grpc.NewServer()
	proto.RegisterAuthorGrpcServer(grpcServer, &authorHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
