package storage

import (
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Upsert
// @Summary Upsert storage
// @Description create a storage
// @Tags Storage
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.UpsertStorageRequest
// @Success 200 {object} presenter.StorageResponseWrapper
// @Router /storage [post] .
func (r *Route) Upsert(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = &payload.UpsertStorageRequest{}
		resp *presenter.StorageResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Storage.UpsertStorage(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
