package account

import "time"

type AccountHolder struct {
	HolderName     string    `json:"holderName"`
	HolderDocument string    `json:"holderDocument"`
	DateCreation   time.Time `json:"dateCreation"`
}

type Account struct {
	AccountHolder AccountHolder
	Balance       float64                `json:"balance"`
	Number        int                    `json:"number"`
	Branch        int                    `json:"branch"`
	DateCreation  time.Time              `json:"dateCreation"`
	DateUpdate    time.Time              `json:"dateUpdate"`
	Status        int                    `json:"status"`
	Active        bool                   `json:"active"`
	Transactions  map[string]Transaction `json:"transactions"`
}

type Transaction struct {
	Id              string    `json:"id"`
	Type            int       `json:"type"`
	Amount          float64   `json:"amount"`
	DateTransaction time.Time `json:"dateTransaction"`
	Status          int       `json:"status"`
}
