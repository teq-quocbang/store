package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrSemesterGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30000",
		Message:   "Failed to get semester.",
		IsSentry:  true,
	}
}

func ErrSemesterCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30001",
		Message:   "Failed to create semester.",
		IsSentry:  true,
	}
}

func ErrSemesterUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30002",
		Message:   "Failed to update semester.",
		IsSentry:  true,
	}
}

func ErrSemesterDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30003",
		Message:   "Failed to delete semester.",
		IsSentry:  true,
	}
}

func ErrSemesterNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "30004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrSemesterInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "30005",
		Message:   fmt.Sprintf("Invalid parameter: `%s`.", param),
		IsSentry:  false,
	}
}
