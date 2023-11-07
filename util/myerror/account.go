package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrAccountGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "20000",
		Message:   "Failed to get account.",
		IsSentry:  true,
	}
}

func ErrAccountCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "20001",
		Message:   "Failed to create account.",
		IsSentry:  true,
	}
}

func ErrAccountUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "20002",
		Message:   "Failed to update account.",
		IsSentry:  true,
	}
}

func ErrAccountDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "20003",
		Message:   "Failed to delete account.",
		IsSentry:  true,
	}
}

func ErrAccountNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "20004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrAccountInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "20005",
		Message:   fmt.Sprintf("Invalid parameter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrAccountConflictUniqueConstraint(message string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "20006",
		Message:   message,
		IsSentry:  false,
	}
}

func ErrAccountGenerateToken(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "20007",
		Message:   "failed to generate token",
		IsSentry:  false,
	}
}

func ErrAccountComparePassword(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		ErrorCode: "20008",
		HTTPCode:  http.StatusForbidden,
		Message:   "Failed to compare password",
		IsSentry:  false,
	}
}
