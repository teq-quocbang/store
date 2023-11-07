package register

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/test"
)

func TestGetListHistories(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	cache := database.InitCache()
	r := Route{
		UseCase: usecase.New(repo, cache),
	}

	testSemesterID := "S0001"
	testCourseID := "M0001"
	testClassID := "CL0001"
	testSecondCourseID := "B0002"
	testSecondClassID := "CL0002"

	// good case
	{
		// Arrange
		db.TruncateTables()
		opts := []Option{
			WithCreateAccount("test_username", "test@teqnological.asia", "test_password"),
			WithCreateSemester(testSemesterID),
			WithCreateCourse(testCourseID, testSemesterID),
			WithCreateClass(testClassID, testCourseID, testSemesterID),
		}
		err := CreateForeignKeyDataHelper(db, opts...)
		assertion.NoError(err)
		opts = []Option{
			WithCreateCourse(testSecondCourseID, testSemesterID),
			WithCreateClass(testSecondClassID, testSecondCourseID, testSemesterID),
		}
		err = CreateForeignKeyDataHelper(db, opts...)
		if err != nil {
			if !errors.Is(err, myerror.ErrAccountConflictUniqueConstraint("Username or Email was registered")) {
				assertion.NoError(err)
			}
		}

		// first register
		firstRegisterRequest := &payload.CreateRegisterRequest{
			SemesterID: testSemesterID,
			ClassID:    testClassID,
			CourseID:   testCourseID,
		}
		resp, ctx := setUpTestCreate(firstRegisterRequest)
		err = r.Create(ctx)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)

		// second register
		secondRegisterRequest := &payload.CreateRegisterRequest{
			SemesterID: testSemesterID,
			ClassID:    testSecondClassID,
			CourseID:   testSecondCourseID,
		}
		resp, ctx = setUpTestCreate(secondRegisterRequest)
		err = r.Create(ctx)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)

		req := &payload.ListRegisteredHistories{
			SemesterID: testSemesterID,
		}
		resp, ctx = setUpTestGetHistories(req)

		// Act
		err = r.GetHistories(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.ListRegisterResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.Equal(2, len(actual.Register))
	}
	// *NOTE: no bad case
}

func setUpTestGetHistories(input *payload.ListRegisteredHistories) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodGet, "/api/register", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("")

	return rec, c
}
