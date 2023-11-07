package class

import (
	"context"
	"time"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/repository/class"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestDelete() {
	assertion := s.Assertions

	testClassID := "CL0001"
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

	testClassStartTimeFormat := time.Now().Add(time.Minute * 2).Format(time.RFC3339)
	testClassEndTimeFormat := time.Now().Add(time.Hour * 2).Format(time.RFC3339)
	// good case
	{
		// Arrange
		mockClassRepo := class.NewMockRepository(s.T())
		testClassStartTime, err := times.StringToTime(testClassStartTimeFormat)
		assertion.NoError(err)
		testClassEndTime, err := times.StringToTime(testClassEndTimeFormat)
		assertion.NoError(err)
		mockClassRepo.EXPECT().GetByID(s.ctx, testClassID).ReturnArguments = mock.Arguments{
			model.Class{
				StartTime: testClassStartTime,
				EndTime:   testClassEndTime,
			}, nil,
		}
		mockClassRepo.EXPECT().Delete(s.ctx, testClassID).ReturnArguments = mock.Arguments{nil}

		u := s.useCase(mockClassRepo, course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		err = u.Delete(s.ctx, testClassID)

		// Assert
		assertion.NoError(err)
	}

	// bad case
	{ // invalid param missing id
		// Arrange
		u := s.useCase(class.NewMockRepository(s.T()), course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		err := u.Delete(s.ctx, "")

		// Assert
		assertion.Error(err)
		expected := myerror.ErrClassInvalidParam("id")
		assertion.Equal(expected, err)
	}
}
