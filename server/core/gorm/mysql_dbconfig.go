package gorm

import (
	"strings"
)

type MysqlDbConfig struct {
	DbConfig
}

func (dc MysqlDbConfig) Dsn() string {
	var sb strings.Builder

	sb.WriteString(dc.Username)
	sb.WriteString(":")
	sb.WriteString(dc.Password)

	sb.WriteString("@tcp(")
	sb.WriteString(dc.Host)
	sb.WriteString(":")
	sb.WriteString(dc.Port)
	sb.WriteString(")/")
	sb.WriteString(dc.DbName)

	sb.WriteString("?")
	sb.WriteString(dc.Config)
	return sb.String()
}

func NewMysqlDbConfig(
	host string, username string, password string,
	dbname string, port string, config string,
) *MysqlDbConfig {
	mysqldbconfig := MysqlDbConfig{}
	mysqldbconfig.Host = host
	mysqldbconfig.Username = username
	mysqldbconfig.Password = password
	mysqldbconfig.DbName = dbname
	mysqldbconfig.Port = port
	mysqldbconfig.Config = config
	return &mysqldbconfig
}
