package connection

import (
	"fmt"
	"log"
	"server_with_database/pkg/models"
	"server_with_database/pkg/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var db *gorm.DB

func dbConnect() {
	config := utils.GetConfig()

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUser, config.DbPass, config.DbIP, config.DbName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("error connecting to Database", err)
	}

	log.Println("Database Connected")
	db = d

	db.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(10).
			SetMaxOpenConns(100))
}

func migrate() {
	// db.Migrator().DropTable(&models.Product{})
	// db.Migrator().DropTable(&models.Profile{})
	db.AutoMigrate(&models.Profile{}, &models.Product{})
}

func GetDB() *gorm.DB {
	// if db == nil {
	dbConnect()
	// }
	migrate()
	return db
}
