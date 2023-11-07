package example_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/teq-quocbang/store/delivery/http/example"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
)

func TestGetList(t *testing.T) {
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := example.Route{UseCase: usecase.New(repo, nil)}

	t.Run("200", func(t *testing.T) {
		t.Run("Get list", func(t *testing.T) {
			rec, c := setUpTestGetList(payload.GetListExampleRequest{})

			require.NoError(t, r.GetList(c))
			require.Equal(t, http.StatusOK, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})
}
