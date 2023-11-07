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
	"gorm.io/gorm"
)

func (s *TestSuite) TestCreate() {
	assertion := s.Assertions

	testSemesterID := "S0001"
	testCourseID := "M0001"
	testClassID := "CL0001"
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
		mockSemesterRepo := semester.NewMockRepository(s.T())
		mockCourseRepo := course.NewMockRepository(s.T())
		mockClassRepo := class.NewMockRepository(s.T())
		mockRegisterRepo := register.NewMockRepository(s.T())
		mockCache := cache.NewMockICache(s.T())
		mockRegisterRepo.EXPECT().Get(s.ctx, &model.Register{
			AccountID:  uint(1),
			SemesterID: testSemesterID,
			ClassID:    testClassID,
			CourseID:   testCourseID,
		}).ReturnArguments = mock.Arguments{&model.Register{}, gorm.ErrRecordNotFound}
		mockRegisterRepo.EXPECT().GetListByFirstCourseChar(s.ctx, string(testCourseID[0]), uint(1), testSemesterID).
			ReturnArguments = mock.Arguments{[]model.Register{}, nil}

		mockRegisterRepo.EXPECT().Create(s.ctx, &model.Register{
			AccountID:  1,
			SemesterID: testSemesterID,
			CourseID:   testCourseID,
			ClassID:    testClassID,
			CreatedBy:  uint(1),
		}).ReturnArguments = mock.Arguments{nil}

		mockSemesterRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				ID: testSemesterID,
			}, nil,
		}
		mockClassRepo.EXPECT().GetByID(s.ctx, testClassID).ReturnArguments = mock.Arguments{
			model.Class{
				ID: testSemesterID,
			}, nil,
		}
		mockCourseRepo.EXPECT().GetByID(s.ctx, testCourseID).ReturnArguments = mock.Arguments{
			model.Course{
				ID: testCourseID,
			}, nil,
		}
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

		u := s.useCase(mockRegisterRepo, mockSemesterRepo, mockClassRepo, mockCourseRepo, mockCache)

		// Act
		reply, err := u.Create(s.ctx, &payload.CreateRegisterRequest{
			SemesterID: testSemesterID,
			ClassID:    testClassID,
			CourseID:   testCourseID,
		})

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply.Register.AccountID)
		assertion.NotNil(reply.Register.Semester)
		assertion.NotNil(reply.Register.Class)
		assertion.NotNil(reply.Register.Course)
	}

	// bad case
	{
		// Arrange
		req := &payload.CreateRegisterRequest{
			ClassID:  testClassID,
			CourseID: testCourseID,
		}
		u := s.useCase(register.NewMockRepository(s.T()), semester.NewMockRepository(s.T()), class.NewMockRepository(s.T()), course.NewMockRepository(s.T()), cache.NewMockICache(s.T()))

		// Act
		_, err := u.Create(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrRegisterInvalidParam("Key: 'CreateRegisterRequest.SemesterID' Error:Field validation for 'SemesterID' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
