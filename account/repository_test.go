package account

import (
	"reflect"
	"testing"
)

func TestInMemoryRepository_ListAllAccountHolder(t *testing.T) {

	tests := []struct {
		name              string
		ListAccountHolder []AccountHolder
		want              []AccountHolder
	}{
		{
			name: "Case with success list",
			ListAccountHolder: []AccountHolder{
				{
					HolderName:     "José",
					HolderDocument: "32165498766",
				}, {
					HolderName:     "Maria",
					HolderDocument: "32165498765",
				}, {
					HolderName:     "Pam",
					HolderDocument: "32165498761",
				},
			},
			want: []AccountHolder{
				{
					HolderName:     "José",
					HolderDocument: "32165498766",
				}, {
					HolderName:     "Maria",
					HolderDocument: "32165498765",
				}, {
					HolderName:     "Pam",
					HolderDocument: "32165498761",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InMemoryRepository{
				ListAccountHolder: tt.ListAccountHolder,
			}
			if got := i.ListAllAccountHolder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryRepository.ListAllAccountHolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryRepository_RemoveAccountHolder(t *testing.T) {

	tests := []struct {
		name              string
		ListAccountHolder []AccountHolder
		document          string
		wantErr           bool
	}{
		{
			name: "Case with success remove",
			ListAccountHolder: []AccountHolder{
				{
					HolderName:     "José",
					HolderDocument: "32165498766",
				}, {
					HolderName:     "Maria",
					HolderDocument: "32165498765",
				}, {
					HolderName:     "Pam",
					HolderDocument: "32165498761",
				},
			},
			document: "32165498766",
			wantErr:  false,
		},
		{
			name: "Case with fail remove",
			ListAccountHolder: []AccountHolder{
				{
					HolderName:     "José",
					HolderDocument: "32165498766",
				}, {
					HolderName:     "Maria",
					HolderDocument: "32165498765",
				}, {
					HolderName:     "Pam",
					HolderDocument: "32165498761",
				},
			},
			document: "32165498763",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InMemoryRepository{
				ListAccountHolder: tt.ListAccountHolder,
			}
			if err := i.RemoveAccountHolder(tt.document); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.RemoveAccountHolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryRepository_PersistAccountHolder(t *testing.T) {

	tests := []struct {
		name              string
		ListAccountHolder []AccountHolder
		accountHolder     AccountHolder
		wantErr           bool
	}{
		{
			name: "case with success creation",
			accountHolder: AccountHolder{
				HolderName:     "Pedro",
				HolderDocument: "98765432155",
			},
			ListAccountHolder: []AccountHolder{
				{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InMemoryRepository{
				ListAccountHolder: tt.ListAccountHolder,
			}
			if err := i.PersistAccountHolder(tt.accountHolder); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.PersistAccountHolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryRepository_GetAccountHolderByDocument(t *testing.T) {
	tests := []struct {
		name              string
		ListAccountHolder []AccountHolder
		document          string
		want              AccountHolder
		wantErr           bool
	}{
		{
			name: "case with success get",
			ListAccountHolder: []AccountHolder{
				{
					HolderName:     "José",
					HolderDocument: "32165498766",
				}, {
					HolderName:     "Maria",
					HolderDocument: "32165498765",
				}, {
					HolderName:     "Pam",
					HolderDocument: "32165498761",
				},
			},
			document: "32165498766",
			want: AccountHolder{
				HolderName:     "José",
				HolderDocument: "32165498766",
			},
			wantErr: false,
		}, {
			name: "Case with fail get",
			ListAccountHolder: []AccountHolder{
				{
					HolderName:     "José",
					HolderDocument: "32165498766",
				}, {
					HolderName:     "Maria",
					HolderDocument: "32165498765",
				}, {
					HolderName:     "Pam",
					HolderDocument: "32165498761",
				},
			},
			document: "32165498763",
			want:     AccountHolder{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InMemoryRepository{
				ListAccountHolder: tt.ListAccountHolder,
			}
			got, err := i.GetAccountHolderByDocument(tt.document)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.GetAccountHolderByDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryRepository.GetAccountHolderByDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInMemoryRepository(t *testing.T) {
	tests := []struct {
		name string
		want InMemoryRepository
	}{
		{
			name: "success constructor",
			want: InMemoryRepository{
				[]AccountHolder{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemoryRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemoryRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
