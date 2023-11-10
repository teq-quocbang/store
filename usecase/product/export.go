package product

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
	"gopkg.in/yaml.v3"
)

func (u *UseCase) Export(ctx context.Context, req *payload.ExportProductRequest) (*presenter.ListProductResponseWrapper, error) {
	resp := &presenter.ListProductResponseWrapper{}
	products, err := u.Product.GetList(ctx)
	if err != nil {
		return nil, myerror.ErrProductGet(err)
	}
	b := &bytes.Buffer{}
	switch {
	case req.IsYAML():
		e := yaml.NewEncoder(b)
		defer e.Close()

		err := e.Encode(products)
		if err != nil {
			return nil, myerror.ErrProductExportFailed(err)
		}
		resp.Meta = b.String()
	case req.IsJSON():
		e := json.NewEncoder(b)
		err := e.Encode(products)
		if err != nil {
			return nil, myerror.ErrProductExportFailed(err)
		}
		resp.Meta = b.String()
	case req.IsCSV():
		records := make([][]string, len(products))
		for lineNbr, product := range products {
			records[lineNbr] = []string{product.ID.String(), product.Name, product.ProductType, product.ProducerID.String()}
		}

		wr := csv.NewWriter(b)
		wr.WriteAll(records)
		wr.Flush()

		resp.Meta = b.String()
	default:
		records := make([][]string, len(products))
		for lineNbr, product := range products {
			records[lineNbr] = []string{product.Name, product.ProductType, product.ProducerID.String()}
		}
		resp.Meta = records
	}

	return resp, nil
}
