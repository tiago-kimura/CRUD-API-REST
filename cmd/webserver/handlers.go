package main

import (
	"Github/desafio-dev-api-rest/account"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	headerKey   = "Content-Type"
	headerValue = "application/json"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok!"))
}

func writeResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func createAccountHolderHandler(service account.Service) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		body, errb := ioutil.ReadAll(r.Body)
		if errb != nil {
			w.Write([]byte(errb.Error()))
		}
		var accountHolder account.AccountHolder
		json.Unmarshal(body, &accountHolder)
		accountHolder, errc := service.CreateAccountHolder((accountHolder))
		if errc != nil {
			accountError := account.AccountError{
				Description: errc.Error(),
			}
			body, _ := json.Marshal(accountError)
			writeResponse(w, http.StatusBadRequest, body)
			return
		}
		body, _ = json.Marshal(accountHolder)

		writeResponse(w, http.StatusOK, body)
	}
}

func findAccountHolderHandler(service account.Service) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerKey, headerValue)
		vars := mux.Vars(r)
		accountHolder, errfi := service.FindAccountHolderByDocument(vars["document"])
		if errfi != nil {
			accountError := account.AccountError{
				Description: errfi.Error(),
			}
			body, _ := json.Marshal(accountError)
			writeResponse(w, http.StatusNotFound, body)
			return
		}
		body, _ := json.Marshal(accountHolder)
		writeResponse(w, http.StatusOK, body)
	}
}

func listAllAccountHolderHandler(service account.Service) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerKey, headerValue)
		listAccountHolder := service.GetAccountHolders()
		body, _ := json.Marshal(listAccountHolder)
		writeResponse(w, http.StatusOK, body)
	}
}

func RemoveAccountHolderByDocumentHandler(service account.Service) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		errrem := service.RemoveAccountHolderByDocument(vars["document"])
		if errrem != nil {
			accountError := account.AccountError{
				Description: errrem.Error(),
			}
			body, _ := json.Marshal(accountError)
			writeResponse(w, http.StatusNotFound, body)
			return
		}
		body, _ := json.Marshal("{}")
		writeResponse(w, http.StatusOK, body)
	}
}
