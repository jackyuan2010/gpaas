package model

type MetadataObject struct {
	GPaasModel
	ObjectApiName  string `gorm:"unique:true"`
	Description    string
	Module         string
	DefineType     string
	StoreTableName string
	ObjectFields   []MetadataField `gorm:"foreignKey:ObjectApiName;references:ObjectApiName"`
}

func (entity *MetadataObject) TableName() string {
	return "metadata_object"
}
