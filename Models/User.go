package Models

import (
	"log"
	Database "logins/Database"
)

// GetAllUsers Fetch all User data
func GetAllUsers(user *[]User) (err error) {
	db := Database.Init()
	// if err = Config.DB.Find(user).Error; err != nil {
	// 	return err
	// }
	if err := db.Raw("SELECT id, name, age FROM users").Scan(&user).Error; err != nil {
		return err
	}
	log.Println("User:", user)
	return nil
}

// CreateUser ... Insert New data
//
//	if err = Config.DB.Create(user).Error; err != nil {
//		return err
//	}
func CreateUser(user *User) (err error) {
	db := Database.Init()
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	db := Database.Init()
	// if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
	// 	return err
	// }
	if err := db.Raw("SELECT id, name, age FROM users where id ='" + id + "'").Scan(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ... Update user
// func UpdateUser(user *User, id string) (err error) {
// 	fmt.Println(user)
// 	Config.DB.Save(user)
// 	return nil
// }

// DeleteUser ... Delete user
// func DeleteUser(user *User, id string) (err error) {
// 	Config.DB.Where("id = ?", id).Delete(user)
// 	return nil
// }
