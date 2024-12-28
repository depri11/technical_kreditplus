package db

import (
	"fmt"

	"github.com/depri11/technical_kreditplus/transaction_service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(db *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.DB.User, db.DB.Password, db.DB.Host, db.DB.Port, db.DB.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
