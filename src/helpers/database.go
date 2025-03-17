package helpers

import (
	"database/sql"

	"loyalty_system/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/IvanSkripnikov/go-gormdb"
	"github.com/IvanSkripnikov/go-logger"
)

var DB *sql.DB

func InitDatabase(config gormdb.Database) {
	gormDatabase, err := gormdb.AddMysql(models.ServiceDatabase, config)
	if err != nil {
		logger.Fatalf("Cant initialize DB: %v", err)
	}
	db, err := gormDatabase.DB()
	if err != nil {
		logger.Fatalf("Cant get DB: %v", err)
	}
	DB = db
}
