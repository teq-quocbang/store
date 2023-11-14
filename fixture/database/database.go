package database

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Register using Golang migrate.
	"github.com/teq-quocbang/store/cache"
	"github.com/teq-quocbang/store/cache/connection"
	"github.com/teq-quocbang/store/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// list tables in DB.
var tables = []string{
	"examples",
	"account",
	"product",
	"producer",
	"storage",
}

type Database struct {
	DB *gorm.DB
}

func InitDatabase() *Database {
	connectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_TEST_USERNAME"),
		os.Getenv("DB_TEST_PASSWORD"),
		os.Getenv("DB_TEST_HOST"),
		os.Getenv("DB_TEST_PORT"),
		os.Getenv("DB_TEST_NAME"),
	)

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: connectionString}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		teqlogger.GetLogger().Fatal(err.Error())
	}

	// migration to test database
	migrations.Up(db, os.Getenv("DB_TEST_MIGRATION_PATH"), os.Getenv("DB_TEST_NAME"))

	return &Database{DB: db.Session(&gorm.Session{})}
}

func InitCache() cache.ICache {
	port, err := strconv.Atoi(os.Getenv("REDIS_TEST_PORT"))
	if err != nil {
		teqlogger.GetLogger().Fatal(err.Error())
	}

	return connection.NewRedisCache(connection.RedisConfig{
		Address:  os.Getenv("REDIS_TEST_HOST"),
		Port:     port,
		Password: os.Getenv("REDIS_TEST_PASSWORD"),
	})
}
func (d *Database) TruncateTables() {
	d.DB.Exec("SET FOREIGN_KEY_CHECKS=0")
	defer d.DB.Exec("SET FOREIGN_KEY_CHECKS=1")

	for i := range tables {
		err := d.DB.Table(tables[i]).Exec(fmt.Sprintf("TRUNCATE TABLE %s", tables[i])).Error
		if err != nil {
			teqlogger.GetLogger().Fatal(err.Error())
		}
	}
}

func (d *Database) ExecFixture(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	for scanner.Scan() {
		query := scanner.Text()
		if err = d.DB.Exec(query).Error; err != nil {
			return err
		}
	}

	return nil
}

func (d *Database) GetClient(_ context.Context) *gorm.DB {
	return d.DB
}
