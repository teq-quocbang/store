package product

import (
	"encoding/csv"
	"fmt"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

// Create List Product With Import File
// @Summary Create List Product
// @Description create a Product
// @Tags Product
// @Accept  json
// @Produce json
// @Security AuthToken
// @Success 200 {object} presenter.ListProductResponseWrapper
// @Router /products/import [post] .
func (r *Route) CreateListWithImportFile(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		resp *presenter.ListProductResponseWrapper
	)

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return teq.Response.Error(c, myerror.ErrProductInvalidParam(err.Error()))
	}

	f, err := fileHeader.Open()
	if err != nil {
		return teq.Response.Error(c, myerror.ErrProductInvalidParam(err.Error()))
	}
	defer f.Close()

	reader := csv.NewReader(f)
	record, err := reader.ReadAll()
	if err != nil {
		return teq.Response.Error(c, myerror.ErrProductInvalidParam(err.Error()))
	}

	req := &payload.CreateListProductRequest{
		Products: make([]payload.Product, len(record)),
	}
	expectedColumn := 3
	for i, v := range record {
		if len(v) < expectedColumn {
			return teq.Response.Error(c, myerror.ErrProductInvalidParam(fmt.Sprintf("line [%d] insufficient column, expected is %d", i, expectedColumn)))
		}
		req.Products[i] = payload.Product{
			Name:        v[0],
			ProductType: v[1],
			ProducerID:  v[2],
		}
	}

	resp, err = r.UseCase.Product.CreateList(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}

// Create List Product
// @Summary Create List Product
// @Description create a Product
// @Tags Product
// @Accept  json
// @Produce json
// @Security AuthToken
// @Param req body payload.CreateProductRequest true "Product info"
// @Success 200 {object} presenter.ListProductResponseWrapper
// @Router /products [post] .
func (r *Route) CreateList(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = &payload.CreateListProductRequest{}
		resp *presenter.ListProductResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Product.CreateList(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
