package gorm

import (
	gpaasgorm "github.com/jackyuan2010/gpaas/server/core/gorm"
	model "github.com/jackyuan2010/gpaas/server/core/model"
	"gorm.io/gorm"
)

type PostgresMetadataRepository struct {
	dsn       *gpaasgorm.DbDsn
	dbcontext *gpaasgorm.DbContext
	DB        *gorm.DB
}

func NewPostgresMetadataRepository(dsn *gpaasgorm.DbDsn) *PostgresMetadataRepository {
	pgrepository := PostgresMetadataRepository{}
	pgrepository.dsn = dsn
	var dbcontext gpaasgorm.DbContext = PostgresDbContext{}
	pgrepository.dbcontext = &dbcontext
	pgrepository.DB = (*pgrepository.dbcontext).GetDb(pgrepository.dsn)
	return &pgrepository
}

func (pgrepository PostgresMetadataRepository) QueryByObjectApiName(objectApiName string) *model.MetadataObject {
	metadataObject := model.MetadataObject{}

	pgrepository.DB.Preload("ObjectFields").First(&metadataObject, "object_api_name=?", objectApiName)

	return &metadataObject
}

func (pgrepository PostgresMetadataRepository) QueryById(id string) *model.MetadataObject {
	metadataObject := model.MetadataObject{}

	pgrepository.DB.Preload("ObjectFields").First(&metadataObject, "id=?", id)

	return &metadataObject
}
