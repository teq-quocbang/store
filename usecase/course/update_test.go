package course

import (
	"context"
	"time"

	"bou.ke/monkey"
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestUpdate() {
	assertion := s.Assertions
	testCourseID := "M0001"
	changedSemesterID := "S0002"

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
		mockSemesterRepo := semester.NewMockRepository(s.T())
		mockSemesterRepo.EXPECT().GetByID(s.ctx, changedSemesterID).ReturnArguments = mock.Arguments{model.Semester{EndTime: time.Now().Add(time.Hour * 1)}, nil}
		mockCourseRepo.EXPECT().Update(s.ctx, &model.Course{
			ID:         testCourseID,
			SemesterID: changedSemesterID,
			UpdatedBy:  teq.Uint(1),
		}).ReturnArguments = mock.Arguments{
			nil,
		}
		u := s.useCase(mockCourseRepo, mockSemesterRepo)

		// Act
		reply, err := u.Update(s.ctx, &payload.UpdateCourseRequest{
			ID:         testCourseID,
			SemesterID: changedSemesterID,
		})

		// Assert
		assertion.NoError(err)
		assertion.Equal(testCourseID, reply.Course.ID)
		assertion.Equal(changedSemesterID, reply.Course.SemesterID)
	}

	// bad case
	{ // missing id
		// Arrange
		u := s.useCase(course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		_, err := u.Update(s.ctx, &payload.UpdateCourseRequest{})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrCourseInvalidParam("Key: 'UpdateCourseRequest.ID' Error:Field validation for 'ID' failed on the 'required' tag\nKey: 'UpdateCourseRequest.SemesterID' Error:Field validation for 'SemesterID' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
