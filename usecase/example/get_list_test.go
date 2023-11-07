package example_test

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/pkg/errors"

	"github.com/teq-quocbang/store/codetype"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/util/myerror"
)

func (suite *TestSuite) TestGetList_Success() {
	req := &payload.GetListExampleRequest{
		Paginator: codetype.Paginator{
			Page:  1,
			Limit: 20,
		},
		CreatedBy: teq.Int64(1),
		OrderBy:   "name",
		SortBy:    codetype.SortTypeASC,
	}

	mockExamples := []model.Example{
		{
			ID:        1,
			Name:      "test",
			CreatedBy: 1,
		},
	}

	suite.mockExampleRepo.
		On("GetList", suite.ctx, req.Search, req.Paginator, map[string]interface{}{"created_by": req.CreatedBy}, []string{"name ASC"}).
		Return(mockExamples, int64(1), nil)

	// execute
	_, err := suite.useCase.GetList(suite.ctx, req)

	suite.Nil(err)
}

func (suite *TestSuite) TestGetList_GetList() {
	req := &payload.GetListExampleRequest{
		Paginator: codetype.Paginator{
			Page:  1,
			Limit: 20,
		},
	}

	suite.mockExampleRepo.
		On("GetList", suite.ctx, req.Search, req.Paginator, map[string]interface{}{}, []string{}).
		Return(nil, int64(0), errors.New("error"))

	// execute
	_, err := suite.useCase.GetList(suite.ctx, req)
	expectErr := myerror.ErrExampleGet(err)
	myErr := err.(teqerror.TeqError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}
