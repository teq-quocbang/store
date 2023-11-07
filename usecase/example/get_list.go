package example

import (
	"context"
	"fmt"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListExampleRequest,
) (*presenter.ListExampleResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}

	if req.CreatedBy != nil && *req.CreatedBy > 0 {
		conditions["created_by"] = req.CreatedBy
	}

	myExamples, total, err := u.ExampleRepo.GetList(ctx, req.Search, req.Paginator, conditions, order)
	if err != nil {
		return nil, myerror.ErrExampleGet(err)
	}

	return &presenter.ListExampleResponseWrapper{
		Examples: myExamples,
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}, nil
}
