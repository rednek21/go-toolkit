package database

type Config interface {
	GetHost() string
	GetPort() int
	GetUsername() string
	GetPassword() string
	GetDBName() string
	GetSSLMode() string
	GetMaxIdleConns() int
	GetMaxOpenConns() int
}

type BaseConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DBName       string
	SSLMode      string
	MaxIdleConns int
	MaxOpenConns int
}

func NewBaseConfig(host string, port int, username, password, dbName, sslMode string, maxIdleConns, maxOpenConns int) *BaseConfig {
	return &BaseConfig{
		Host:         host,
		Port:         port,
		Username:     username,
		Password:     password,
		DBName:       dbName,
		SSLMode:      sslMode,
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
	}
}

func (c *BaseConfig) GetHost() string      { return c.Host }
func (c *BaseConfig) GetPort() int         { return c.Port }
func (c *BaseConfig) GetUsername() string  { return c.Username }
func (c *BaseConfig) GetPassword() string  { return c.Password }
func (c *BaseConfig) GetDBName() string    { return c.DBName }
func (c *BaseConfig) GetSSLMode() string   { return c.SSLMode }
func (c *BaseConfig) GetMaxIdleConns() int { return c.MaxIdleConns }
func (c *BaseConfig) GetMaxOpenConns() int { return c.MaxOpenConns }
