package account

import (
	"reflect"
	"testing"
)

func TestNewAccountService(t *testing.T) {

	tests := []struct {
		name       string
		repository Repository
		want       AccountService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountService(tt.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountService() = %v, want %v", got, tt.want)
			}
		})
	}
}

type repositoryMock struct {
}
