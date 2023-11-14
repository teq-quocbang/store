package product

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/teq-quocbang/store/delivery/http/auth"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/test"
	"github.com/teq-quocbang/store/util/token"
)

func TestCreateListWithImportFile(t *testing.T) {
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
			Username: fake.Name(),
			ID:       accountID,
			Email:    fake.Email(),
		},
	}

	// good case
	{
		// Arrange
		resp, ctx := setupTestCreateListWithImportFile(producerID, 10)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err := r.CreateListWithImportFile(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}
}

func TestCreateList(t *testing.T) {
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
			Username: fake.Name(),
			ID:       accountID,
			Email:    fake.Email(),
		},
	}

	// good case
	{
		// Arrange
		req := &payload.CreateListProductRequest{
			Products: make([]payload.Product, 10),
		}
		for i := 0; i < 10; i++ {
			req.Products[i] = payload.Product{
				Name:        fake.Name(),
				ProductType: fake.Car().Type,
				ProducerID:  producerID.String(),
			}
		}
		resp, ctx := setupCreateList(req)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err := r.CreateList(ctx)

		// Assert
		assertion.NoError(err)
		actual, err := test.UnmarshalBody[*presenter.ListProductResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		assertion.Equal(10, len(actual.Product))
	}

	// bad case
	{
		// Arrange
		resp, ctx := setupCreateList(&payload.CreateListProductRequest{})

		// Act
		r.CreateList(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setupCreateList(input *payload.CreateListProductRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/products", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func setupTestCreateListWithImportFile(producerID uuid.UUID, orderRows int) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	// prepare csv data
	records := make([][]string, orderRows)
	for i := 0; i < orderRows; i++ {
		records[i] = []string{fake.Name(), fake.Car().Type, producerID.String()}
	}

	// create null csv file
	f, err := os.Create("test.csv")
	if err != nil {
		log.Fatalf("failed to create, error: %v", err)
	}
	defer f.Close()

	// write data to csv file
	wr := csv.NewWriter(f)
	err = wr.WriteAll(records)
	if err != nil {
		log.Fatalf("failed to write records, error: %v", err)
	}

	// create null request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// open csv file
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatalf("failed to open csv, error: %v", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", "test.csv")
	if err != nil {
		log.Fatalf("failed to create from file, error: %v", err)
	}

	// copy csv file to part
	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatalf("error copying file content: %v", err)
	}
	writer.Close()

	// remove test file
	err = os.Remove("test.csv")
	if err != nil {
		log.Fatalf("failed to remove csv file error: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/products", &requestBody)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func TestCreateListThirtyPart(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	accountID, _, err := SetUpForeignKeyData(db)
	assertion.NoError(err)

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			Username: fake.Name(),
			ID:       accountID,
			Email:    fake.Email(),
		},
	}

	// good case
	{
		// Arrange
		req := &payload.CreateListWithThirtyPartRequest{
			ThirtyParts: []payload.ThirtyPartRequire{
				{
					Url: "https://api.eazymock.net/mock/d60c7e6e-6819-4a59-b3ba-fc31fe0f019c/api/products",
				},
				{
					Url: "https://api.eazymock.net/mock/d60c7e6e-6819-4a59-b3ba-fc31fe0f019c/api/products/macbooks",
					Params: map[string]string{
						"product_type": "macbook",
					},
				},
			},
		}
		resp, ctx := setupCreateListWithThirtyPart(req)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)

		// Act
		err := r.CreateListWithThirtyPart(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.ListProductResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.NotNil(actual.Product)
	}
}

func setupCreateListWithThirtyPart(input *payload.CreateListWithThirtyPartRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/products/thirty-part", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
