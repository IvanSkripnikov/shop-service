package helpers

import (
	"database/sql"
	"fmt"

	"loyalty_system/models"

	_ "github.com/go-sql-driver/mysql"

	logger "github.com/IvanSkripnikov/go-logger"
)

var DB *sql.DB

// InitDataBase подключение к БД сервиса
func InitDataBase() (*sql.DB, error) {
	dataSource := GetDatabaseConnectionString(Config.Database)
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		logger.Fatalf("Failed to connect to service database. Error: %v", err)
		return nil, err
	}

	DB = db

	return db, nil
}

func GetDatabaseConnectionString(config models.Database) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.User, config.Password, config.Address, config.Port, config.DB)
	logger.Info(connectionString)
	return connectionString
}
