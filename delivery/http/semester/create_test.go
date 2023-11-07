package semester

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/test"
	"github.com/teq-quocbang/store/util/times"
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
		req := &payload.CreateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits,
			StartTime:         testStartTime,
			EndTime:           testEndTime,
			RegisterStartAt:   registerStartAt,
			RegisterExpiresAt: registerExpiresAt,
		}
		resp, ctx := setUpTestCreate(req)

		// Act
		err := r.CreateSemester(ctx)

		// Assert
		assertion.NoError(err)
		startTime, err := times.StringToTime(testStartTime)
		assertion.NoError(err)
		endTime, err := times.StringToTime(testEndTime)
		assertion.NoError(err)
		rStartAt, err := times.StringToTime(registerStartAt)
		assertion.NoError(err)
		rExpiresAt, err := times.StringToTime(registerExpiresAt)
		assertion.NoError(err)
		expected := &presenter.SemesterResponseWrapper{
			Semester: model.Semester{
				ID:                testSemesterID,
				MinCredits:        testMinCredits,
				StartTime:         startTime,
				EndTime:           endTime,
				RegisterStartAt:   rStartAt,
				RegisterExpiresAt: rExpiresAt,
			},
		}
		actual, err := test.UnmarshalBody[presenter.SemesterResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		assertion.Equal(expected.Semester.ID, actual.Semester.ID)
		assertion.Equal(expected.Semester.MinCredits, actual.Semester.MinCredits)
		assertion.Equal(expected.Semester.StartTime, actual.Semester.StartTime)
		assertion.Equal(expected.Semester.EndTime, actual.Semester.EndTime)
		assertion.Equal(expected.Semester.RegisterExpiresAt, actual.Semester.RegisterExpiresAt)
		assertion.Equal(expected.Semester.RegisterStartAt, actual.Semester.RegisterStartAt)
	}

	// bad case
	{
		// Arrange
		resp, ctx := setUpTestCreate(nil)

		// Act
		err := r.CreateSemester(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(400, resp.Code)
	}
}

func setUpTestCreate(input *payload.CreateSemesterRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/semester/", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
