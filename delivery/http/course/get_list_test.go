package course

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
	"github.com/teq-quocbang/store/delivery/http/semester"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/test"
	"github.com/teq-quocbang/store/util/times"
	"github.com/teq-quocbang/store/util/token"
)

func TestGetList(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	rSemester := semester.Route{
		UseCase: usecase.New(repo, nil),
	}

	// good case
	{
		// Arrange
		db.TruncateTables() // clean database
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
		createSemesterRequest := &payload.CreateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits,
			StartTime:         testStartTime,
			EndTime:           testEndTime,
			RegisterStartAt:   registerStartAt,
			RegisterExpiresAt: registerExpiresAt,
		}
		// create semester
		_, ctx := setUpTestCreateSemester(createSemesterRequest)
		err := rSemester.CreateSemester(ctx)
		assertion.NoError(err)

		// create courses
		createMCourse := &payload.CreateCourseRequest{
			ID:         "M0001",
			SemesterID: testSemesterID,
		}
		_, ctx = setUpTestCreate(createMCourse)
		err = r.CreateCourse(ctx)
		assertion.NoError(err)

		createBCourse := &payload.CreateCourseRequest{
			ID:         "B0001",
			SemesterID: testSemesterID,
		}
		_, ctx = setUpTestCreate(createBCourse)
		err = r.CreateCourse(ctx)
		assertion.NoError(err)

		req := &payload.ListCourseBySemesterRequest{
			SemesterID: testSemesterID,
		}
		resp, ctx := setUpTestGetList(req)

		// Act
		err = r.GetList(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.ListCourseResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.Equal(2, len(actual.Course))
	}

	// bad case
	{
		// Arrange
		resp, ctx := setUpTestGetList(nil)

		// Act
		r.GetList(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setUpTestGetList(input *payload.ListCourseBySemesterRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/course", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
