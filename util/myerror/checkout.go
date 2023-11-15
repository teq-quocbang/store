package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrCartGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "60000",
		Message:   "Failed to get cart.",
		IsSentry:  true,
	}
}

func ErrCartCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "60001",
		Message:   "Failed to create cart.",
		IsSentry:  true,
	}
}

func ErrCartUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "60002",
		Message:   "Failed to update cart.",
		IsSentry:  true,
	}
}

func ErrCartDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "60003",
		Message:   "Failed to delete cart.",
		IsSentry:  true,
	}
}

func ErrCartNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "60004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrCartInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "60005",
		Message:   fmt.Sprintf("Invalid paramemter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrCustomerOrderGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "61000",
		Message:   "Failed to get customer order.",
		IsSentry:  true,
	}
}

func ErrCustomerOrderCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "61001",
		Message:   "Failed to create customer order.",
		IsSentry:  true,
	}
}

func ErrCustomerOrderUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "61002",
		Message:   "Failed to update customer order.",
		IsSentry:  true,
	}
}

func ErrCustomerOrderDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "61003",
		Message:   "Failed to delete customer order.",
		IsSentry:  true,
	}
}

func ErrCustomerOrderNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "61004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrCustomerOrderInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "61005",
		Message:   fmt.Sprintf("Invalid paramemter: `%s`.", param),
		IsSentry:  false,
	}
}
