package class

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/test"
)

func TestGetList(t *testing.T) {
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
	testSecondClassID := "CL0002"

	// good case
	{
		// Arrange
		err := CreateForeignKeyDataHelper(testSemesterID, testCourseID, db)
		assertion.NoError(err)
		// create first class
		createFirstClass := &payload.CreateClassRequest{
			ID:         testClassID,
			SemesterID: testSemesterID,
			CourseID:   testCourseID,
			Credits:    5,
			MaxSlot:    40,
			StartTime:  testStartTime,
			EndTime:    testEndTime,
		}
		resp, ctx := setUpTestCreate(createFirstClass)
		err = r.CreateClass(ctx)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)

		// create second class
		createSecondClass := &payload.CreateClassRequest{
			ID:         testSecondClassID,
			SemesterID: testSemesterID,
			CourseID:   testCourseID,
			Credits:    5,
			MaxSlot:    40,
			StartTime:  testStartTime,
			EndTime:    testEndTime,
		}
		resp, ctx = setUpTestCreate(createSecondClass)
		err = r.CreateClass(ctx)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)

		req := &payload.ListClassBySemesterRequest{
			SemesterID: testSemesterID,
		}
		resp, ctx = setUpTestGetList(req)

		// Act
		err = r.GetList(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.ListClassResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.Equal(2, len(actual.Class))
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

func setUpTestGetList(input *payload.ListClassBySemesterRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodGet, "/api/class", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
