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
	"github.com/teq-quocbang/store/util/times"
)

func (u *UseCase) GetProductSoldChart(ctx context.Context, req *payload.GetProductSoldChartRequest) (*presenter.ListStatisticsSoldProductChartResponseWrapper, error) {
	req.Format()
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

		return &presenter.ListStatisticsSoldProductChartResponseWrapper{
			Sold: parseChartSorting(req.TimeDuration, filteredCdrs),
		}, nil
	}

	return &presenter.ListStatisticsSoldProductChartResponseWrapper{
		Sold: parseChartSorting(req.TimeDuration, cdrs),
	}, nil
}

func parseChartSorting(timeDuration string, filteredCdrs []model.CustomerOrder) []presenter.SoldProduct {
	mSoldProduct := map[string]presenter.SoldProduct{}
	switch timeDuration {
	case string(payload.WEEK):
		for _, cdr := range filteredCdrs {
			year, week := cdr.CreatedAt.ISOWeek()
			start, end := times.WeekRange(year, week)
			if value, ok := mSoldProduct[fmt.Sprintf("%d%d", year, week)]; ok {
				mSoldProduct[fmt.Sprintf("%d%d", year, week)] = presenter.SoldProduct{
					ProductIDs: append(value.ProductIDs, cdr.ProductID),
					SoldQty:    value.SoldQty + cdr.SoldQty,
					SoldStart:  start,
					SoldEnd:    end,
				}
			} else {
				mSoldProduct[fmt.Sprintf("%d%d", year, week)] = presenter.SoldProduct{
					ProductIDs: []uuid.UUID{cdr.ProductID},
					SoldQty:    cdr.SoldQty,
					SoldStart:  start,
					SoldEnd:    end,
				}
			}
		}
	default:
		for _, cdr := range filteredCdrs {
			year, month, day := cdr.CreatedAt.Date()
			if value, ok := mSoldProduct[fmt.Sprintf("%d%v%d", year, month, day)]; ok {
				mSoldProduct[fmt.Sprintf("%d%v%d", year, month, day)] = presenter.SoldProduct{
					ProductIDs: append(value.ProductIDs, cdr.ProductID),
					SoldQty:    value.SoldQty + cdr.SoldQty,
					SoldStart:  cdr.CreatedAt,
					SoldEnd:    cdr.CreatedAt,
				}
			} else {
				mSoldProduct[fmt.Sprintf("%d%v%d", year, month, day)] = presenter.SoldProduct{
					ProductIDs: []uuid.UUID{cdr.ProductID},
					SoldQty:    cdr.SoldQty,
					SoldStart:  cdr.CreatedAt,
					SoldEnd:    cdr.CreatedAt,
				}
			}
		}
	}

	result := make([]presenter.SoldProduct, len(mSoldProduct))
	n := 0
	for _, sp := range mSoldProduct {
		result[n] = sp
		n++
	}

	return result
}
