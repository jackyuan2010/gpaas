package gorm

import (
	"strings"
)

type MysqlDbConfig struct {
	DbConfig
}

func (dc *MysqlDbConfig) Dns() string {
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
