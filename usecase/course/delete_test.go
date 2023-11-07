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

func (s *TestSuite) TestDelete() {
	assertion := s.Assertions
	testCourseID := "M0001"

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
			model.Course{}, nil,
		}
		mockCourseRepo.EXPECT().Delete(s.ctx, testCourseID).ReturnArguments = mock.Arguments{nil}
		u := s.useCase(mockCourseRepo, semester.NewMockRepository(s.T()))

		// Act
		err := u.Delete(s.ctx, testCourseID)

		// Assert
		assertion.NoError(err)
	}

	// bad case
	{ // missing id
		// Arrange
		u := s.useCase(course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		err := u.Delete(s.ctx, "")

		// Assert
		assertion.Error(err)
		expected := myerror.ErrCourseInvalidParam("id")
		assertion.Equal(expected, err)
	}
}
