package statistics

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
)

func (u *UseCase) prepareProductSold(ctx context.Context, req *payload.GetChartRequest) ([]model.CustomerOrder, error) {
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

		return filteredCdrs, nil
	}

	return cdrs, nil
}

func (u *UseCase) GetProductSoldChart(ctx context.Context, req *payload.GetChartRequest) (*presenter.ListStatisticsSoldProductChartResponseWrapper, error) {
	cdrs, err := u.prepareProductSold(ctx, req)
	if err != nil {
		return nil, err
	}

	mSoldProduct, err := u.parseChart(ctx, req.TimeDuration, cdrs)
	if err != nil {
		return nil, err
	}

	result := make([]presenter.SoldProduct, len(mSoldProduct))
	n := 0
	for _, sp := range mSoldProduct {
		result[n] = sp
		n++
	}

	return &presenter.ListStatisticsSoldProductChartResponseWrapper{
		Sold: result,
	}, nil
}

func (u *UseCase) parseChart(ctx context.Context, timeDuration string, filteredCdrs []model.CustomerOrder) (map[string]presenter.SoldProduct, error) {
	mSoldProduct := map[string]presenter.SoldProduct{}
	switch timeDuration {
	case string(payload.WEEK):
		for _, cdr := range filteredCdrs {
			year, week := cdr.CreatedAt.ISOWeek()
			start, end := times.WeekRange(year, week)
			if value, ok := mSoldProduct[fmt.Sprintf("%d%d", year, week)]; ok {
				totalPrice := value.TotalPrice.CoefficientInt64() + (cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty)
				mSoldProduct[fmt.Sprintf("%d%d", year, week)] = presenter.SoldProduct{
					ProductIDs: append(value.ProductIDs, cdr.ProductID),
					SoldQty:    value.SoldQty + cdr.SoldQty,
					TotalPrice: decimal.NewFromInt(totalPrice),
					SoldStart:  start,
					SoldEnd:    end,
				}
			} else {
				totalPrice := cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty
				mSoldProduct[fmt.Sprintf("%d%d", year, week)] = presenter.SoldProduct{
					ProductIDs: []uuid.UUID{cdr.ProductID},
					SoldQty:    cdr.SoldQty,
					TotalPrice: decimal.NewFromInt(totalPrice),
					SoldStart:  start,
					SoldEnd:    end,
				}
			}
		}
	default:
		for _, cdr := range filteredCdrs {
			year, month, day := cdr.CreatedAt.Date()
			if value, ok := mSoldProduct[fmt.Sprintf("%d%v%d", year, month, day)]; ok {
				totalPrice := value.TotalPrice.CoefficientInt64() + (cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty)
				mSoldProduct[fmt.Sprintf("%d%v%d", year, month, day)] = presenter.SoldProduct{
					ProductIDs: append(value.ProductIDs, cdr.ProductID),
					SoldQty:    value.SoldQty + cdr.SoldQty,
					TotalPrice: decimal.NewFromInt(totalPrice),
					SoldStart:  cdr.CreatedAt,
					SoldEnd:    cdr.CreatedAt,
				}
			} else {
				totalPrice := cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty
				mSoldProduct[fmt.Sprintf("%d%v%d", year, month, day)] = presenter.SoldProduct{
					ProductIDs: []uuid.UUID{cdr.ProductID},
					SoldQty:    cdr.SoldQty,
					TotalPrice: decimal.NewFromInt(totalPrice),
					SoldStart:  cdr.CreatedAt,
					SoldEnd:    cdr.CreatedAt,
				}
			}
		}
	}

	return mSoldProduct, nil
}

func (u *UseCase) parseGrowthChart(ctx context.Context, timeDuration string, filteredCdrs []model.CustomerOrder) (map[string]presenter.SoldProduct, error) {
	mSoldProduct := map[string]presenter.SoldProduct{}
	switch timeDuration {
	case string(payload.WEEK):
		var (
			oldKey      string
			growthPrice int64
			growthQty   int64
		)
		for _, cdr := range filteredCdrs {
			year, week := cdr.CreatedAt.ISOWeek()
			start, end := times.WeekRange(year, week)
			currentKey := fmt.Sprintf("%d%d", year, week)
			if value, ok := mSoldProduct[currentKey]; ok {
				totalPrice := value.TotalPrice.CoefficientInt64() + (cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty)
				totalQty := value.SoldQty + cdr.SoldQty

				mSoldProduct[currentKey] = presenter.SoldProduct{
					ProductIDs: append(value.ProductIDs, cdr.ProductID),
					SoldQty:    totalQty,
					TotalPrice: decimal.NewFromInt(totalPrice),
					SoldStart:  start,
					SoldEnd:    end,
				}

				// growth price and qty always up-to-date and remove the start price and product because
				// it was increase in map before
				growthPrice += cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty
				growthQty += cdr.SoldQty
			} else {
				totalPrice := cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty
				totalQty := cdr.SoldQty

				// nếu khác tuần thì cộng dồn
				if oldKey != currentKey {
					totalPrice = totalPrice + growthPrice
					totalQty = totalQty + growthQty
				}
				mSoldProduct[fmt.Sprintf("%d%d", year, week)] = presenter.SoldProduct{
					ProductIDs: []uuid.UUID{cdr.ProductID},
					SoldQty:    totalQty,
					TotalPrice: decimal.NewFromInt(totalPrice),
					SoldStart:  start,
					SoldEnd:    end,
				}

				// growth price and qty always up-to-date
				growthPrice += cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty
				growthQty += cdr.SoldQty
			}
			oldKey = currentKey
		}
	default:
		var (
			oldKey      string
			growthPrice int64
			growthQty   int64
		)
		for _, cdr := range filteredCdrs {
			year, month, day := cdr.CreatedAt.Date()
			currentKey := fmt.Sprintf("%d%v%d", year, month, day)

			if value, ok := mSoldProduct[currentKey]; ok {
				totalPrice := value.TotalPrice.CoefficientInt64() + (cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty)
				totalQty := value.SoldQty + cdr.SoldQty

				mSoldProduct[fmt.Sprintf("%d%v%d", year, month, day)] = presenter.SoldProduct{
					ProductIDs: append(value.ProductIDs, cdr.ProductID),
					SoldQty:    totalQty,
					TotalPrice: decimal.NewFromInt(totalPrice),
					SoldStart:  cdr.CreatedAt,
					SoldEnd:    cdr.CreatedAt,
				}

				// growth price and qty always up-to-date and remove the start price and product because
				// it was increase in map before
				growthPrice += cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty
				growthQty += cdr.SoldQty
			} else {
				totalPrice := cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty
				totalQty := cdr.SoldQty

				// nếu khác ngày thì cộng dồn
				if oldKey != currentKey {
					totalPrice = totalPrice + growthPrice
					totalQty = totalQty + growthQty
				}

				mSoldProduct[currentKey] = presenter.SoldProduct{
					ProductIDs: []uuid.UUID{cdr.ProductID},
					SoldQty:    totalQty,
					TotalPrice: decimal.NewFromInt(totalPrice),
					SoldStart:  cdr.CreatedAt,
					SoldEnd:    cdr.CreatedAt,
				}

				// growth price and qty always up-to-date
				growthPrice += cdr.PriceOfPer.CoefficientInt64() * cdr.SoldQty
				growthQty += cdr.SoldQty
			}
			oldKey = currentKey
		}
	}

	return mSoldProduct, nil
}

func (u *UseCase) GetProductGrowthChart(ctx context.Context, req *payload.GetChartRequest) (*presenter.ListStatisticsSoldProductChartResponseWrapper, error) {
	cdrs, err := u.prepareProductSold(ctx, req)
	if err != nil {
		return nil, err
	}

	mSoldProduct, err := u.parseGrowthChart(ctx, req.TimeDuration, cdrs)
	if err != nil {
		return nil, err
	}

	result := make([]presenter.SoldProduct, len(mSoldProduct))
	n := 0
	for _, sp := range mSoldProduct {
		result[n] = sp
		n++
	}

	return &presenter.ListStatisticsSoldProductChartResponseWrapper{
		Sold: result,
	}, nil
}
