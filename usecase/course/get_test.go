package course

import (
	"context"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestGet() {
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
		mockCourseRepo.EXPECT().GetByID(s.ctx, testCourseID).ReturnArguments = mock.Arguments{
			model.Course{
				ID:         "M0001",
				SemesterID: testSemesterID,
			}, nil,
		}
		u := s.useCase(mockCourseRepo, semester.NewMockRepository(s.T()))

		// Act
		reply, err := u.GetByID(s.ctx, testCourseID)

		// Assert
		assertion.NoError(err)
		assertion.Equal(testCourseID, reply.Course.ID)
	}

	// bad case
	{ // missing id
		// Arrange
		u := s.useCase(course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		_, err := u.GetByID(s.ctx, "")

		// Assert
		assertion.Error(err)
		expected := myerror.ErrCourseInvalidParam("id")
		assertion.Equal(expected, err)
	}
}
