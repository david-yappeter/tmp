package infrastructure

import (
	"fmt"
	"myapp/global"

	"github.com/golang-migrate/migrate/v4"
	migratePgx "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfrastructureManager interface {
	GetSqlDB() *sqlx.DB
	GetGormDB() *gorm.DB
	MigrateDB(migrationDir string, isRollingBack bool, steps int, force *int) error
	RefreshDB() error
	CloseDB() error
}

type infrastructureManager struct {
	sqlDB  *sqlx.DB
	gormDB *gorm.DB
}

func NewInfrastructureManager(config global.YamlConfig) InfrastructureManager {
	postgresConfig := config.Postgres

	sqlDB := NewPostgresSqlDB(postgresConfig, config.IsDebug)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		panic(err)
	}

	return &infrastructureManager{
		sqlDB:  sqlDB,
		gormDB: gormDB,
	}
}

func (i *infrastructureManager) createDB() error {
	dbConfig := global.GetPostgresConfig()
	dbConfig.Database = ""

	isDebug := global.IsDebug()

	sqlDB := NewPostgresSqlDB(dbConfig, isDebug)

	if _, err := sqlDB.Exec(fmt.Sprintf(`CREATE DATABASE "%s" WITH ENCODING='UTF8';`, global.GetPostgresConfig().Database)); err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	i.sqlDB = NewPostgresSqlDB(global.GetPostgresConfig(), isDebug)

	return nil
}

func (i *infrastructureManager) dropDB() error {
	dbConfig := global.GetPostgresConfig()
	dbConfig.Database = ""

	isDebug := global.IsDebug()

	sqlDB := NewPostgresSqlDB(dbConfig, isDebug)

	if _, err := sqlDB.Exec(fmt.Sprintf(`DROP DATABASE IF EXISTS "%s" WITH (FORCE);`, global.GetPostgresConfig().Database)); err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}

func (i *infrastructureManager) GetSqlDB() *sqlx.DB {
	return i.sqlDB
}

func (i *infrastructureManager) GetGormDB() *gorm.DB {
	return i.gormDB
}

func (i *infrastructureManager) MigrateDB(migrationDir string, isRollingBack bool, steps int, force *int) error {
	dbDriver, err := migratePgx.WithInstance(i.GetSqlDB().DB, &migratePgx.Config{})
	if err != nil {
		return err
	}

	d, err := (&file.File{}).Open(fmt.Sprintf("file://%s", migrationDir))
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithInstance("", d, "pgx", dbDriver)
	if err != nil {
		return err
	}

	if force != nil {
		migrator.Force(*force)
		return nil
	}

	if isRollingBack {
		_, _, err := migrator.Version()
		if err != nil {
			return err
		}

		if steps > 0 {
			err = migrator.Steps(-1 * int(steps))
		} else {
			err = migrator.Down()
		}

		if err != nil {
			return err
		}
	} else {
		var err error
		if steps > 0 {
			err = migrator.Steps(int(steps))
		} else {
			err = migrator.Up()
		}

		if err != nil && err != migrate.ErrNoChange {
			return err
		}
	}

	return nil
}

func (i *infrastructureManager) RefreshDB() error {
	if err := i.dropDB(); err != nil {
		return err
	}

	if err := i.createDB(); err != nil {
		return err
	}

	if err := i.MigrateDB(global.GetMigrationDir(), false, 0, nil); err != nil {
		return err
	}

	return nil
}

func (i *infrastructureManager) CloseDB() error {
	if sqlDB := i.GetSqlDB(); sqlDB != nil {
		sqlDB.Close()
	}
	return nil
}
