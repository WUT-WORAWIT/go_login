package Controllers

import (
	"fmt"
	"logins/Models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers ... Get all users
func GetRemark(c *gin.Context) {
	str := `
	func: CreateUser
	POST(http://localhost:8080/user-api/creatuser)
	Body -> JSON
	{
		"Username": "worawit",
	   	"Password": "12345678",
		"Prefix": "นาย",
		"First_name": "วรวิทย์",
	   	"Last_name": "จันรอง",
	   	"Email": "worawit@test.com",
	   	"Phone_number": "0812345622",
		"Date_of_birth": "25390709"
	}
	 
	-----------------------------------------------
	func: GetUsers
	GET(http://localhost:8080/user-api/getuserall)
	paramiter:

	-----------------------------------------------
	func: GetUserByID
	GET(http://localhost:8080/user-api/getuserbyid?id=1)
	paramiter:
	Username=

	-----------------------------------------------
	func: UpdateUser
	http://localhost:8080/user-api/updateuser?id=2
	Body -> JSON
	 {
	 	"ID": 2,
	 	"Name": "test2",
	 	"Age": 19
	 }
	paramiter:
	id
	 `
	c.String(http.StatusOK, str)
}

// GetUsers ... Get all users
func GetUsersAll(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON request"})
		return
	}
	fmt.Println(user.Password)
	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate hashed password"})
		return
	}
	user.Password = string(hashedPassword)

	// Create user in database
	err = Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByUsername ... Get the user by id
func GetUserByID(c *gin.Context) {
	// id := c.Params.ByName("id")
	Username := c.Query("Username")
	var user Models.User
	err := Models.GetUserByID(&user, Username)
	if err != nil {
		errorResponse := gin.H{"status": "error", "message": err.Error()}
		// ส่ง JSON กลับไปยังผู้ใช้พร้อมกับ HTTP status code 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
		// c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	// id := c.Params.ByName("id")
	id := c.Query("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		errorResponse := gin.H{"status": "error", "message": err.Error()}
		// ส่ง JSON กลับไปยังผู้ใช้พร้อมกับ HTTP status code 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
		// c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(c, &user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	// else {
	// 	c.JSON(http.StatusOK, user)
	// }
}

// DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	// id := c.Params.ByName("id")
	id := c.Query("id")

	err := Models.GetUserByID(&user, id)
	if err != nil {
		// ส่ง JSON กลับไปยังผู้ใช้พร้อมกับ HTTP status code 500 Internal Server Error
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User with id " + id + " not found"})
		return
		// c.JSON(http.StatusNotFound, user)
	}
	err = Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	// หากข้อมูลถูกลบสำเร็จ
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User with id " + id + " has been deleted"})
}

// Login ... Login the user
func Login(c *gin.Context) {

	// - c.Query: เป็นการดึงข้อมูลจาก parameter ของ request (คล้ายกับ context.params ใน koa)
	// - c.PostForm: จะเป็นการดึงข้อมูลจาก Multipart หรือ Urlencoded Form (คล้ายกับ context.request.body ใน koa)
	// - c.ShouldBindJSON: คือการ bind request body ให้เป็นในรูปแบบ JSON กับ struct
	// - c.Json: เป็นการคืนค่า response กลับไปให้ request

	var user Models.User
	var userjson Models.UserJson

	// ดำเนินการแปลง JSON request body เป็นโครงสร้าง User
	if err := c.ShouldBindJSON(&userjson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Models.GetUserByID(&user, userjson.Username)
	if err != nil {
		// หากไม่พบข้อมูลผู้ใช้
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
		return
	}

	password := []byte(user.Password)
	password1 := userjson.Password

	// ตรวจสอบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword(password, []byte(password1)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid password"})
		return
	}

	// สร้าง JWT (JSON Web Token) ด้วยข้อมูลผู้ใช้และกำหนดเวลาหมดอายุ
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Username,                         // ระบุผู้ออก token ให้เป็นชื่อผู้ใช้
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // กำหนดเวลาหมดอายุให้กับ token เป็น 1 วัน
	})

	// ลงชื่อเซ็นต์กับคีย์เรื่องรหัสลับเพื่อสร้าง token
	token, err := claims.SignedString([]byte("secret"))

	// ตรวจสอบว่ามีข้อผิดพลาดในขั้นตอนการสร้าง token หรือไม่
	if err != nil {
		// ถ้าเกิดข้อผิดพลาด ส่งคำตอบกลับในรูปแบบข้อความข้อผิดพลาด
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to generate JWT token"})
		return
	}

	// ส่ง token กลับไปยังผู้ใช้ในรูปแบบ JSON พร้อมกับสถานะ "success"
	c.JSON(http.StatusOK, gin.H{"status": "success", "jwt": token})
}
