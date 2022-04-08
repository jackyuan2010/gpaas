package main

import (
	"fmt"

	model "github.com/jackyuan2010/gpaas/server/core/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("starting......")
	mt := model.MetadataObject{ObjectApiName: "1212"}
	fmt.Println(mt)

	var fieldType model.FieldType = 4

	switch fieldType {
	case model.Text:
		fmt.Println("Text")
	case model.LongText:
		fmt.Println("LongText")
	case model.Numerric:
		fmt.Println("Numerric")
	case model.Image:
		fmt.Println("Image")
	default:
		fmt.Println("Unknown")
	}

	fmt.Println(model.Image)

	initDB()
}

func initDB() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=172.17.0.2 user=gormuser password=gormuser dbname=gormpg port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database:" + err.Error())
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(100)

	db.AutoMigrate(&model.MetadataObject{})
	db.AutoMigrate(&model.MetadataField{})

	field := model.MetadataField{IsRequired: true, IsUnique: false,
		FieldApiName: "FieldApiName",
		FieldType:    model.Text, FieldLabel: "FieldLabel", Description: "MetaDataFiled"}

	object := model.MetadataObject{ObjectApiName: "MetaDataField"}
	object.ObjectFields = append(object.ObjectFields, field)

	db.Create(&object)

	firstDBEntity := model.MetadataObject{}

	db.First(&firstDBEntity, "object_api_name=?", "MetaDataField")

	fmt.Println(firstDBEntity)

	db.Preload("ObjectFields").First(&firstDBEntity, "object_api_name=?", "MetaDataField")

	fmt.Println(firstDBEntity)
}
