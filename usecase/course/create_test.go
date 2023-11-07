package course

import (
	"context"
	"time"

	"bou.ke/monkey"
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestCreate() {
	assertion := s.Assertions
	testCourseID := "M0001"
	testSemesterID := "S0001"

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
		mockCourseRepo := course.NewMockRepository(s.T())
		mockSemesterRepo := semester.NewMockRepository(s.T())
		mockSemesterRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				EndTime: time.Now().Add(time.Hour * 1),
			}, nil,
		}
		mockCourseRepo.EXPECT().Create(s.ctx, &model.Course{
			ID:         testCourseID,
			SemesterID: testSemesterID,
			CreatedBy:  teq.Uint(1),
		}).ReturnArguments = mock.Arguments{nil}
		u := s.useCase(mockCourseRepo, mockSemesterRepo)

		// Act
		reply, err := u.CreateCourse(s.ctx, &payload.CreateCourseRequest{
			ID:         testCourseID,
			SemesterID: testSemesterID,
		})

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}

	// bad case
	{ // semester ended
		// Arrange
		mockSemesterRepo := semester.NewMockRepository(s.T())
		mockSemesterRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				EndTime: time.Now(),
			}, nil,
		}
		u := s.useCase(course.NewMockRepository(s.T()), mockSemesterRepo)

		// Act
		_, err := u.CreateCourse(s.ctx, &payload.CreateCourseRequest{
			ID:         testCourseID,
			SemesterID: testSemesterID,
		})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrCourseInvalidParam("semester ended")
		assertion.Equal(expected, err)
	}
}
