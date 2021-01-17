package main

import (
	"log"
	"net/http"

	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/delivery/api"
	"github.com/tebrizetayi/cleanarchitecture/repository/inmemory"
)

func main() {

	authorRepository := inmemory.NewAuthorInmemoryRepo()
	authorService := businessservice.NewAuthorBS(&authorRepository)
	authorhandler := api.NewAuthorHandler(&authorService)
	r := api.InitRoutes(&authorhandler)
	log.Println("Server begin to run in 9090 port")
	http.ListenAndServe(":9090", r)

}
