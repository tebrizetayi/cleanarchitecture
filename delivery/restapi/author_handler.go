package api

/*
"http://localhost:8080/Authors/"
"http://localhost:8080/Authors/Id:3"  //HttpMethod Delete
"http://localhost:8080/Authors/Id:3"  //HttpMethod Get
"http://localhost:8080/Authors/Id:3"  //HttpMethod Update
"http://localhost:8080/Authors/Id:3/Articles"  //HttpMethod Update
*/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tebrizetayi/cleanarchitecture/businessservice/contract"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorHandler struct {
	contract.AuthorBS
}

type AuthorRequest struct {
	model.Author
}

type AuthorResponse struct {
	model.Author
}

func NewAuthorHandler(authorBs contract.AuthorBS) AuthorHandler {
	return AuthorHandler{AuthorBS: authorBs}
}

func (a *AuthorHandler) Get(w http.ResponseWriter, r *http.Request) {

	//decoder := json.NewDecoder(r.Body)
	//var authorRequest AuthorRequest
	//decoder.Decode(&authorRequest)

	authors, _ := a.AuthorBS.GetAll()

	encoder := json.NewEncoder(w)
	encoder.Encode(authors)

	fmt.Fprintln(w, authors)
}
