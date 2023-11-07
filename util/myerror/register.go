package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrRegisterGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50000",
		Message:   "Failed to get register.",
		IsSentry:  true,
	}
}

func ErrRegisterCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50001",
		Message:   "Failed to create register.",
		IsSentry:  true,
	}
}

func ErrRegisterUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50002",
		Message:   "Failed to update register.",
		IsSentry:  true,
	}
}

func ErrRegisterDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "50003",
		Message:   "Failed to delete register.",
		IsSentry:  true,
	}
}

func ErrRegisterNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "50004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrRegisterInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "50005",
		Message:   fmt.Sprintf("Invalid parameter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrCanNotRegisterSameCourse(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "50006",
		Message:   param,
		IsSentry:  false,
	}
}

func ErrFailedToGetCache(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "51001",
		Message:   "failed to get cache",
	}
}

func ErrFailedToSaveCache(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "51002",
		Message:   "failed to save cache",
	}
}

func ErrFailedToRemoveCache(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "51003",
		Message:   "failed to remove cache",
	}
}
