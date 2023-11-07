package hashing

import (
	"errors"

	"github.com/teq-quocbang/store/util/myerror"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrWrongPassword error = errors.New("wrong password")
)

func ToHashPassword(password string) ([]byte, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, myerror.ErrAccountCreate(err)
	}
	return hashPassword, nil
}

func CompareHashPassword(password string, hashedPassword []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrWrongPassword
		}
		return err
	}
	return nil
}
