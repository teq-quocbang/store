package class

import (
	"context"
	"time"

	"bou.ke/monkey"
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/class"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/times"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestUpdate() {
	assertion := s.Assertions

	testSemesterID := "S0001"
	testClassID := "CL0001"
	testCourseID := "M0001"
	testCredits := 5
	testMaxSlot := 40
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
				ID:        "CL0001",
				StartTime: testClassStartTime,
				EndTime:   testClassEndTime,
			}, nil,
		}
		mockClassRepo.EXPECT().Update(s.ctx, &model.Class{
			ID:         testClassID,
			CourseID:   testCourseID,
			SemesterID: testSemesterID,
			StartTime:  testClassStartTime,
			EndTime:    testClassEndTime,
			MaxSlot:    uint(testMaxSlot),
			Credits:    uint(testCredits),
			UpdatedBy:  teq.Uint(1),
		}).ReturnArguments = mock.Arguments{nil}

		u := s.useCase(mockClassRepo, course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		reply, err := u.Update(s.ctx, &payload.UpdateClassRequest{
			ID:         testClassID,
			CourseID:   testCourseID,
			SemesterID: testSemesterID,
			StartTime:  testClassStartTimeFormat,
			EndTime:    testClassEndTimeFormat,
			MaxSlot:    testMaxSlot,
			Credits:    testCredits,
		})

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}

	// *NOTE: no bad case
}
