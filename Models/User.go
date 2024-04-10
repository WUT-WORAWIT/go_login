package Models

import (
	"fmt"
	Db "logins/Database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers Fetch all User data
func GetAllUsers(users *[]User) (err error) {
	db := Db.Init()
	if err := db.Raw("SELECT * FROM users").Scan(users).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser ... Insert New data
func CreateUser(users *User) (err error) {
	db := Db.Init()
	if err := db.Create(&users).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(users *User, username string) (err error) {
	db := Db.Init()
	if err := db.Raw("SELECT * FROM users where username ='" + username + "'").First(&users).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ... Update user
func UpdateUser(c *gin.Context, user *User, id string) (err error) {
	db := Db.Init()
	fmt.Println(user)
	// db.Save(user)
	// ทำการบันทึกข้อมูลผู้ใช้
	if err := db.Save(user).Error; err != nil {
		// ถ้ามีข้อผิดพลาดเกิดขึ้นในการบันทึกข้อมูล สร้าง JSON object สำหรับข้อผิดพลาด
		errorResponse := gin.H{"status": "error", "message": err.Error()}
		// ส่ง JSON กลับไปยังผู้ใช้พร้อมกับ HTTP status code 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, errorResponse)
		// ส่งข้อความข้อผิดพลาดกลับและออกจากฟังก์ชัน
		return err
	}
	// กรณีที่ไม่เกิดข้อผิดพลาด ส่ง JSON กลับไปยังผู้ใช้พร้อมกับ HTTP status code 200 OK
	successResponse := gin.H{"status": "success", "message": "User updated successfully"}
	c.JSON(http.StatusOK, successResponse)
	return nil
}

// DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	db := Db.Init()
	db.Where("id = ?", id).Delete(user)
	return nil
}
