package product

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/test"
	"github.com/teq-quocbang/store/util/token"
)

func TestUpdate(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}
	testProductID := uuid.New()
	testProductType := gofakeit.Car().Type
	testProductName := gofakeit.Name()

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
		createProduct := &payload.CreateProductRequest{
			Name:        gofakeit.Name(),
			ProductType: gofakeit.Car().Type,
			ProducerID:  producerID.String(),
		}
		resp, ctx := setupCreate(createProduct)
		err := r.Create(ctx)
		assertion.NoError(err)
		actualResponse, err := test.UnmarshalBody[*presenter.ProductResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)

		req := &payload.UpdateProductRequest{
			ID:          actualResponse.Product.ID.String(),
			Name:        testProductName,
			ProductType: testProductType,
			ProducerID:  producerID.String(),
		}
		resp, ctx = setupUpdate(req, actualResponse.Product.ID.String())

		// Act
		err = r.Update(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		req := &payload.UpdateProductRequest{
			ID:          testProductID.String(),
			Name:        testProductName,
			ProductType: testProductType,
			ProducerID:  producerID.String(),
		}
		resp, ctx := setupUpdate(req, "")

		// Act
		r.Update(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupUpdate(input *payload.UpdateProductRequest, id string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/api/product", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	return rec, c
}
