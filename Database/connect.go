package Database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=TEST port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	// db.AutoMigrate(&Models.User{})

	return db
}

// import (
// 	"logins/Models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect() {

// 	// กำหนดข้อมูลเชื่อมต่อฐานข้อมูล PostgreSQL
// 	dsn := "host=localhost user=postgres password=123456 dbname=TEST port=5432 sslmode=disable"

// 	// เชื่อมต่อฐานข้อมูล
// 	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	DB = database

// 	database.AutoMigrate(&Models.User{})
// }
