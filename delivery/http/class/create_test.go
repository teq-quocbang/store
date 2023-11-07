package class

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/delivery/http/course"
	"github.com/teq-quocbang/store/delivery/http/semester"
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

	testSemesterID := "S0001"
	testCourseID := "M0001"
	testClassID := "CL0001"
	testStartTime := time.Now().Add(time.Minute * 2).Format(time.RFC3339)
	testEndTime := time.Now().Add(time.Hour * 1).Format(time.RFC3339)

	// good case
	{
		// Arrange
		db.TruncateTables()
		err := CreateForeignKeyDataHelper(testSemesterID, testCourseID, db)
		assertion.NoError(err)

		req := &payload.CreateClassRequest{
			ID:         testClassID,
			SemesterID: testSemesterID,
			CourseID:   testCourseID,
			StartTime:  testStartTime,
			EndTime:    testEndTime,
			Credits:    5,
			MaxSlot:    40,
		}
		resp, ctx := setUpTestCreate(req)

		// Act
		err = r.CreateClass(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.ClassResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		start, err := times.StringToTime(req.StartTime)
		assertion.NoError(err)
		end, err := times.StringToTime(req.EndTime)
		assertion.NoError(err)
		expected := &presenter.ClassResponseWrapper{
			Class: model.Class{
				ID:         testClassID,
				CourseID:   testCourseID,
				SemesterID: testSemesterID,
				MaxSlot:    40,
				Credits:    5,
				StartTime:  start,
				EndTime:    end,
			},
		}
		assertion.Equal(expected.Class.ID, actual.Class.ID)
		assertion.Equal(expected.Class.SemesterID, actual.Class.SemesterID)
		assertion.Equal(expected.Class.CourseID, actual.Class.CourseID)
		assertion.Equal(expected.Class.StartTime, actual.Class.StartTime)
		assertion.Equal(expected.Class.EndTime, actual.Class.EndTime)
		assertion.Equal(expected.Class.MaxSlot, actual.Class.MaxSlot)
		assertion.Equal(expected.Class.Credits, actual.Class.Credits)
	}

	// bad case
	{
		// Arrange
		resp, ctx := setUpTestCreate(nil)

		// Act
		r.CreateClass(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func CreateForeignKeyDataHelper(semesterID string, courseID string, db *database.Database) error {
	repo := repository.New(db.GetClient)
	rSemester := semester.Route{
		UseCase: usecase.New(repo, nil),
	}
	rCourse := course.Route{
		UseCase: usecase.New(repo, nil),
	}

	testMinCredits := 15
	testStartTime := time.Now().Add(time.Second * 5).Format(time.RFC3339)
	testEndTime := time.Now().Add(time.Nanosecond * times.ThreeMonth * 2).Format(time.RFC3339)
	registerStartAt := time.Now().Add(time.Second * 5).Format(time.RFC3339)
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
		ID:                semesterID,
		MinCredits:        testMinCredits,
		StartTime:         testStartTime,
		EndTime:           testEndTime,
		RegisterStartAt:   registerStartAt,
		RegisterExpiresAt: registerExpiresAt,
	}
	resp, ctx := setUpTestCreateSemester(createSemesterRequest)
	err := rSemester.CreateSemester(ctx)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return fmt.Errorf("failed to create semester, error: %v", resp.Body)
	}

	req := &payload.CreateCourseRequest{
		ID:         courseID,
		SemesterID: semesterID,
	}
	resp, ctx = setUpTestCreateCourse(req)
	err = rCourse.CreateCourse(ctx)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return fmt.Errorf("failed to create course, error: %v", resp.Body)
	}

	return nil
}

func setUpTestCreateCourse(input *payload.CreateCourseRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/course", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func setUpTestCreateSemester(input *payload.CreateSemesterRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/semester", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func setUpTestCreate(input *payload.CreateClassRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/class", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
