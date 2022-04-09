package gorm

type DbDsn interface {
	Dsn() string
}

type DbConfig struct {
	Host         string `json:"host" yaml:"host"`
	Port         string `json:"port" yaml:"port"`
	Config       string `json:"config" yaml:"config"`
	DbName       string `json:"dbname" yaml:"dbname"`
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	MaxIdleConns int    `json:"max_idle_connections" yaml:"max_idle_connections"`
	MaxOpenConns int    `json:"max_open_connections" yaml:"max_open_connections"`
}

func (dc DbConfig) Dsn() string {
	return ""
}
