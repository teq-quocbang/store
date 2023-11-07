package example_test

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/util/myerror"
)

func (suite *TestSuite) TestUpdate_Success() {
	req := &payload.UpdateExampleRequest{
		ID:   1,
		Name: teq.String("test"),
	}

	mockExample := &model.Example{
		ID:        1,
		Name:      "test",
		CreatedBy: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(mockExample, nil)
	suite.mockExampleRepo.On("Update", suite.ctx, mockExample).Return(nil)

	// execute
	_, err := suite.useCase.Update(suite.ctx, req)

	suite.Nil(err)
}

func (suite *TestSuite) TestUpdate_NotFound() {
	req := &payload.UpdateExampleRequest{
		ID:   1,
		Name: teq.String("test"),
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(nil, gorm.ErrRecordNotFound)

	// execute
	_, err := suite.useCase.Update(suite.ctx, req)
	expectErr := myerror.ErrExampleNotFound()
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestUpdate_GetByID() {
	req := &payload.UpdateExampleRequest{
		ID:   1,
		Name: teq.String("test"),
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(nil, errors.New("error"))

	// execute
	_, err := suite.useCase.Update(suite.ctx, req)
	expectErr := myerror.ErrExampleGet(err)
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestUpdate_NameInvalid() {
	req := &payload.UpdateExampleRequest{
		ID:   1,
		Name: teq.String(""),
	}

	mockExample := &model.Example{
		ID:        1,
		Name:      "test",
		CreatedBy: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(mockExample, nil)

	// execute
	_, err := suite.useCase.Update(suite.ctx, req)
	expectErr := myerror.ErrExampleInvalidParam("name")
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestUpdate_Update() {
	req := &payload.UpdateExampleRequest{
		ID:   1,
		Name: teq.String("test"),
	}

	mockExample := &model.Example{
		ID:        1,
		Name:      "test",
		CreatedBy: 1,
	}

	suite.mockExampleRepo.On("GetByID", suite.ctx, req.ID).Return(mockExample, nil)
	suite.mockExampleRepo.On("Update", suite.ctx, mockExample).Return(errors.New("error"))

	// execute
	_, err := suite.useCase.Update(suite.ctx, req)
	expectErr := myerror.ErrExampleUpdate(err)
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}
