package semester

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/times"
	"github.com/teq-quocbang/store/util/token"
)

func TestGet(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}
	testSemesterID := "TEST_S0001"
	testMinCredits := 15
	testStartTime := time.Now().Add(time.Minute * 2).Format(time.RFC3339)
	testEndTime := time.Now().Add(time.Nanosecond * times.ThreeMonth * 2).Format(time.RFC3339)
	registerStartAt := time.Now().Add(time.Minute * 2).Format(time.RFC3339)
	registerExpiresAt := time.Now().Add(time.Hour * 48).Format(time.RFC3339)
	userPrinciple := monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return &token.JWTClaimCustom{
			SessionID: uuid.New(),
			User: token.UserInfo{
				ID:       1,
				Username: "test_username",
				Email:    "test@teqnological.asia",
			},
		}
	})
	defer monkey.Unpatch(userPrinciple)

	// good case
	{
		// Arrange
		db.TruncateTables() // clean table before test
		// create semester
		createRequest := &payload.CreateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits,
			StartTime:         testStartTime,
			EndTime:           testEndTime,
			RegisterStartAt:   registerStartAt,
			RegisterExpiresAt: registerExpiresAt,
		}
		_, ctx := setUpTestCreate(createRequest)
		err := r.CreateSemester(ctx)
		assertion.NoError(err)

		resp, ctx := setUpTestGet(testSemesterID)

		// Act
		err = r.Get(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		db.TruncateTables() // clean table before test
		resp, ctx := setUpTestGet("")

		// Act
		r.Get(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setUpTestGet(id string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/semester/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	return rec, c
}
