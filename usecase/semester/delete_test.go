package semester

import (
	"context"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestDelete() {
	assertion := s.Assertions
	testSemesterID := "TEST_S0001"
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
		mockRepo := semester.NewMockRepository(s.T())
		mockRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{}, nil,
		}
		mockRepo.EXPECT().Delete(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			nil,
		}
		u := s.useCase(mockRepo)

		// Act
		err := u.Delete(s.ctx, testSemesterID)

		// Assert
		assertion.NoError(err)
	}

	// bad case
	{ // missing id
		// Arrange
		u := s.useCase(semester.NewMockRepository(s.T()))

		// Act
		err := u.Delete(s.ctx, "")

		// Assert
		assertion.Error(err)
		expected := myerror.ErrSemesterInvalidParam("id")
		assertion.Equal(expected, err)
	}
}
