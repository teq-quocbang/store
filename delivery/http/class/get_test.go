package class

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/test"
)

func TestGet(t *testing.T) {
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
		// create class
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

		resp, ctx = setUpTestGet(testClassID)

		// Act
		err = r.Get(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.ClassResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		expected := presenter.ClassResponseWrapper{
			Class: model.Class{
				ID:         testClassID,
				SemesterID: testSemesterID,
				CourseID:   testCourseID,
			},
		}
		assertion.Equal(expected.Class.ID, actual.Class.ID)
		assertion.Equal(expected.Class.SemesterID, actual.Class.SemesterID)
		assertion.Equal(expected.Class.CourseID, actual.Class.CourseID)
	}

	// bad case
	{
		// Arrange
		resp, ctx := setUpTestGet("")

		// Act
		r.Get(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setUpTestGet(id string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/class/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	return rec, c
}
