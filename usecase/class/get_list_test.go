package class

import (
	"context"
	"time"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/class"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestGetList() {
	assertion := s.Assertions

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
		mockClassRepo.EXPECT().GetListBySemester(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			[]model.Class{
				{
					ID:        "CL0001",
					StartTime: testClassStartTime,
					EndTime:   testClassEndTime,
				},
				{
					ID:        "CL0002",
					StartTime: testClassStartTime,
					EndTime:   testClassEndTime,
				},
			}, nil,
		}

		u := s.useCase(mockClassRepo, course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		reply, err := u.GetList(s.ctx, &payload.ListClassBySemesterRequest{
			SemesterID: testSemesterID,
		})

		// Assert
		assertion.NoError(err)
		assertion.Equal(2, len(reply.Class))
	}

	// bad case
	{ // invalid param missing id
		// Arrange
		u := s.useCase(class.NewMockRepository(s.T()), course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		_, err := u.GetList(s.ctx, &payload.ListClassBySemesterRequest{})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrClassInvalidParam("Key: 'ListClassBySemesterRequest.SemesterID' Error:Field validation for 'SemesterID' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
