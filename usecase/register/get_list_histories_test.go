package register

import (
	"context"
	"fmt"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/cache"
	"github.com/teq-quocbang/store/codetype"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository/class"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/register"
	"github.com/teq-quocbang/store/repository/semester"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestListHistories() {
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
		mockRegisterRepo.EXPECT().GetListRegistered(s.ctx, uint(1), testSemesterID, []string{}, codetype.Paginator{
			Page:  1,
			Limit: 20,
		}).ReturnArguments = mock.Arguments{
			[]model.Register{
				{
					AccountID:  uint(1),
					SemesterID: testSemesterID,
					ClassID:    testClassID,
					CourseID:   testCourseID,
				},
			}, int64(1), nil,
		}

		mockSemesterRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				ID: testSemesterID,
			}, nil,
		}
		mockClassRepo.EXPECT().GetByID(s.ctx, testClassID).ReturnArguments = mock.Arguments{
			model.Class{
				ID: testClassID,
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
				register.EXPECT().GetRegisterHistories(
					s.ctx,
					fmt.Sprintf("%d|%s|%d|%d|%s|%s", 1, testSemesterID, 1, 20, "", "ASC"),
				).ReturnArguments = mock.Arguments{nil, nil}
				register.EXPECT().CreateRegisterHistories(
					s.ctx,
					fmt.Sprintf("%d|%s|%d|%d|%s|%s", 1, testSemesterID, 1, 20, "", "ASC"),
					&presenter.ListRegisterResponseWrapper{
						Register: []presenter.RegisterResponseCustom{
							{
								AccountID: 1,
								Semester: model.Semester{
									ID: testSemesterID,
								},
								Class: model.Class{
									ID: testClassID,
								},
								Course: model.Course{
									ID: testCourseID,
								},
							},
						},
						Meta: map[string]interface{}{
							"page":  1,
							"limit": 20,
							"total": int64(1),
						},
					}).ReturnArguments = mock.Arguments{nil}

				return &cache.MockRegisterService{
					Mock: mock.Mock{
						ExpectedCalls: register.ExpectedCalls,
						Calls:         register.Calls,
					},
				}
			}(),
		}
		u := s.useCase(mockRegisterRepo, mockSemesterRepo, mockClassRepo, mockCourseRepo, mockCache)
		req := &payload.ListRegisteredHistories{
			SemesterID: testSemesterID,
		}

		// Act
		reply, err := u.GetListRegisteredHistories(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.Equal(1, len(reply.Register))
	}
	{ // get with cache
		// Arrange
		mockCache := cache.NewMockICache(s.T())
		mockCache.EXPECT().Register().ReturnArguments = mock.Arguments{
			func() cache.RegisterService {
				register := cache.NewMockRegisterService(s.T())
				register.EXPECT().GetRegisterHistories(s.ctx, fmt.Sprintf("%d|%s|%d|%d|%s|%s", 1, testSemesterID, 1, 20, "", "ASC")).ReturnArguments = mock.Arguments{
					&presenter.ListRegisterResponseWrapper{
						Register: []presenter.RegisterResponseCustom{
							{
								AccountID: 1,
								Semester: model.Semester{
									ID: testSemesterID,
								},
								Class: model.Class{
									ID: testClassID,
								},
								Course: model.Course{
									ID: testCourseID,
								},
							},
						},
						Meta: map[string]interface{}{
							"page":  1,
							"limit": 20,
							"total": int64(1),
						},
					}, nil,
				}
				return &cache.MockRegisterService{
					Mock: mock.Mock{
						ExpectedCalls: register.ExpectedCalls,
					},
				}
			}(),
		}
		req := &payload.ListRegisteredHistories{
			SemesterID: testSemesterID,
		}
		u := s.useCase(register.NewMockRepository(s.T()), semester.NewMockRepository(s.T()), class.NewMockRepository(s.T()), course.NewMockRepository(s.T()), mockCache)

		// Act
		reply, err := u.GetListRegisteredHistories(s.ctx, req)

		// Arrange
		assertion.NoError(err)
		assertion.Equal(1, len(reply.Register))
	}
}
