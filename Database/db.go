package database

import (
	"fmt"
	"github.com/caffeine-driven-developers/modi-backend-go/lib"

)
// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbUser := lib.GoDotEnvVariable("DB_USER")
	dbPassword := lib.GoDotEnvVariable("DB_PASSWORD")
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     dbUser,
		DBName:   "modi",
		Password: dbPassword,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

