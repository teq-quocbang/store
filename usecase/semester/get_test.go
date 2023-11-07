package semester

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
)

func (s *TestSuite) TestGet() {
	assertion := s.Assertions
	testSemesterID := "TEST_S0001"
	testMinCredits := 15
	testStartTime := time.Now().Add(time.Minute * 2)
	testEndTime := time.Now().Add(time.Nanosecond * times.ThreeMonth * 2)
	registerStartAt := time.Now().Add(time.Minute * 2)
	registerExpiresAt := time.Now().Add(time.Hour * 48)

	// good case
	{
		// Arrange
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

		u := s.useCase(mockRepo)

		// Act
		reply, err := u.GetByID(s.ctx, testSemesterID)

		// Assert
		assertion.NoError(err)
		assertion.Equal(testSemesterID, reply.Semester.ID)
		assertion.Equal(testMinCredits, reply.Semester.MinCredits)
		assertion.Equal(testStartTime, reply.Semester.StartTime)
		assertion.Equal(testEndTime, reply.Semester.EndTime)
	}

	// bad case
	{ // missing id
		// Arrange
		u := s.useCase(semester.NewMockRepository(s.T()))

		// Act
		_, err := u.GetByID(s.ctx, "")

		// Assert
		assertion.Error(err)
		expected := myerror.ErrSemesterInvalidParam("id")
		assertion.Equal(expected, err)
	}
}
