package storage

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
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
// @Router /storage/:locat [get] .
func (r *Route) GetList(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = c.Param("locat")
		resp *presenter.ListStorageResponseWrapper
	)

	resp, err := r.UseCase.Storage.GetListByLocat(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
