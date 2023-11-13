package product

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

// Create List Product With Thirty Part
// @Summary Create List Product
// @Description create a Product
// @Tags Product
// @Accept  json
// @Produce json
// @Security AuthToken
// @Param req body payload.CreateProductRequest true "Product info"
// @Success 200 {object} presenter.ListProductResponseWrapper
// @Router /products/thirty-part [post] .
func (r *Route) CreateListWithThirtyPart(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = &payload.CreateListWithThirtyPartRequest{}
		resp *presenter.ListProductResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	if err := req.Validate(); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	// get and parse products.
	products := []payload.Product{}
	for _, thirtyPart := range req.ThirtyParts {
		pds, err := getThirtyPartData(thirtyPart)
		if err != nil {
			return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
		}
		products = append(products, pds...)
	}
	createListRequest := &payload.CreateListProductRequest{
		Products: make([]payload.Product, len(products)),
	}
	createListRequest.Products = products

	// create products
	resp, err := r.UseCase.Product.CreateList(ctx, createListRequest)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}

func getThirtyPartData(tpr payload.ThirtyPartRequire) ([]payload.Product, error) {
	// with param required
	if tpr.WithParams() {
		n := 0
		buildParam := ""
		for k, param := range tpr.Params {
			buildParam += fmt.Sprintf("%s%s=%s", getQuerySyntax(n), k, param)
			n++
		}
		tpr.Url = fmt.Sprintf("%s%s", tpr.Url, buildParam)
	}

	req, err := http.NewRequest(http.MethodGet, tpr.Url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request, error: %v", err)
	}

	// with header required
	if tpr.WithHeaders() {
		for k, header := range tpr.Headers {
			req.Header.Set(k, header)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request, error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		responseErr, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read error body, error: %v", err)
		}
		return nil, fmt.Errorf("response error code: [%d] and details: %s", resp.StatusCode, string(responseErr))
	}

	var createListProduct payload.CreateListProductRequest
	if err := json.NewDecoder(resp.Body).Decode(&createListProduct); err != nil {
		return nil, fmt.Errorf("failed to decode body, error: %v", err)
	}

	return createListProduct.Products, nil
}

func getQuerySyntax(n int) string {
	if n > 0 {
		return "&"
	} else {
		return "?"
	}
}
