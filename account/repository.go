package account

import "fmt"

type Repository interface {
	PersistAccountHolder(accountHolder AccountHolder) error
	GetAccountHolderByDocument(document string) (AccountHolder, error)
	RemoveAccountHolder(document string) error
	ListAllAccountHolder() []AccountHolder
}

// InMemoryRepository implements repository interface storing data on memory
//Represents a list of Account Holder ///TODO:and Accout
type InMemoryRepository struct {
	ListAccountHolder []AccountHolder
	///TODO: List Account
}

//Constructor of InMemoryRepository
func NewInMemoryRepository() InMemoryRepository {
	fmt.Println("Passando aqui!")
	return InMemoryRepository{
		ListAccountHolder: []AccountHolder{},
		///TODO: List Account
	}
}

func (i *InMemoryRepository) GetAccountHolderByDocument(document string) (AccountHolder, error) {
	var accountHolder AccountHolder
	for _, v := range i.ListAccountHolder {
		if v.HolderDocument == document {
			accountHolder = v
			return accountHolder, nil
		}
	}
	return accountHolder, ErrDocumentNotFound
}

func (i *InMemoryRepository) PersistAccountHolder(accountHolder AccountHolder) error {
	listHolders := append(i.ListAccountHolder, accountHolder)
	i.ListAccountHolder = listHolders
	fmt.Println("Pers i.List: \n", i.ListAccountHolder)
	return nil
}

func (i *InMemoryRepository) RemoveAccountHolder(document string) error {
	indice := -1
	for i, holder := range i.ListAccountHolder {
		if holder.HolderDocument == document {
			indice = i
			break
		}
	}
	if indice < 0 {
		return ErrDocumentNotFound
	}
	listAccount := append(i.ListAccountHolder[0:indice], i.ListAccountHolder[indice+1:]...)
	i.ListAccountHolder = listAccount
	return nil
}

func (i *InMemoryRepository) ListAllAccountHolder() []AccountHolder {
	return i.ListAccountHolder
}
