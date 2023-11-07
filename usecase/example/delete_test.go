package example_test

import (
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/util/myerror"
)

func (suite *TestSuite) TestDelete_Success() {
	req := &payload.DeleteExampleRequest{
		ID: 1,
	}

	mockExample := &model.Example{
		ID:        1,
		Name:      "test",
		CreatedBy: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(mockExample, nil)
	suite.mockExampleRepo.On("Delete", suite.ctx, mockExample, false).Return(nil)

	// execute
	err := suite.useCase.Delete(suite.ctx, req)

	suite.Nil(err)
}

func (suite *TestSuite) TestDelete_NotFound() {
	req := &payload.DeleteExampleRequest{
		ID: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(nil, gorm.ErrRecordNotFound)

	// execute
	err := suite.useCase.Delete(suite.ctx, req)
	expectErr := myerror.ErrExampleNotFound()
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestDelete_GetByID() {
	req := &payload.DeleteExampleRequest{
		ID: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(nil, errors.New("error"))

	// execute
	err := suite.useCase.Delete(suite.ctx, req)
	expectErr := myerror.ErrExampleGet(err)
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestDelete_Delete() {
	req := &payload.DeleteExampleRequest{
		ID: 1,
	}

	mockExample := &model.Example{
		ID:        1,
		Name:      "test",
		CreatedBy: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(mockExample, nil)
	suite.mockExampleRepo.On("Delete", suite.ctx, mockExample, false).Return(errors.New("error"))

	// execute
	err := suite.useCase.Delete(suite.ctx, req)
	expectErr := myerror.ErrExampleDelete(err)
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}
