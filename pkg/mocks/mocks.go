package mocks

import (
//	  "github.com/stretchr/testify/mock"
)

func HashPasswordMock(password string) (string, error) {
	// this monkey patched mock just returns the password
	return password, nil
}

func CheckPasswordHashMock(password, hash string) bool {
	return true
}
