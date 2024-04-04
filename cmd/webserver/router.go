package main

import (
	"Github/desafio-dev-api-rest/account"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type handler func(w http.ResponseWriter, r *http.Request)

// Start ...
func Start(port string, service account.Service) {
	router := mux.NewRouter()
	router.HandleFunc("/health", healthHandler).Methods("GET")
	router.HandleFunc("/createAccountHolder", createAccountHolderHandler(service)).Methods("POST")
	router.HandleFunc("/findAccountHolder/{document}", findAccountHolderHandler(service)).Methods("GET")
	router.HandleFunc("/listAccountHolders", listAllAccountHolderHandler(service)).Methods("GET")
	router.HandleFunc("/removeAccountHolder/{document}", RemoveAccountHolderByDocumentHandler(service)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
