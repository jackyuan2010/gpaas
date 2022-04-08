package model

type MetadataField struct {
	GPaasModel
	MetadataObjectID string
	FieldApiName     string
	ObjectApiName    string
	FieldType        int
	Description      string
	FieldLabel       string
	DefineType       string
	Length           int
	IsRequired       bool
	IsUnique         bool
}

func (entity *MetadataField) TableName() string {
	return "metadata_object_field"
}
