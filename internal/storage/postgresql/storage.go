package postgresql

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/V1merX/upserv-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewConnection(ctx context.Context, log *logrus.Logger, dbConfig *config.PostgreSQLConfig) (*pgxpool.Pool, error) {
	const op = "internal.storage.postgresql.NewConnection"

	var (
		db     *pgxpool.Pool
		pgOnce sync.Once
	)

	connString := strings.TrimSpace(fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d",
		dbConfig.User, dbConfig.Password, dbConfig.DBName,
		dbConfig.Host, dbConfig.Port))

	log.WithFields(logrus.Fields{
		"op": op,
	}).Debug("parsing config")

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"op": op,
		}).Error(err)

		return nil, err
	}

	config.MaxConns = int32(dbConfig.MaxConnections)
	config.MinConns = int32(dbConfig.MinConnections)
	config.MaxConnLifetime = dbConfig.MaxConnLifeTime
	config.MaxConnIdleTime = dbConfig.MaxConnIdleTime
	config.HealthCheckPeriod = dbConfig.HealthCheckPeriod

	pgOnce.Do(func() {
		db, err = pgxpool.NewWithConfig(ctx, config)
	})

	if err = db.Ping(ctx); err != nil {
		log.WithFields(logrus.Fields{
			"op": op,
		}).Error(err)

		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	log.WithFields(logrus.Fields{
		"op": op,
	}).Debug("successfully connected to database")

	return db, nil
}
