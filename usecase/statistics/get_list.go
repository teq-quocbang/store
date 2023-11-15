package statistics

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetProductSoldChart(ctx context.Context, req *payload.GetProductSoldChartRequest) (*presenter.ListStatisticsResponseWrapper, error) {
	var (
		order = make([]string, 0)
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	startTime, err := time.Parse("2006-01-02", req.StartTime)
	if err != nil {
		return nil, myerror.ErrCustomerOrderInvalidParam(err.Error())
	}
	endTime, err := time.Parse("2006-01-02", req.EndTime)
	if err != nil {
		return nil, myerror.ErrCustomerOrderInvalidParam(err.Error())
	}

	// get all ordered product in duration time
	cdrs, err := u.Checkout.GetListOrdered(ctx, userPrinciple.User.ID, startTime, endTime, order)
	if err != nil {
		return nil, myerror.ErrCustomerOrderGet(err)
	}

	if req.IsNeedToFilter() {
		// filter product by product type
		productIDs := make([]uuid.UUID, len(cdrs))
		for i, cdr := range cdrs {
			productIDs[i] = cdr.ProductID
		}
		products, err := u.Product.GetListByProductIDs(ctx, productIDs)
		if err != nil {
			return nil, myerror.ErrProductGet(err)
		}
		// filter according product type request
		filteredProducts := lo.Filter[model.Product](products, func(item model.Product, index int) bool {
			return item.ProductType == req.ProductType
		})
		// convert productIds to map
		// this map contain all product ID that filtered
		mProductIDs := map[uuid.UUID]struct{}{}
		for _, p := range filteredProducts {
			mProductIDs[p.ID] = struct{}{}
		}

		// filter and remove
		filteredCdrs := lo.Filter[model.CustomerOrder](cdrs, func(item model.CustomerOrder, index int) bool {
			if _, ok := mProductIDs[item.ProductID]; ok {
				return true
			}
			return false
		})

		return &presenter.ListStatisticsResponseWrapper{
			CustomerOrder: filteredCdrs,
		}, nil
	}

	return &presenter.ListStatisticsResponseWrapper{
		CustomerOrder: cdrs,
	}, nil
}
