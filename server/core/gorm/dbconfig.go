package gorm

type DbDsn interface {
	Dsn() string
}

type DbConfig struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	DbName       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max_idle_connections" json:"max_idle_connections" yaml:"max_idle_connections"`
	MaxOpenConns int    `mapstructure:"max_open_connections" json:"max_open_connections" yaml:"max_open_connections"`
}

func (dc DbConfig) Dsn() string {
	return ""
}
