package example_test

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/pkg/errors"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/util/myerror"
)

func (suite *TestSuite) TestCreate_Success() {
	req := &payload.CreateExampleRequest{
		Name: teq.String("test"),
	}

	mockExample := &model.Example{
		Name:      *req.Name,
		CreatedBy: 1,
	}

	suite.mockExampleRepo.On("Create", suite.ctx, mockExample).Return(nil)

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	suite.Nil(err)
}

func (suite *TestSuite) TestCreate_Name_Invalid() {
	req := &payload.CreateExampleRequest{
		Name: teq.String(" "),
	}

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)
	expectErr := myerror.ErrExampleInvalidParam("name")
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestCreate_Name_Required() {
	req := &payload.CreateExampleRequest{
		Name: nil,
	}

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)
	expectErr := myerror.ErrExampleInvalidParam("name")
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestCreate_Fail() {
	req := &payload.CreateExampleRequest{
		Name: teq.String("test"),
	}

	mockExample := &model.Example{
		Name:      *req.Name,
		CreatedBy: 1,
	}

	suite.mockExampleRepo.On("Create", suite.ctx, mockExample).Return(errors.New("error"))

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)
	expectErr := myerror.ErrExampleCreate(err)
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}
