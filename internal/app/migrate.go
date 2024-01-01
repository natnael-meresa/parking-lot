package app

import (
	"context"
	"database/sql"
	"parking-lot/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func Migrate(ctx context.Context, url string, log *logger.Logger) (*migrate.Migrate, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Error(ctx, "error connecting to database", zap.Error(err))
		return nil, err
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Error(ctx, "error migrating up", zap.Error(err))
		return nil, err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error(ctx, "error migrating up", zap.Error(err))
		return nil, err
	}

	return m, nil
}
