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
	"gorm.io/gorm"
)

func (s *TestSuite) TestUpdate() {
	assertion := s.Assertions
	testSemesterID := "TEST_S0001"
	testMinCredits := 15
	testStartTime := time.Now().Add(time.Minute * 2)
	testEndTime := time.Now().Add(time.Nanosecond * times.ThreeMonth * 2)
	registerStartAt := time.Now().Add(time.Minute * 2)
	registerExpiresAt := time.Now().Add(time.Hour * 48)
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
		req := &payload.UpdateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits - 5, // 15 - 5
			StartTime:         testStartTime.Format(time.RFC3339),
			EndTime:           testEndTime.Format(time.RFC3339),
			RegisterStartAt:   registerStartAt.Format(time.RFC3339),
			RegisterExpiresAt: registerExpiresAt.Format(time.RFC3339),
		}
		mockRepo := semester.NewMockRepository(s.T())
		mockRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				ID:                testSemesterID,
				MinCredits:        testMinCredits,
				StartTime:         testStartTime,
				EndTime:           testEndTime,
				RegisterStartAt:   registerStartAt,
				RegisterExpiresAt: registerExpiresAt,
			}, nil,
		}
		start, err := times.StringToTime(testStartTime.Format(time.RFC3339))
		assertion.NoError(err)
		end, err := times.StringToTime(testEndTime.Format(time.RFC3339))
		assertion.NoError(err)
		rStartAt, err := times.StringToTime(registerStartAt.Format(time.RFC3339))
		assertion.NoError(err)
		rExpires, err := times.StringToTime(registerExpiresAt.Format(time.RFC3339))
		assertion.NoError(err)
		modelMockUpdate := &model.Semester{
			ID:                testSemesterID,
			MinCredits:        testMinCredits - 5, // 15 - 5
			StartTime:         start,
			EndTime:           end,
			RegisterStartAt:   rStartAt,
			RegisterExpiresAt: rExpires,
			UpdatedBy:         teq.Uint(1),
		}
		mockRepo.EXPECT().Update(s.ctx, modelMockUpdate).ReturnArguments = mock.Arguments{
			nil,
		}
		u := s.useCase(mockRepo)

		// Act
		reply, err := u.Update(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.Equal(testSemesterID, reply.Semester.ID)
		assertion.Equal(10, reply.Semester.MinCredits)
	}

	// bad case
	{ // semester does not exist
		// Arrange
		req := &payload.UpdateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits - 5, // 15 - 5
			StartTime:         testStartTime.Format(time.RFC3339),
			EndTime:           testEndTime.Format(time.RFC3339),
			RegisterStartAt:   registerStartAt.Format(time.RFC3339),
			RegisterExpiresAt: registerExpiresAt.Format(time.RFC3339),
		}
		mockRepo := semester.NewMockRepository(s.T())
		mockRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{}, gorm.ErrRecordNotFound,
		}
		u := s.useCase(mockRepo)

		// Act
		_, err := u.Update(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrSemesterGet(gorm.ErrRecordNotFound)
		assertion.Equal(expected, err)
	}
	{ // semester ended
		// Arrange
		req := &payload.UpdateSemesterRequest{
			ID:                testSemesterID,
			MinCredits:        testMinCredits - 5, // 15 - 5
			StartTime:         testStartTime.Format(time.RFC3339),
			EndTime:           testEndTime.Format(time.RFC3339),
			RegisterStartAt:   registerStartAt.Format(time.RFC3339),
			RegisterExpiresAt: registerExpiresAt.Format(time.RFC3339),
		}

		mockRepo := semester.NewMockRepository(s.T())
		timeEnded, err := times.StringToTime(time.Now().Format(time.RFC3339))
		assertion.NoError(err)
		mockRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				EndTime: timeEnded,
			}, nil,
		}
		u := s.useCase(mockRepo)

		// Act
		_, err = u.Update(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrSemesterInvalidParam("the semester was ended")
		assertion.Equal(expected, err)
	}
}
