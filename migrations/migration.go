package migrations

//nolint:revive
import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Up(db *gorm.DB, migratePath string, databaseName string) {
	getDB, err := db.DB()
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	driver, err := mysql.WithInstance(getDB, &mysql.Config{MigrationsTable: "migration"})
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(migratePath, databaseName, driver)
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	teqlogger.GetLogger().Info("Up done!")
}
