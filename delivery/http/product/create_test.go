package product

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
	"github.com/teq-quocbang/store/delivery/http/auth"
	"github.com/teq-quocbang/store/delivery/http/producer"
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

	accountID, producerID, err := SetUpForeignKeyData(db)
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
		req := &payload.CreateProductRequest{
			Name:        gofakeit.Name(),
			ProductType: gofakeit.Car().Type,
			ProducerID:  producerID.String(),
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
		req := &payload.CreateProductRequest{}
		resp, ctx := setupCreate(req)

		// Act
		r.Create(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupCreate(input *payload.CreateProductRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/product", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func SetUpForeignKeyData(db *database.Database) (uuid.UUID, uuid.UUID, error) {
	repo := repository.New(db.GetClient)
	rAccount := account.Route{
		UseCase: usecase.New(repo, nil),
	}
	rProducer := producer.Route{
		UseCase: usecase.New(repo, nil),
	}

	// create account
	resp, ctx := setUpTestSignUp(&payload.SignUpRequest{
		Username: gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Name(),
	})
	err := rAccount.SignUp(ctx)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	if resp.Code != 200 {
		return uuid.UUID{}, uuid.UUID{}, fmt.Errorf("failed to sign up, error: %v", resp.Body)
	}
	accountResponse, err := test.UnmarshalBody[*presenter.AccountResponseWrapper](resp.Body.Bytes())
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}

	// define user principle
	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			Username: gofakeit.Name(),
			ID:       accountResponse.Account.ID,
			Email:    gofakeit.Email(),
		},
	}

	// create producer
	strProducerID := "2e51ab2e-11a0-4c7e-823e-b5643e40489b"
	producerID, err := uuid.Parse(strProducerID)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, fmt.Errorf("failed to parse producerID, error: %v", err)
	}
	producerUUID := monkey.Patch(uuid.New, func() uuid.UUID {
		return producerID
	})

	producerResp, ctx := setupCreateProducer(&payload.CreateProducerRequest{
		Name:    gofakeit.Name(),
		Country: gofakeit.Country(),
	})
	ctx.Set(string(auth.UserPrincipleKey), userPrinciple)
	err = rProducer.Create(ctx)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	if resp.Code != 200 {
		return uuid.UUID{}, uuid.UUID{}, fmt.Errorf("failed to create producer, error: %v", producerResp.Body)
	}
	producerResponse, err := test.UnmarshalBody[*presenter.ProducerResponseWrapper](producerResp.Body.Bytes())
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	monkey.Unpatch(producerUUID)

	return accountResponse.Account.ID, producerResponse.Producer.ID, nil
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

func setupCreateProducer(input *payload.CreateProducerRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/producer", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
