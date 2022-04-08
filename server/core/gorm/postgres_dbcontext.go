package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDbContext struct {
}

func (pgdc *PostgresDbContext) GetDb(dsn PostgresDbConfig) *gorm.DB {
	pgsqlconfig := postgres.Config{
		DSN:                  dsn.Dsn(),
		PreferSimpleProtocol: false,
	}

	if db, err := gorm.Open(postgres.New(pgsqlconfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dsn.MaxIdleConnections)
		sqlDB.SetMaxOpenConns(dsn.MaxOpenConnections)
		return db
	}
}
