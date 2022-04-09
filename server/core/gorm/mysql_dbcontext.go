package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDbContext struct {
}

func (mysqldc MysqlDbContext) GetDb(dsn *DbDsn) *gorm.DB {
	mysqlconfig := mysql.Config{
		DSN:                       (*dsn).Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}

	if db, err := gorm.Open(mysql.New(mysqlconfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		mysqlconfig := (*dsn).(MysqlDbConfig)
		sqlDB.SetMaxIdleConns(mysqlconfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysqlconfig.MaxOpenConns)
		return db
	}
}
