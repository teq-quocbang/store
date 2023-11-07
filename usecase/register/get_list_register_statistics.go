package register

import (
	"context"
	"fmt"
	"time"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"github.com/teq-quocbang/store/codetype"
	"github.com/teq-quocbang/store/util/myerror"
	"go.uber.org/zap"
)

func (u *UseCase) TracingInsufficientCreditsStatistics(ctx context.Context) error {
	// get all account
	accounts, err := u.Account.GetList(ctx)
	if err != nil {
		return myerror.ErrAccountGet(err)
	}
	currentYear := time.Now().Year()
	// get semester by year
	semesters, err := u.Semester.GetListByYear(ctx, fmt.Sprint(currentYear))
	if err != nil {
		return myerror.ErrSemesterGet(err)
	}

	teqlogger.GetLogger().Warn("Start tracing insufficient credits")

	var totalCredits int
	for _, s := range semesters {
		if time.Now().After(s.RegisterStartAt) && time.Now().Before(s.RegisterExpiresAt) {
			for _, a := range accounts {
				// get registered by account
				registers, _, err := u.Register.GetListRegistered(ctx, a.ID, s.ID, []string{}, codetype.Paginator{})
				if err != nil {
					return myerror.ErrRegisterGet(err)
				}

				// for loop  to receive class id and start to sum
				for _, r := range registers {
					// continue if that order was canceled
					if r.IsCanceled {
						continue
					}

					// get credit of each class
					class, err := u.Class.GetByID(ctx, r.ClassID)
					if err != nil {
						return myerror.ErrClassGet(err)
					}

					totalCredits += int(class.Credits)
				}

				// check credits
				if totalCredits < s.MinCredits {
					fields := []zap.Field{
						zap.Uint("student_id", a.ID),
						zap.String("semester_id", s.ID),
						zap.Int("total_registered_credits", totalCredits),
						zap.Int("required_credits", s.MinCredits),
					}
					teqlogger.GetLogger().Warn("Insufficient credits", fields...)
				}

				totalCredits = 0
			}
		}
	}

	teqlogger.GetLogger().Warn("End tracing")

	return nil
}
