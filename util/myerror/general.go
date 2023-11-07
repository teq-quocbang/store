package myerror

import (
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrSendEmail(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusNotAcceptable,
		ErrorCode: "001",
		Message:   err.Error(),
		IsSentry:  true,
	}
}

func ErrCommitTransaction(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "002",
		Message:   "Failed to commit transaction",
		IsSentry:  true,
	}
}

func ErrJSONMarshal(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "003",
		Message:   "Failed to json marshal",
		IsSentry:  true,
	}
}

func ErrJSONUnmarshal(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "003",
		Message:   "Failed to json unmarshal",
		IsSentry:  true,
	}
}

func ErrForbidden(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusForbidden,
		ErrorCode: "004",
		Message:   "Failed to access",
		IsSentry:  false,
	}
}
