package storage

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// GetList storage
// @Summary Get an storage
// @Description Get storage by id
// @Tags Storage
// @Accept json
// @Produce json
// @Security AuthToken
// @Success 200 {object} presenter.ListStorageResponseWrapper
// @Router /storage [get] .
func (r *Route) GetList(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.GetStorageByLocatRequest{}
		resp *presenter.ListStorageResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}
	resp, err := r.UseCase.Storage.GetList(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
