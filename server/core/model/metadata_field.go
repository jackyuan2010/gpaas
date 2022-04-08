package model

type MetadataField struct {
	GPaasModel
	FieldApiName  string
	ObjectApiName string
	FieldType     int
	Description   string
	FieldLabel    string
	DefineType    string
	Length        int
	IsRequired    bool
	IsUnique      bool
}
