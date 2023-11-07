package class

import (
	"context"
	"time"

	"bou.ke/monkey"
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository/class"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestCreate() {
	assertion := s.Assertions

	testSemesterID := "S0001"
	testCourseID := "M0001"
	testClassID := "CL0001"
	testCredits := 5
	testMaxSlot := 40
	testStartTime, err := times.StringToTime(time.Now().Format(time.RFC3339))
	assertion.NoError(err)
	testEndTime, err := times.StringToTime(time.Now().Add(times.ThreeMonth * time.Nanosecond * 2).Format(time.RFC3339))
	assertion.NoError(err)
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
		mockSemesterRepo := semester.NewMockRepository(s.T())
		mockCourseRepo := course.NewMockRepository(s.T())
		mockClassRepo := class.NewMockRepository(s.T())
		mockSemesterRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				StartTime: testStartTime,
				EndTime:   testEndTime,
			}, nil,
		}
		mockCourseRepo.EXPECT().GetByID(s.ctx, testCourseID).ReturnArguments = mock.Arguments{
			model.Course{}, nil,
		}
		testClassStartTime, err := times.StringToTime(testClassStartTimeFormat)
		assertion.NoError(err)
		testClassEndTime, err := times.StringToTime(testClassEndTimeFormat)
		assertion.NoError(err)
		createClassModel := &model.Class{
			ID:         testClassID,
			CourseID:   testCourseID,
			SemesterID: testSemesterID,
			StartTime:  testClassStartTime,
			EndTime:    testClassEndTime,
			Credits:    uint(testCredits),
			MaxSlot:    uint(testMaxSlot),
			CreatedBy:  teq.Uint(1),
		}
		mockClassRepo.EXPECT().Create(s.ctx, createClassModel).ReturnArguments = mock.Arguments{nil}

		u := s.useCase(mockClassRepo, mockCourseRepo, mockSemesterRepo)

		// Act
		reply, err := u.CreateClass(s.ctx, &payload.CreateClassRequest{
			ID:         testClassID,
			SemesterID: testSemesterID,
			CourseID:   testCourseID,
			StartTime:  testClassStartTimeFormat,
			EndTime:    testClassEndTimeFormat,
			Credits:    testCredits,
			MaxSlot:    testMaxSlot,
		})

		// Assert
		assertion.NoError(err)
		expected := &presenter.ClassResponseWrapper{
			Class: model.Class{
				ID:         testClassID,
				CourseID:   testCourseID,
				SemesterID: testSemesterID,
				StartTime:  testClassStartTime,
				EndTime:    testClassEndTime,
				Credits:    uint(testCredits),
				MaxSlot:    uint(testMaxSlot),
			},
		}
		assertion.Equal(expected.Class.ID, reply.Class.ID)
		assertion.Equal(expected.Class.CourseID, reply.Class.CourseID)
		assertion.Equal(expected.Class.SemesterID, reply.Class.SemesterID)
		assertion.Equal(expected.Class.StartTime, reply.Class.StartTime)
		assertion.Equal(expected.Class.EndTime, reply.Class.EndTime)
		assertion.Equal(expected.Class.Credits, reply.Class.Credits)
		assertion.Equal(expected.Class.MaxSlot, reply.Class.MaxSlot)
	}

	// bad case
	{ // invalid param missing id
		// Arrange
		u := s.useCase(class.NewMockRepository(s.T()), course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		_, err := u.CreateClass(s.ctx, &payload.CreateClassRequest{
			CourseID:   testCourseID,
			SemesterID: testSemesterID,
			StartTime:  testClassStartTimeFormat,
			EndTime:    testClassEndTimeFormat,
			Credits:    testCredits,
			MaxSlot:    testMaxSlot,
		})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrClassInvalidParam("Key: 'CreateClassRequest.ID' Error:Field validation for 'ID' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
	{ // semester not found
		// Arrange
		mockSemesterRepo := semester.NewMockRepository(s.T())
		mockSemesterRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{}, gorm.ErrRecordNotFound,
		}

		u := s.useCase(class.NewMockRepository(s.T()), course.NewMockRepository(s.T()), mockSemesterRepo)

		// Act
		_, err := u.CreateClass(s.ctx, &payload.CreateClassRequest{
			ID:         testClassID,
			CourseID:   testCourseID,
			SemesterID: testSemesterID,
			StartTime:  testClassStartTimeFormat,
			EndTime:    testClassEndTimeFormat,
			Credits:    testCredits,
			MaxSlot:    testMaxSlot,
		})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrSemesterGet(gorm.ErrRecordNotFound)
		assertion.Equal(expected, err)
	}
	{ // course not found
		// Arrange
		mockSemesterRepo := semester.NewMockRepository(s.T())
		mockCourseRepo := course.NewMockRepository(s.T())
		mockSemesterRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				StartTime: testStartTime,
				EndTime:   testEndTime,
			}, nil,
		}
		mockCourseRepo.EXPECT().GetByID(s.ctx, testCourseID).ReturnArguments = mock.Arguments{
			model.Course{}, gorm.ErrRecordNotFound,
		}

		u := s.useCase(class.NewMockRepository(s.T()), mockCourseRepo, mockSemesterRepo)

		// Act
		_, err := u.CreateClass(s.ctx, &payload.CreateClassRequest{
			ID:         testClassID,
			CourseID:   testCourseID,
			SemesterID: testSemesterID,
			StartTime:  testClassStartTimeFormat,
			EndTime:    testClassEndTimeFormat,
			Credits:    testCredits,
			MaxSlot:    testMaxSlot,
		})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrCourseGet(gorm.ErrRecordNotFound)
		assertion.Equal(expected, err)
	}
}
