package database

type Config struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DBName       string
	SSLMode      string
	MaxIdleConns int
	MaxOpenConns int
}
