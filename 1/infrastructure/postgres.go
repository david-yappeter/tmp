package infrastructure

import (
	"myapp/global"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresSqlDB(postgresConfig global.PostgresConfig, isDebug bool) *sqlx.DB {
	pgxConnConfig, err := pgx.ParseConfig("")
	if err != nil {
		panic(err)
	}

	postgresConnectionConfig := &pgxConnConfig.Config
	postgresConnectionConfig.Host = postgresConfig.Host
	postgresConnectionConfig.Port = postgresConfig.Port
	postgresConnectionConfig.Database = postgresConfig.Database
	postgresConnectionConfig.User = postgresConfig.Username
	postgresConnectionConfig.Password = postgresConfig.Password
	postgresConnectionConfig.RuntimeParams["timezone"] = "UTC"

	pgxDB := stdlib.OpenDB(*pgxConnConfig)
	if err = pgxDB.Ping(); err != nil {
		pgxDB.Close()
		panic(err)
	}

	db := sqlx.NewDb(pgxDB, "pgx")

	return db
}
