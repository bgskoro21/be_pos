package app

import (
	"bgskoro21/be-pos/helper"
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

	dsn := os.Getenv("DATABASE_URL")

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