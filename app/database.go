package app

import (
	"bgskoro21/be-pos/helper"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB{
	err := godotenv.Load();
	helper.PanicIfError(err);

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	gormConfig := &gorm.Config{
		PrepareStmt: true,
		SkipDefaultTransaction: true,
		Logger: dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}

	db,err := gorm.Open(postgres.Open(dsn), gormConfig)

	helper.PanicIfError(err)

	return db;
}