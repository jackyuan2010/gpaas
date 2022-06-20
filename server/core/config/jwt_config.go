package config

type JWTConfig struct {
	SecretKey  string `mapstructure:"scretkey" json:"scretkey" yaml:"scretkey"`
	ExpireTime int64
	BufferTime int64
	Issuer     string
}
