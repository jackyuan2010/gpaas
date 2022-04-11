package gorm

import (
	"strings"

	gpaasgorm "github.com/jackyuan2010/gpaas/server/core/gorm"
)

type PostgresDbConfig struct {
	gpaasgorm.DbConfig `mapstructure:"dbconfig" json:"dbconfig" yaml:"dbconfig"`
}

func (dc PostgresDbConfig) Dsn() string {
	var sb strings.Builder
	sb.WriteString("host=")
	sb.WriteString(dc.Host)

	sb.WriteString(" user=")
	sb.WriteString(dc.Username)

	sb.WriteString(" password=")
	sb.WriteString(dc.Password)

	sb.WriteString(" dbname=")
	sb.WriteString(dc.DbName)

	sb.WriteString(" port=")
	sb.WriteString(dc.Port)

	sb.WriteString(" ")
	sb.WriteString(dc.Config)
	return sb.String()
}

func NewPostgresDbConfig(
	host string, username string, password string,
	dbname string, port string, config string,
) *PostgresDbConfig {
	pgdbconfig := PostgresDbConfig{}
	pgdbconfig.Host = host
	pgdbconfig.Username = username
	pgdbconfig.Password = password
	pgdbconfig.DbName = dbname
	pgdbconfig.Port = port
	pgdbconfig.Config = config
	return &pgdbconfig
}

func Converte2PostgresDbConfig(dbconfig *gpaasgorm.DbConfig) *PostgresDbConfig {
	pgdbconfig := PostgresDbConfig{}
	pgdbconfig.DbConfig = *dbconfig
	return &pgdbconfig
}
