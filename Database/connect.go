package Database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// host=127.0.0.1 หรือ host=localhost ใช้ในเครื่อง windows
	// dsn := "host=127.0.0.1 user=postgres password=123456 dbname=TEST port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// host=host.docker.internal ที่ใช้ บน docker เพื่อเข้าถึง 127.0.0.1 ในเครื่อง windows
	dsn := "host=host.docker.internal user=postgres password=123456 dbname=TEST port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	if dsn == "" {
		log.Fatal("❌ DSN is empty. Please check your environment variables.")
	}
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
