package register

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
)

func TestUpdate(t *testing.T) {
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
	testThirstCourseID := "P0002"
	testThirstClassID := "CL0003"
	testUsername := "test_user"
	testEmail := "test@teqnological.asia"
	testPassword := "test_password"

	// good case
	{
		// Arrange
		db.TruncateTables()
		// arrange first data
		opts := []Option{
			WithCreateAccount(testUsername, testEmail, testPassword),
			WithCreateClass(testClassID, testCourseID, testSemesterID),
			WithCreateCourse(testCourseID, testSemesterID),
			WithCreateSemester(testSemesterID),
		}
		err := CreateForeignKeyDataHelper(db, opts...)
		assertion.NoError(err)

		// arrange second data
		opts = []Option{
			WithCreateClass(testSecondClassID, testSecondCourseID, testSemesterID),
			WithCreateCourse(testSecondCourseID, testSemesterID),
		}
		err = CreateForeignKeyDataHelper(db, opts...)
		assertion.NoError(err)

		// arrange thirst data
		opts = []Option{
			WithCreateClass(testThirstClassID, testThirstCourseID, testSemesterID),
			WithCreateCourse(testThirstCourseID, testSemesterID),
		}
		err = CreateForeignKeyDataHelper(db, opts...)
		assertion.NoError(err)

		// first register
		firstCreateRequest := &payload.CreateRegisterRequest{
			SemesterID: testSemesterID,
			ClassID:    testClassID,
			CourseID:   testCourseID,
		}
		resp, ctx := setUpTestCreate(firstCreateRequest)
		err = r.Create(ctx)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)

		// second register
		secondCreateRequest := &payload.CreateRegisterRequest{
			SemesterID: testSemesterID,
			ClassID:    testSecondClassID,
			CourseID:   testSecondCourseID,
		}
		resp, ctx = setUpTestCreate(secondCreateRequest)
		err = r.Create(ctx)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)

		// thirst register
		thirstCreateRequest := &payload.CreateRegisterRequest{
			SemesterID: testSemesterID,
			ClassID:    testThirstClassID,
			CourseID:   testThirstCourseID,
		}
		resp, ctx = setUpTestCreate(thirstCreateRequest)
		err = r.Create(ctx)
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)

		// arrange update unregister second data
		cancelRequest := &payload.UnRegisterRequest{
			SemesterID: testSemesterID,
			ClassID:    testSecondClassID,
			CourseID:   testSecondCourseID,
		}
		resp, ctx = setUpTestUpdate(cancelRequest)

		// Act
		err = r.Update(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{
		// Arrange
		resp, ctx := setUpTestUpdate(nil)

		// Act
		r.Update(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setUpTestUpdate(input *payload.UnRegisterRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/api/register/cancel", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
