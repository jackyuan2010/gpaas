package config

type JWTConfig struct {
	SigningKey string `mapstructure:"dbtype" json:"dbtype" yaml:"dbtype"`
	ExpireTime int64
	BufferTime int64
	Issuer     string
}
