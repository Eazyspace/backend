package db

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBContext context.Context

func getDBContext() context.Context {
	if DBContext == nil {
		DBContext, _ = context.WithTimeout(context.Background(), 10*time.Second)
	}
	return DBContext
}

// Create only 1 database
func CreateUniversalDB(uri string, db string) *gorm.DB {
	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	DB, _ := database.DB()
	err = DB.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}
	database.Logger.LogMode(logger.Error)
	
	return database
	// client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// err = client.Connect(getDBContext())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return client.Database(db)
}
