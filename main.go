package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	gpaasgorm "github.com/jackyuan2010/gpaas/server/core/gorm"
	gpaasgormpg "github.com/jackyuan2010/gpaas/server/core/gorm/postgres"
	model "github.com/jackyuan2010/gpaas/server/core/model"
	gpaasjwt "github.com/jackyuan2010/gpaas/server/gin/jwt"
	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("starting......")

	// config.Viper()
	// initDB()

	jwtTest()
}

func initDB() {
	dbconfig := gpaasgorm.DbConfig{
		Host:     "172.17.0.2",
		Username: "gormuser", Password: "gormuser",
		DbName: "gormpg", Port: "5432",
		Config: "sslmode=disable TimeZone=Asia/Shanghai",
	}

	pgdbconfig := gpaasgormpg.Converte2PostgresDbConfig(&dbconfig)

	// pgdbconfig := gpaasgormpg.NewPostgresDbConfig(
	// 	"172.17.0.2",
	// 	"gormuser",
	// 	"gormuser",
	// 	"gormpg",
	// 	"5432",
	// 	"sslmode=disable TimeZone=Asia/Shanghai",
	// )
	// pgdbconfig.Host = "172.17.0.2"
	// pgdbconfig.Username = "gormuser"
	// pgdbconfig.Password = "gormuser"
	// pgdbconfig.DbName = "gormpg"
	// pgdbconfig.Port = "5432"
	// pgdbconfig.Config = "sslmode=disable TimeZone=Asia/Shanghai"

	var dbdsn gpaasgorm.DbDsn = *pgdbconfig
	// dbdsn = pgdbconfig

	println("database source name:", dbdsn.Dsn())

	// pgdbcontext := gpaasgorm.PostgresDbContext{}

	var dbcontext gpaasgorm.DbContext = gpaasgormpg.PostgresDbContext{}
	// dbcontext = pgdbcontext

	var metadatarepository gpaasgorm.MetadataRepository = gpaasgormpg.NewPostgresMetadataRepository(&dbdsn)

	fmt.Println("Repository method start......")

	dd := metadatarepository.QueryByObjectApiName("MetaDataField")

	fmt.Println(dd)

	fmt.Println("Repository method end......")

	db := dbcontext.GetDb(&dbdsn)

	sqlDB, _ := db.DB()

	defer sqlDB.Close()

	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  "host=172.17.0.2 user=gormuser password=gormuser dbname=gormpg port=5432 sslmode=disable TimeZone=Asia/Shanghai",
	// 	PreferSimpleProtocol: true,
	// }), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database:" + err.Error())
	// }

	// sqlDB, err := db.DB()

	// sqlDB.SetMaxIdleConns(5)
	// sqlDB.SetMaxOpenConns(100)

	// db.AutoMigrate(&model.MetadataObject{})
	// db.AutoMigrate(&model.MetadataField{})

	// field := model.MetadataField{IsRequired: true, IsUnique: false,
	// 	FieldApiName: "FieldApiName",
	// 	FieldType:    model.Text, FieldLabel: "FieldLabel", Description: "MetaDataFiled"}

	// object := model.MetadataObject{ObjectApiName: "MetaDataField"}
	// object.ObjectFields = append(object.ObjectFields, field)

	// db.Create(&object)

	firstDBEntity := model.MetadataObject{}

	db.First(&firstDBEntity, "object_api_name=?", "MetaDataField")

	fmt.Println(firstDBEntity)

	db.Preload("ObjectFields").First(&firstDBEntity, "object_api_name=?", "MetaDataField")

	fmt.Println(firstDBEntity)
}

func jwtTest() {
	jwtClaims := gpaasjwt.JWTClaims{
		UserId:         uuid.NewV4(),
		UserName:       "longyan",
		StandardClaims: jwt.StandardClaims{},
	}
	tokenString := gpaasjwt.GernerateToken(&jwtClaims)
	fmt.Println(tokenString)

	dd := gpaasjwt.ParseToken(tokenString)
	fmt.Println(*dd)

	tokenString = gpaasjwt.Refresh(tokenString)
	fmt.Println(tokenString)

	dd = gpaasjwt.ParseToken(tokenString)
	fmt.Println(*dd)
}
