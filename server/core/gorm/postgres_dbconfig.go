package gorm

import (
	"strings"
)

type PostgresDbConfig struct {
	DbConfig
}

func (dc *PostgresDbConfig) Dns() string {
	sb := strings.Builder()
	sb.append

	return "host=172.17.0.2 user=gormuser password=gormuser dbname=gormpg port=5432 sslmode=disable TimeZone=Asia/Shanghai"
}
