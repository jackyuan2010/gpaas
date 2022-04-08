package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDbContext struct {
}

func (pgdc *MysqlDbContext) GetDb(dsn MysqlDbConfig) *gorm.DB {
	mysqlconfig := mysql.Config{
		DSN:                      dsn.Dsn(),
		DefaultStringSize:        191,
		SkipInitalizeWithVersion: false,
	}

	if db, err := gorm.Open(mysql.New(mysqlconfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dsn.MaxIdleConnections)
		sqlDB.SetMaxOpenConns(dsn.MaxOpenConnections)
		return db
	}
}
