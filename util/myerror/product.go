package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrProductGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30000",
		Message:   "Failed to get product.",
		IsSentry:  true,
	}
}

func ErrProductCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30001",
		Message:   "Failed to create product.",
		IsSentry:  true,
	}
}

func ErrProductUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30002",
		Message:   "Failed to update product.",
		IsSentry:  true,
	}
}

func ErrProductDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30003",
		Message:   "Failed to delete product.",
		IsSentry:  true,
	}
}

func ErrProductNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "30004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrProductInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "30005",
		Message:   fmt.Sprintf("Invalid parameter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrProductConflictUniqueConstraint(message string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30006",
		Message:   message,
		IsSentry:  false,
	}
}

func ErrProductExportFailed(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30006",
		Message:   "Failed to export file",
		IsSentry:  false,
	}
}
