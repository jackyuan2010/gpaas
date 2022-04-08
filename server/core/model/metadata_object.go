package model

type MetadataObject struct {
	GPaasModel
	ObjectApiName string
	Description   string
	Module        string
	DefineType    string
	ObjectFields  []MetadataField
}
