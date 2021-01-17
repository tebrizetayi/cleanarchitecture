package api

import (
	"encoding/json"
	"errors"
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
	aResp := AuthorResponse{}
	var err error
	if ok {
		id, err := strconv.Atoi(idInput)
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
		aResp.Authors, err = ac.authorBS.GetByIds([]int{id})
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
	} else {
		aResp.Authors, err = ac.authorBS.GetAll()
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
	}
	resultResponse(w, aResp, http.StatusOK)
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
	resultResponse(w, "", http.StatusNoContent)
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

func (ac *AuthorHandler) Update(w http.ResponseWriter, r *http.Request) {

	idInput, ok := mux.Vars(r)["id"]
	if ok {
		id, err := strconv.Atoi(idInput)
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}

		var aReq AuthorRequest = AuthorRequest{}
		err = json.NewDecoder(r.Body).Decode(&aReq)
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
		aReq.Authors[0].ID = id

		aResp := AuthorResponse{}

		//Checking if id exists in the database
		authors, err := ac.authorBS.GetByIds([]int{id})
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}

		//If id is in the database then update
		if len(authors) == 1 {
			//If id is in the database then update
			aResp.Authors, err = ac.authorBS.Update(aReq.Authors)
			if err != nil {
				errorResponse(w, err, http.StatusBadRequest)
				return
			}
		} else {
			//if id is not in the database then add

			aResp.Authors, err = ac.authorBS.Create(aReq.Authors)
			if err != nil {
				errorResponse(w, err, http.StatusBadRequest)
				return
			}
		}
		resultResponse(w, aResp, http.StatusOK)
	} else {
		errorResponse(w, errors.New("No Data"), http.StatusBadRequest)
	}

}

func errorResponse(w http.ResponseWriter, err error, statuscode int) {
	http.Error(w, err.Error(), statuscode)
}

func resultResponse(w http.ResponseWriter, responseData interface{}, status int) {
	data, err := json.Marshal(responseData)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(status)
	w.Write(data)
}
