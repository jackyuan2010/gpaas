package gorm

import (
	model "github.com/jackyuan2010/gpaas/server/core/model"
)

type Repository interface {
	QueryById(metadataObject model.MetadataObject, id string) map[string]interface{}
}
