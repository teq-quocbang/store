package storage

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
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
		resp, ctx := setupGetList(&payload.GetStorageByLocatRequest{
			Locat: locat,
		})

		// Act
		err := r.GetList(ctx)

		// Assert
		assertion.NoError(err)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}
}

func setupGetList(input *payload.GetStorageByLocatRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/storage", bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
