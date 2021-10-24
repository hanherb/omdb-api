package services

import (
	"fmt"
	"log"

	"github.com/hanherb/omdb-api/lib/variables"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func MysqlInitialization() {
	config := variables.Mysql

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Host,
		config.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	DB = db
}
