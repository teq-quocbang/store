package course

import (
	"context"

	"bou.ke/monkey"
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

	// good case
	{
		// Arrange
		mockCourseRepo := course.NewMockRepository(s.T())
		mockCourseRepo.EXPECT().GetListBySemester(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			[]model.Course{
				{
					ID:         "M0001",
					SemesterID: testSemesterID,
				},
				{
					ID:         "B0001",
					SemesterID: testSemesterID,
				},
			}, nil,
		}
		u := s.useCase(mockCourseRepo, semester.NewMockRepository(s.T()))

		// Act
		replies, err := u.GetList(s.ctx, &payload.ListCourseBySemesterRequest{
			SemesterID: testSemesterID,
		})

		// Assert
		assertion.NoError(err)
		assertion.Equal(2, len(replies.Course))
	}

	// bad case
	{ // missing id
		// Arrange
		u := s.useCase(course.NewMockRepository(s.T()), semester.NewMockRepository(s.T()))

		// Act
		_, err := u.GetList(s.ctx, &payload.ListCourseBySemesterRequest{})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrCourseInvalidParam("Key: 'ListCourseBySemesterRequest.SemesterID' Error:Field validation for 'SemesterID' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
