package mysql

import (
	"context"
	"fmt"
	"time"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/teq-quocbang/store/client/logging"
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/util"
)

var db *gorm.DB

func init() {
	var (
		err error
		cfg = config.GetConfig()
	)

	connectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Pass,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName,
	)

	db, err = gorm.Open(mysql.New(mysql.Config{DSN: connectionString}), &gorm.Config{
		Logger: logging.NewGormLogger(),
	})
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	if cfg.Stage.IsLocal() {
		db = db.Debug()
	}

	rawDB, _ := db.DB()
	// rawDB.SetConnMaxIdleTime(time.Hour)
	rawDB.SetMaxIdleConns(cfg.MySQL.DBMaxIdleConns)
	rawDB.SetMaxOpenConns(cfg.MySQL.DBMaxOpenConns)
	rawDB.SetConnMaxLifetime(time.Minute * 5)

	err = rawDB.Ping()
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	teqlogger.GetLogger().Info("Connected mysql db")
}

func GetClient(ctx context.Context) *gorm.DB {
	if util.IsEnableTx(ctx) {
		return util.GetTx(ctx)
	}

	return db.Session(&gorm.Session{})
}
