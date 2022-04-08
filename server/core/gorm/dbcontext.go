package gorm

import (
	"gorm.io/gorm"
)

type DbContext interface {
	GetDb(dsn DbDsn) *gorm.DB
}
