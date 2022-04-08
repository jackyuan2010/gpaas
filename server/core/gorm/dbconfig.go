package gorm

type DbConfig struct {
	Host               string `json:"host" yaml:"host"`
	Port               string `json:"port" yaml:"port"`
	Config             string `json:"config" yaml:"config"`
	DbName             string `json:"dbname" yaml:"dbname"`
	Username           string `json:"username" yaml:"username"`
	Password           string `json:"password" yaml:"password"`
	MaxIdleConnections int    `json:"max_idle_connections" yaml:"max_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections" yaml:"max_open_connections"`
}

type DbDns interface {
	Dns() string
}
