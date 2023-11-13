package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestExport(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	// good case
	{
		// Arrange
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

		_, ctx := setupTestCreateListWithImportFile(producerID, 20)
		ctx.Set(string(auth.UserPrincipleKey), userPrinciple)
		err = r.CreateListWithImportFile(ctx)
		assertion.NoError(err)

		req := &payload.ExportProductRequest{
			FileExtension: "csv",
		}
		resp, ctx := setupExportTest(req)

		// Act
		err = r.Export(ctx)

		// Assert
		assertion.NoError(err)
		actual, err := test.UnmarshalBody[*presenter.ListProductResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.NotNil(actual.Meta)
	}
}

func setupExportTest(input *payload.ExportProductRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/product/export", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
