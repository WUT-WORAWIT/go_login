package Database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// dsn := "host=127.0.0.1 user=postgres password=123456 dbname=TEST port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=172.17.0.2 user=postgres database=TEST password=banana port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	// db.AutoMigrate(&Models.User{})

	return db
}
func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}
