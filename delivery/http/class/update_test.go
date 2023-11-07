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
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/test"
)

func TestUpdate(t *testing.T) {
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

		req := &payload.UpdateClassRequest{
			MaxSlot: 20, // 40 -> 20
		}
		resp, ctx = setUpTestUpdate(req, testClassID)

		// Act
		err = r.Update(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.ClassResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		expected := presenter.ClassResponseWrapper{
			Class: model.Class{
				MaxSlot: 20,
			},
		}
		assertion.Equal(expected.Class.MaxSlot, actual.Class.MaxSlot)
	}
}

func setUpTestUpdate(input *payload.UpdateClassRequest, id string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/api/class/:id", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	return rec, c
}
