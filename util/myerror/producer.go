package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrProducerGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "40000",
		Message:   "Failed to get producer.",
		IsSentry:  true,
	}
}

func ErrProducerCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "40001",
		Message:   "Failed to create producer.",
		IsSentry:  true,
	}
}

func ErrProducerUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "40002",
		Message:   "Failed to update producer.",
		IsSentry:  true,
	}
}

func ErrProducerDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "40003",
		Message:   "Failed to delete producer.",
		IsSentry:  true,
	}
}

func ErrProducerNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "40004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrProducerInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "40005",
		Message:   fmt.Sprintf("Invalid parameter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrProducerConflictUniqueConstraint(message string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "40006",
		Message:   message,
		IsSentry:  false,
	}
}
