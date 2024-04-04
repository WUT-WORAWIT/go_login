package main

import (
	Router "logins/Router"
)

// var err error

// func main() {

// 	DB := database.Init()

//		// Config.DB, err = gorm.Open("postgres", Config.DbURL(Config.BuildDBConfig()))
//		// if err != nil {
//		// 	fmt.Println("Status:", err)
//		// }
//		// defer Config.DB.Close()
//		// Config.DB.AutoMigrate(&Models.User{})
//		r := Routes.SetupRouter()
//		//running
//		r.Run()
//	}
// type User struct {
// 	ID   uint
// 	Name string
// 	Age  uint
// }

func main() {
	// db := Database.Init()
	// var user User

	// if err := db.Raw("SELECT id, name, age FROM users").Scan(&user).Error; err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println("User:", user)
	r := Router.SetupRouter()
	//running
	r.Run()
}
