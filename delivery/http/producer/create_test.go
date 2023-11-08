package producer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/teq-quocbang/store/delivery/http/account"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/test"
	"github.com/teq-quocbang/store/util/token"
)

func TestCreate(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	accountID, err := SetUpForeignKeyData(db)
	assertion.NoError(err)

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			Username: gofakeit.Name(),
			ID:       accountID,
			Email:    gofakeit.Email(),
		},
	}
	monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return userPrinciple
	})

	defer monkey.UnpatchAll()

	// good case
	{
		// Arrange
		req := &payload.CreateProducerRequest{
			Name:    gofakeit.Name(),
			Country: gofakeit.Country(),
		}
		resp, ctx := setupCreate(req)

		// Act
		err = r.Create(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		req := &payload.CreateProducerRequest{}
		resp, ctx := setupCreate(req)

		// Act
		r.Create(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupCreate(input *payload.CreateProducerRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/producer", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func SetUpForeignKeyData(db *database.Database) (uuid.UUID, error) {
	repo := repository.New(db.GetClient)
	rAccount := account.Route{
		UseCase: usecase.New(repo, nil),
	}

	resp, ctx := setUpTestSignUp(&payload.SignUpRequest{
		Username: gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Name(),
	})

	err := rAccount.SignUp(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}
	if resp.Code != 200 {
		return uuid.UUID{}, fmt.Errorf("failed to sign up, error: %v", resp.Body)
	}

	accountResponse, err := test.UnmarshalBody[*presenter.AccountResponseWrapper](resp.Body.Bytes())
	if err != nil {
		return uuid.UUID{}, err
	}

	return accountResponse.Account.ID, nil
}

func setUpTestSignUp(input *payload.SignUpRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/user/sign-up", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
