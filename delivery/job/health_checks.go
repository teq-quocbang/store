package job

import (
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

type healthChecks struct {
	url string
}

func NewHealthChecks(url string) IJob {
	return healthChecks{
		url: url,
	}
}

func (h healthChecks) callHealthCheck() {
	if h.url != "" {
		resp, err := http.Get(h.url)
		if err != nil {
			teqlogger.GetLogger().Fatal("failed to call request to health check")
			teqsentry.Fatal(errors.Wrap(err, "failed to call request to health check"))

			return
		}

		if resp != nil {
			resp.Body.Close()
		}
	}
}

func (h healthChecks) Run() {
	c := cron.New()

	_, err := c.AddFunc("*/1 * * * *", func() { h.callHealthCheck() })
	if err != nil {
		teqlogger.GetLogger().Fatal("failed to schedule health check short period")
		teqsentry.Fatal(errors.Wrap(err, "failed to schedule health check short period"))
	}

	c.Start()
}
