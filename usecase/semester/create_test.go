package semester

import (
	"context"
	"time"

	"bou.ke/monkey"
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestCreate() {
	assertion := s.Assertions
	testSemesterID := "TEST_S0001"
	testMinCredits := 15
	testStartTime := time.Now().Add(time.Minute * 2).Format(time.RFC3339)
	testEndTime := time.Now().Add(time.Nanosecond * times.ThreeMonth * 2).Format(time.RFC3339)
	registerStartAt := time.Now().Add(time.Minute * 2).Format(time.RFC3339)
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

	// good case
	{
		// Arrange
		req := &payload.CreateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits,
			StartTime:         testStartTime,
			EndTime:           testEndTime,
			RegisterStartAt:   registerStartAt,
			RegisterExpiresAt: registerExpiresAt,
		}
		mockRepo := semester.NewMockRepository(s.T())
		startTime, err := times.StringToTime(req.StartTime)
		assertion.NoError(err)
		endTime, err := times.StringToTime(req.EndTime)
		assertion.NoError(err)
		rStartAt, err := times.StringToTime(req.RegisterStartAt)
		assertion.NoError(err)
		rExpiresAt, err := times.StringToTime(req.RegisterExpiresAt)
		assertion.NoError(err)
		createRequest := &model.Semester{
			ID:                req.ID,
			MinCredits:        req.MinCredits,
			CreatedBy:         teq.Uint(1),
			StartTime:         startTime,
			EndTime:           endTime,
			RegisterStartAt:   rStartAt,
			RegisterExpiresAt: rExpiresAt,
		}
		mockRepo.EXPECT().Create(s.ctx, createRequest).ReturnArguments = mock.Arguments{
			nil,
		}
		u := s.useCase(mockRepo)

		// Act
		reply, err := u.CreateSemester(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}

	// bad case
	{ // invalid param
		// Arrange
		u := s.useCase(semester.NewMockRepository(s.T()))

		// Act
		_, err := u.CreateSemester(s.ctx, &payload.CreateSemesterRequest{})

		// Assert
		assertion.Error(err)
	}
	{ // a semester at least 3 month
		// Arrange
		req := &payload.CreateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits,
			StartTime:         testStartTime,
			EndTime:           time.Now().Add(time.Hour * 1).Format(time.RFC3339),
			RegisterStartAt:   registerStartAt,
			RegisterExpiresAt: registerExpiresAt,
		}
		u := s.useCase(semester.NewMockRepository(s.T()))

		// Act
		_, err := u.CreateSemester(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrSemesterInvalidParam("a semester at least 3 month")
		assertion.Equal(expected, err)
	}
	{ // a semester maximum is six month
		// Arrange
		req := &payload.CreateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits,
			StartTime:         testStartTime,
			EndTime:           time.Now().Add(time.Nanosecond * times.SixMonth * 2).Format(time.RFC3339),
			RegisterStartAt:   registerStartAt,
			RegisterExpiresAt: registerExpiresAt,
		}
		u := s.useCase(semester.NewMockRepository(s.T()))

		// Act
		_, err := u.CreateSemester(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrSemesterInvalidParam("a semester maximum is six month")
		assertion.Equal(expected, err)
	}
}
