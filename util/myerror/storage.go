package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrStorageGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50000",
		Message:   "Failed to get storage.",
		IsSentry:  true,
	}
}

func ErrStorageCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50001",
		Message:   "Failed to create storage.",
		IsSentry:  true,
	}
}

func ErrStorageUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50002",
		Message:   "Failed to update storage.",
		IsSentry:  true,
	}
}

func ErrStorageDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50003",
		Message:   "Failed to delete storage.",
		IsSentry:  true,
	}
}

func ErrStorageNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "50004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrStorageInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "50005",
		Message:   fmt.Sprintf("Invalid parameter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrStorageConflictUniqueConstraint(message string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50006",
		Message:   message,
		IsSentry:  false,
	}
}
