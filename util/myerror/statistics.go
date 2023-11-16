package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrStatisticsGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30000",
		Message:   "Failed to get statistics.",
		IsSentry:  true,
	}
}

func ErrStatisticsCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30001",
		Message:   "Failed to create statistics.",
		IsSentry:  true,
	}
}

func ErrStatisticsUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30002",
		Message:   "Failed to update statistics.",
		IsSentry:  true,
	}
}

func ErrStatisticsDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30003",
		Message:   "Failed to delete statistics.",
		IsSentry:  true,
	}
}

func ErrStatisticsNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "30004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrStatisticsInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "30005",
		Message:   fmt.Sprintf("Invalid parameter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrStatisticsConflictUniqueConstraint(message string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30006",
		Message:   message,
		IsSentry:  false,
	}
}

func ErrStatisticsExportFailed(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "30006",
		Message:   "Failed to export file",
		IsSentry:  false,
	}
}
