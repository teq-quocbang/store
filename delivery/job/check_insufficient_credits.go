package job

import (
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

type insufficientCreditsCheck struct {
	url string
}

func NewInsufficientCreditsCheck(url string) IJob {
	return insufficientCreditsCheck{
		url: url,
	}
}

func (i insufficientCreditsCheck) call() {
	resp, err := http.Get(i.url)
	if err != nil {
		teqlogger.GetLogger().Fatal("failed to call request to insufficient credits check")
		teqsentry.Fatal(errors.Wrap(err, "failed to call request to insufficient credits check"))

		return
	}

	if resp != nil {
		resp.Body.Close()
	}
}

func (i insufficientCreditsCheck) Run() {
	c := cron.New()

	_, err := c.AddFunc("1 * * * *", func() { i.call() })
	if err != nil {
		teqlogger.GetLogger().Fatal("failed to schedule insufficient credits check short period")
		teqsentry.Fatal(errors.Wrap(err, "failed to schedule insufficient credits check short period"))
	}

	c.Start()
}
