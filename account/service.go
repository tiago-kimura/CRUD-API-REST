package account

import (
	"errors"
	"time"
)

////var ListAccountHolder []AccountHolder

// interface that defines account service
type Service interface {
	CreateAccountHolder(accountHolder AccountHolder) (AccountHolder, error)
	FindAccountHolderByDocument(document string) (AccountHolder, error)
	GetAccountHolders() []AccountHolder
	RemoveAccountHolderByDocument(document string) error
}

// implements service interface
type AccountService struct {
	repository Repository
}

// constructor
func NewAccountService(repository Repository) AccountService {
	return AccountService{
		repository: repository,
	}
}

func (a AccountService) CreateAccountHolder(accountHolder AccountHolder) (AccountHolder, error) {
	errv := obrigatoryFields(accountHolder)
	if errv != nil {
		return accountHolder, errv
	}
	accountHolder.HolderDocument = sanitize(accountHolder.HolderDocument)
	_, errfi := a.FindAccountHolderByDocument(accountHolder.HolderDocument)
	if errfi == nil {
		return accountHolder, errors.New("document already exists!")
	}
	validado := ValidateCPF(accountHolder.HolderDocument)
	if !validado {
		return accountHolder, errors.New("invalid Document")
	}
	accountHolder.DateCreation = time.Now()
	errper := a.repository.PersistAccountHolder(accountHolder)
	if errper != nil {
		return accountHolder, errper
	}
	return accountHolder, nil
}

func (a AccountService) FindAccountHolderByDocument(document string) (AccountHolder, error) {
	document = sanitize(document)
	accountHolder, errget := a.repository.GetAccountHolderByDocument(document)
	if errget != nil {
		return accountHolder, errget
	}
	return accountHolder, nil
}

func (a AccountService) GetAccountHolders() []AccountHolder {
	return a.repository.ListAllAccountHolder()
}

func (a AccountService) RemoveAccountHolderByDocument(document string) error {
	document = sanitize(document)
	errre := a.repository.RemoveAccountHolder(document)
	if errre != nil {
		return errre
	}
	return nil
}
