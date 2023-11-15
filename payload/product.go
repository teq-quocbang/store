package payload

import "github.com/go-playground/validator/v10"

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	ProductType string `json:"product_type" validate:"required"`
	ProducerID  string `json:"producer_id" validate:"required"`
	Price       string `json:"price" validate:"required"`
}

func (p CreateProductRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(p)
}

type UpdateProductRequest struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" `
	ProductType string `json:"product_type"`
	ProducerID  string `json:"producer_id"`
}

func (p UpdateProductRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(p)
}

type Product struct {
	Name        string `json:"name" validate:"required"`
	ProductType string `json:"product_type" validate:"required"`
	ProducerID  string `json:"producer_id" validate:"required"`
	Price       string `json:"price" validate:"required"`
}

type CreateListProductRequest struct {
	Products []Product `json:"products" validate:"required,dive"`
}

func (p CreateListProductRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(p)
}

type ExportProductRequest struct {
	FileExtension string `json:"file_extension"`
}

func (e *ExportProductRequest) IsYAML() bool {
	return e.FileExtension == "yaml"
}

func (e *ExportProductRequest) IsJSON() bool {
	return e.FileExtension == "json"
}

func (e *ExportProductRequest) IsCSV() bool {
	return e.FileExtension == "csv"
}

type ThirtyPartRequire struct {
	Url     string            `json:"url"`
	Params  map[string]string `json:"params"`
	Headers map[string]string `json:"headers"`
}

func (tpr ThirtyPartRequire) WithParams() bool {
	return tpr.Params != nil
}

func (tpr ThirtyPartRequire) WithHeaders() bool {
	return tpr.Headers != nil
}

type CreateListWithThirtyPartRequest struct {
	ThirtyParts []ThirtyPartRequire `json:"thirty_parts" validate:"required,dive"`
}

func (tr *CreateListWithThirtyPartRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(tr)
}
