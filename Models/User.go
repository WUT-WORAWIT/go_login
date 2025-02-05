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

// GetAllUsers Fetch all User data
func GetBranchAll(branchno string) ([]Branch, error) {
	db := Db.Init()
	var branches []Branch
	// ใช้ Rows() เพื่อดึงข้อมูลทั้งหมดโดยตรง
	query := "SELECT branchno, branchname FROM psbranch WHERE branchno = ?"
	rows, err := db.Raw(query, branchno).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close() // ปิด rows หลังใช้เสร็จ

	// อ่านค่าจาก rows ทีละแถว
	for rows.Next() {
		var branch Branch
		err = rows.Scan(&branch.BranchNo, &branch.BranchName) // Scan ค่าแต่ละ Column
		if err != nil {
			return nil, err
		}
		// fmt.Printf("Branch No: %s, Name: %s\n", branchNo, branchName)
		branches = append(branches, branch)
	}

	return branches, nil
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
