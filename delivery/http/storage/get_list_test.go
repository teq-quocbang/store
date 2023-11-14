package storage

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
)

func TestGetList(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	_, _, _, locat, err := SetUpForeignKeyData(db)
	assertion.NoError(err)

	// good case
	{
		// Arrange
		resp, ctx := setupGetList(locat)

		// Act
		err := r.GetList(ctx)

		// Assert
		assertion.NoError(err)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		resp, ctx := setupGetList("")

		// Act
		r.GetList(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupGetList(locat string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/storage", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("locat")
	c.SetParamValues(locat)

	return rec, c
}
