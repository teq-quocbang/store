package semester

import (
	"fmt"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
)

func (s *TestSuite) TestGetList() {
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
		req := &payload.GetListSemesterRequest{
			Year: "2023",
		}
		getListResponse := []model.Semester{
			{
				ID:                testSemesterID,
				MinCredits:        testMinCredits,
				StartTime:         testStartTime,
				EndTime:           testEndTime,
				RegisterStartAt:   registerStartAt,
				RegisterExpiresAt: registerExpiresAt,
			},
			{
				ID:                fmt.Sprintf("%s_%s", testSemesterID, "1"), // TEST_S0001_1
				MinCredits:        testMinCredits,
				StartTime:         testEndTime.Add(time.Hour * 1),
				EndTime:           testEndTime.Add(time.Nanosecond * times.ThreeMonth * 2),
				RegisterStartAt:   registerStartAt,
				RegisterExpiresAt: registerExpiresAt,
			},
		}
		mockRepo.EXPECT().GetListByYear(s.ctx, "2023").ReturnArguments = mock.Arguments{
			getListResponse, nil,
		}
		u := s.useCase(mockRepo)

		// Act
		reply, err := u.GetList(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.Equal(2, len(reply.Semester))
	}

	// bad case
	{ // missing year
		// Arrange
		u := s.useCase(semester.NewMockRepository(s.T()))

		// Act
		_, err := u.GetList(s.ctx, &payload.GetListSemesterRequest{})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrSemesterInvalidParam("Key: 'GetListSemesterRequest.Year' Error:Field validation for 'Year' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
