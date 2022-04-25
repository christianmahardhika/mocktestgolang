package config

type AppConfig struct {
	Database Database
	Port     string
}

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}
