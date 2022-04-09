package gorm

import (
	model "github.com/jackyuan2010/gpaas/server/core/model"
)

type MetadataRepository interface {
	QueryById(id string) *model.MetadataObject
	QueryByObjectApiName(objectApiName string) *model.MetadataObject
}
