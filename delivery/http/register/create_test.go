package register

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
	"github.com/teq-quocbang/store/delivery/http/account"
	"github.com/teq-quocbang/store/delivery/http/class"
	"github.com/teq-quocbang/store/delivery/http/course"
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

type semesterRequest struct {
	SemesterID string
}

type courseRequest struct {
	CourseID   string
	SemesterID string
}

type classRequest struct {
	ClassID    string
	SemesterID string
	CourseID   string
}

type accountRequest struct {
	Username string
	Email    string
	Password string
}

type options struct {
	CreateAccount  map[bool]accountRequest
	CreateSemester map[bool]semesterRequest
	CreateCourse   map[bool]courseRequest
	CreateClass    map[bool]classRequest
}

type Option func(*options)

func WithCreateSemester(semesterID string) Option {
	return func(o *options) {
		o.CreateSemester[true] = semesterRequest{
			SemesterID: semesterID,
		}
	}
}

func WithCreateCourse(courseID string, semesterID string) Option {
	return func(o *options) {
		o.CreateCourse[true] = courseRequest{
			CourseID:   courseID,
			SemesterID: semesterID,
		}
	}
}

func WithCreateClass(classID string, courseID string, semesterID string) Option {
	return func(o *options) {
		o.CreateClass[true] = classRequest{
			ClassID:    classID,
			CourseID:   courseID,
			SemesterID: semesterID,
		}
	}
}

func WithCreateAccount(username string, email string, password string) Option {
	return func(o *options) {
		o.CreateAccount[true] = accountRequest{
			Username: username,
			Email:    email,
			Password: password,
		}
	}
}

func TestCreate(t *testing.T) {
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

	// good case
	{
		// Arrange
		db.TruncateTables()
		opts := []Option{
			WithCreateAccount("test_user", "test@teqnological.asia", "test_password"),
			WithCreateSemester(testSemesterID),
			WithCreateCourse(testCourseID, testSemesterID),
			WithCreateClass(testClassID, testCourseID, testSemesterID),
		}
		err := CreateForeignKeyDataHelper(db, opts...)
		assertion.NoError(err)

		req := &payload.CreateRegisterRequest{
			SemesterID: testSemesterID,
			CourseID:   testCourseID,
			ClassID:    testClassID,
		}
		resp, ctx := setUpTestCreate(req)

		// Act
		err = r.Create(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.RegisterResponseCustom](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.NotNil(actual)
	}

	// bad case
	{
		// Arrange
		db.TruncateTables()
		resp, ctx := setUpTestCreate(nil)

		// Act
		r.Create(ctx)

		// Assert
		assertion.Equal(400, resp.Code)
	}
}

func setUpTestCreate(input *payload.CreateRegisterRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func CreateForeignKeyDataHelper(db *database.Database, option ...Option) error {
	opts := &options{
		CreateAccount:  make(map[bool]accountRequest, 1),
		CreateSemester: make(map[bool]semesterRequest, 1),
		CreateCourse:   make(map[bool]courseRequest, 1),
		CreateClass:    make(map[bool]classRequest, 1),
	}
	for _, o := range option {
		o(opts)
	}

	repo := repository.New(db.GetClient)

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

	if semester, ok := opts.CreateSemester[true]; ok {
		err := createSemester(semester.SemesterID, testMinCredits, testStartTime, testEndTime, registerStartAt, registerExpiresAt, repo)
		if err != nil {
			return err
		}
	}

	if course, ok := opts.CreateCourse[true]; ok {
		err := createCourse(course.CourseID, course.SemesterID, repo)
		if err != nil {
			return err
		}
	}

	if class, ok := opts.CreateClass[true]; ok {
		err := createClass(class.ClassID, class.CourseID, class.SemesterID, repo)
		if err != nil {
			return err
		}
	}

	if account, ok := opts.CreateAccount[true]; ok {
		err := createAccount(account.Username, account.Email, account.Password, repo)
		if err != nil {
			return err
		}
	}

	return nil
}

func createSemester(
	semesterID string,
	minCredits int,
	startTime string,
	endTime string,
	registerStartAt string,
	registerExpiresAt string,
	repo *repository.Repository) error {
	rSemester := semester.Route{
		UseCase: usecase.New(repo, nil),
	}

	// create semester
	createSemesterRequest := &payload.CreateSemesterRequest{
		ID:                semesterID,
		MinCredits:        minCredits,
		StartTime:         startTime,
		EndTime:           endTime,
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
	return nil
}

func createCourse(courseID string, semesterID string, repo *repository.Repository) error {
	rCourse := course.Route{
		UseCase: usecase.New(repo, nil),
	}
	// create course
	createCourseRequest := &payload.CreateCourseRequest{
		ID:         courseID,
		SemesterID: semesterID,
	}
	resp, ctx := setUpTestCreateCourse(createCourseRequest)
	err := rCourse.CreateCourse(ctx)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return fmt.Errorf("failed to create course, error: %v", resp.Body)
	}
	return nil
}

func createClass(classID string, courseID string, semesterID string, repo *repository.Repository) error {
	rClass := class.Route{
		UseCase: usecase.New(repo, nil),
	}
	// create class
	classStartTime := time.Now().Add(time.Minute * 2).Format(time.RFC3339)
	classEndTime := time.Now().Add(time.Hour * 1).Format(time.RFC3339)
	createClassRequest := &payload.CreateClassRequest{
		ID:         classID,
		CourseID:   courseID,
		SemesterID: semesterID,
		StartTime:  classStartTime,
		EndTime:    classEndTime,
		MaxSlot:    40,
		Credits:    5,
	}
	resp, ctx := setUpTestCreateClass(createClassRequest)
	err := rClass.CreateClass(ctx)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return fmt.Errorf("failed to create course, error: %v", resp.Body)
	}
	return nil
}

func createAccount(username string, email string, password string, repo *repository.Repository) error {
	rAccount := account.Route{
		UseCase: usecase.New(repo, nil),
	}

	// create account
	createAccountRequest := &payload.SignUpRequest{
		Username: username,
		Email:    email,
		Password: password,
	}
	resp, ctx := setUpTestSignUp(createAccountRequest)
	err := rAccount.SignUp(ctx)
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return fmt.Errorf("failed to sign up, error: %v", resp.Body)
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

func setUpTestCreateClass(input *payload.CreateClassRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/class", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}

func setUpTestSignUp(input *payload.SignUpRequest) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/user/sign-up", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	return rec, c
}
