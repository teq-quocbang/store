package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/teq-quocbang/store/delivery/http/account"
	"github.com/teq-quocbang/store/delivery/http/producer"
	"github.com/teq-quocbang/store/delivery/http/product"
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

	accountID, productID, err := SetUpForeignKeyData(db)
	assertion.NoError(err)

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			Username: fake.Name(),
			ID:       accountID,
			Email:    fake.Email(),
		},
	}
	monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return userPrinciple
	})

	defer monkey.UnpatchAll()

	// good case
	{
		// Arrange
		req := &payload.UpsertStorageRequest{
			Locat:     fmt.Sprintf("%s%d", "A", fake.IntRange(100, 1000)),
			ProductID: productID.String(),
			Qty:       int64(fake.Int8()),
		}
		resp, ctx := setupUpsert(req)

		// Act
		err = r.Upsert(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		req := &payload.UpsertStorageRequest{}
		resp, ctx := setupUpsert(req)

		// Act
		r.Upsert(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupUpsert(input *payload.UpsertStorageRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/storage", bytes.NewReader(b))
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
	rProduct := product.Route{
		UseCase: usecase.New(repo, nil),
	}

	resp, ctx := setUpTestSignUp(&payload.SignUpRequest{
		Username: fake.Name(),
		Email:    fake.Email(),
		Password: fake.Name(),
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

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			Username: fake.Name(),
			ID:       accountResponse.Account.ID,
			Email:    fake.Email(),
		},
	}
	monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return userPrinciple
	})

	producerResp, ctx := setupCreateProducer(&payload.CreateProducerRequest{
		Name:    fake.Name(),
		Country: fake.Country(),
	})
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

	productResp, ctx := setupCreateProduct(&payload.CreateProductRequest{
		Name:        fake.Name(),
		ProducerID:  producerResponse.Producer.ID.String(),
		ProductType: fake.Car().Type,
	})
	err = rProduct.Create(ctx)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	productResponse, err := test.UnmarshalBody[*presenter.ProductResponseWrapper](productResp.Body.Bytes())
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}

	return accountResponse.Account.ID, productResponse.Product.ID, nil
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

func setupCreateProduct(input *payload.CreateProductRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/product", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
