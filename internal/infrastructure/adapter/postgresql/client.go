package postgresql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lucaslmuller/technical-test/internal/infrastructure"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/uptrace/bun/driver/pgdriver"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect(res *infrastructure.Resources) {
	cfg := res.Config.Postgres

	if cfg.Host != nil {
		conn, err := connect(cfg, cfg.Host, cfg.TimeoutSec)
		if err != nil {
			res.Logger.Error("failed to connect to PostgreSQL")
			panic(err)
		}

		res.PostgreSQL = conn
		res.Logger.Info("PostgreSQL connection was successful")
	}
}

func connect(cfg *infrastructure.Postgres, host *string, timeoutSec *int) (*bun.DB, error) {
	// Prepare connection string parameterized
	connectionParams := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		*host,
		cfg.Port,
		cfg.DBName,
	)

	timeoutDuration := time.Second * 45
	if timeoutSec != nil {
		timeoutDuration = time.Second * time.Duration(*timeoutSec)
	}

	// Connect to database
	connection := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionParams), pgdriver.WithTimeout(timeoutDuration)))
	if err := connection.Ping(); err != nil {
		return nil, err
	}

	// set connection settings
	connection.SetMaxOpenConns(cfg.Settings.MaxOpenConns)
	connection.SetMaxIdleConns(cfg.Settings.MaxIdleConns)
	connection.SetConnMaxLifetime(time.Duration(cfg.Settings.MaxConnLifeTimeInMin) * time.Minute)

	return bun.NewDB(connection, pgdialect.New()), nil
}

func Disconnect(res *infrastructure.Resources) {
	err := res.PostgreSQL.Close()
	if err != nil {
		res.Logger.Error("error when closing connection with PostgreSQL")
		return
	}
}
