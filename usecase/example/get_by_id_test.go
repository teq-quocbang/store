package example_test

import (
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/util/myerror"
)

func (suite *TestSuite) TestGetByID_Success() {
	req := &payload.GetByIDRequest{
		ID: 1,
	}

	mockExample := &model.Example{
		ID:        1,
		Name:      "test",
		CreatedBy: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(mockExample, nil)

	// execute
	_, err := suite.useCase.GetByID(suite.ctx, req)

	suite.Nil(err)
}

func (suite *TestSuite) TestGetByID_NotFound() {
	req := &payload.GetByIDRequest{
		ID: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(nil, gorm.ErrRecordNotFound)

	// execute
	_, err := suite.useCase.GetByID(suite.ctx, req)
	expectErr := myerror.ErrExampleNotFound()
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestGetByID_GetByID() {
	req := &payload.GetByIDRequest{
		ID: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(nil, errors.New("error"))

	// execute
	_, err := suite.useCase.GetByID(suite.ctx, req)
	expectErr := myerror.ErrExampleGet(err)
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}
