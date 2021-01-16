package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tebrizetayi/cleanarchitecture/businessservice/contract"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorHandler struct {
	authorBS contract.AuthorBS
}

type AuthorResponse struct {
	Authors []model.Author `json:"authors"`
}

type AuthorRequest struct {
	Authors []model.Author `json:"authors"`
}

func NewAuthorHandler(aBS contract.AuthorBS) AuthorHandler {
	return AuthorHandler{
		authorBS: aBS,
	}
}

func (ac *AuthorHandler) Get(w http.ResponseWriter, r *http.Request) {
	idInput, ok := mux.Vars(r)["id"]
	var authors []model.Author
	var err error
	if ok {
		id, err := strconv.Atoi(idInput)
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
		authors, err = ac.authorBS.GetByIds([]int{id})
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
	} else {
		authors, err = ac.authorBS.GetAll()
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
	}
	resultResponse(w, authors, http.StatusOK)
}

func (ac *AuthorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idInput, ok := mux.Vars(r)["id"]
	var err error
	if ok {
		id, err := strconv.Atoi(idInput)
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
		err = ac.authorBS.Delete([]int{id})
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
	} else {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}
	resultResponse(w, "", http.StatusOK)
}

func (ac *AuthorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var aReq AuthorRequest = AuthorRequest{}
	err := json.NewDecoder(r.Body).Decode(&aReq)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	aResp := AuthorResponse{}
	aResp.Authors, err = ac.authorBS.Create(aReq.Authors)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}
	resultResponse(w, aResp, http.StatusOK)
}

func errorResponse(w http.ResponseWriter, err error, statuscode int) {
	w.WriteHeader(statuscode)
	fmt.Fprint(w, err)
}

func resultResponse(w http.ResponseWriter, data interface{}, status int) {

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
