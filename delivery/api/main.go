package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/tebrizetayi/cleanarchitecture/businessservice"
	"github.com/tebrizetayi/cleanarchitecture/repository/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/Academia")
	if err != nil {
		panic(err)
	}
	authorRepository := mysql.NewAuthorMysqlRepo(db)

	authorService := businessservice.NewAuthorBS(&authorRepository)
	authorhandler := NewAuthorHandler(&authorService)
	r := InitRoutes(&authorhandler)
	log.Println("Server begin to run in 9090 port")
	http.ListenAndServe(":9090", r)

}
