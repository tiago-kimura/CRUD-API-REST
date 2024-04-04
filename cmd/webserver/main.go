package main

import "Github/desafio-dev-api-rest/account"

func main() {
	//cfg := newConfig()
	repository := account.NewInMemoryRepository()
	service := account.NewAccountService(&repository)
	Start("8080", service)
}
