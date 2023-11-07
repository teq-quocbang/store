package register

import (
	"context"
	"fmt"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/cache"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/class"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/register"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestUnRegister() {
	assertion := s.Assertions
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
	testSemesterID := "S0001"
	testCourseID := "M0001"
	testClassID := "CL0001"

	// good case
	{
		// Arrange
		mockRegisterRepo := register.NewMockRepository(s.T())
		mockCache := cache.NewMockICache(s.T())
		mockRegisterRepo.EXPECT().Get(s.ctx, &model.Register{
			AccountID:  1,
			SemesterID: testSemesterID,
			ClassID:    testClassID,
			CourseID:   testCourseID,
		}).ReturnArguments = mock.Arguments{
			&model.Register{
				AccountID:  1,
				SemesterID: testSemesterID,
				CourseID:   testCourseID,
				ClassID:    testClassID,
				CreatedBy:  1,
			}, nil,
		}
		mockRegisterRepo.EXPECT().BatchUpdateSwapIsCanCeledStatus(s.ctx, &model.Register{
			AccountID:  1,
			SemesterID: testSemesterID,
			ClassID:    testClassID,
			CourseID:   testCourseID,
		}).ReturnArguments = mock.Arguments{nil}
		mockCache.EXPECT().Register().ReturnArguments = mock.Arguments{
			func() cache.RegisterService {
				register := cache.NewMockRegisterService(s.T())
				register.EXPECT().ClearRegisterHistories(s.ctx, fmt.Sprintf("%d", 1)).ReturnArguments = mock.Arguments{nil}

				return &cache.MockRegisterService{
					Mock: mock.Mock{
						ExpectedCalls: register.ExpectedCalls,
					},
				}
			}(),
		}
		u := s.useCase(mockRegisterRepo, semester.NewMockRepository(s.T()), class.NewMockRepository(s.T()), course.NewMockRepository(s.T()), mockCache)
		req := &payload.UnRegisterRequest{
			SemesterID: testSemesterID,
			CourseID:   testCourseID,
			ClassID:    testClassID,
		}

		// Act
		reply, err := u.UnRegister(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}

	// bad case
	{
		// Arrange
		req := &payload.UnRegisterRequest{
			ClassID:  testClassID,
			CourseID: testCourseID,
		}
		u := s.useCase(register.NewMockRepository(s.T()), semester.NewMockRepository(s.T()), class.NewMockRepository(s.T()), course.NewMockRepository(s.T()), cache.NewMockICache(s.T()))

		// Act
		_, err := u.UnRegister(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrRegisterInvalidParam("Key: 'UnRegisterRequest.SemesterID' Error:Field validation for 'SemesterID' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
