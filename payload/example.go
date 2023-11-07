package payload

import (
	"strings"

	"github.com/teq-quocbang/store/codetype"
)

type CreateExampleRequest struct {
	Name *string `json:"name"`
}

type GetByIDRequest struct {
	ID int64 `json:"-"`
}

var orderByExample = []string{"id", "name", "created_by", "updated_by"}

type GetListExampleRequest struct {
	codetype.Paginator
	SortBy    codetype.SortType `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy   string            `json:"order_by,omitempty" query:"order_by"`
	Search    string            `json:"search,omitempty" query:"search"`
	CreatedBy *int64            `json:"created_by,omitempty" query:"created_by"`
}

func (g *GetListExampleRequest) Format() {
	g.Paginator.Format()
	g.SortBy.Format()
	g.Search = strings.TrimSpace(g.Search)
	g.OrderBy = strings.ToLower(strings.TrimSpace(g.OrderBy))

	for i := range orderByExample {
		if g.OrderBy == orderByExample[i] {
			return
		}
	}

	g.OrderBy = ""
}

type UpdateExampleRequest struct {
	ID   int64   `json:"-"`
	Name *string `json:"name"`
}

type DeleteExampleRequest struct {
	ID int64 `json:"-"`
}
