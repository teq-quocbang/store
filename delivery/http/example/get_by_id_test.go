package example_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/teq-quocbang/store/delivery/http/example"
	"github.com/teq-quocbang/store/fixture/database"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase"
)

func TestGetByID(t *testing.T) {
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := example.Route{UseCase: usecase.New(repo, nil)}

	t.Run("200", func(t *testing.T) {
		t.Run("Get by ID", func(t *testing.T) {
			// prepare data for testing this test case
			require.NoError(t, db.ExecFixture("../../../fixture/example/example.sql"))

			rec, c := setUpTestGetByID("1")

			require.NoError(t, r.GetByID(c))
			require.Equal(t, http.StatusOK, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})

	t.Run("400", func(t *testing.T) {
		t.Run("Invalid param", func(t *testing.T) {
			rec, c := setUpTestGetByID("1n")

			require.NoError(t, r.GetByID(c))
			require.Equal(t, http.StatusBadRequest, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})

	t.Run("404", func(t *testing.T) {
		t.Run("Not found", func(t *testing.T) {
			rec, c := setUpTestGetByID("1")

			require.NoError(t, r.GetByID(c))
			require.Equal(t, http.StatusNotFound, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})
}
