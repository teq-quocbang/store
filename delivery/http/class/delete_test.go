package class

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
)

func TestDelete(t *testing.T) {
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

		createRequest := &payload.CreateClassRequest{
			ID:         testClassID,
			SemesterID: testSemesterID,
			CourseID:   testCourseID,
			StartTime:  testStartTime,
			EndTime:    testEndTime,
			Credits:    5,
			MaxSlot:    40,
		}
		resp, ctx := setUpTestCreate(createRequest)
		err = r.CreateClass(ctx)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)

		resp, ctx = setUpTestDelete(testClassID)

		// Act
		err = r.Delete(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		db.TruncateTables()
		resp, ctx := setUpTestDelete("")

		// Act
		r.Delete(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setUpTestDelete(id string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/class/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	return rec, c
}
