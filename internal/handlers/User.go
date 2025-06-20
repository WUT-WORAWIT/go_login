package handlers

import (
	"fmt"
	M "go_login/internal/middleware"
	"go_login/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers ... Get all users
func GetRemark(c *gin.Context) {
	str := `
	func: CreateUser
	POST(http://localhost:8080/api/creatuser)
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
	GET(http://localhost:8080/api/getuserall)
	paramiter:

	-----------------------------------------------
	func: GetUserByID
	GET(http://localhost:8080/api/getuserbyid?id=1)
	paramiter:
	Username=

	-----------------------------------------------
	func: UpdateUser
	http://localhost:8080/api/updateuser?id=2
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
	resultChan := make(chan []models.User)
	errChan := make(chan error)

	go func() {
		var users []models.User
		err := models.GetAllUsers(&users) // Pass pointer to users slice
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- users
	}()

	// Wait for result
	select {
	case users := <-resultChan:
		c.JSON(http.StatusOK, users)
	case err := <-errChan:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// GetUsers ... Get all users
func GetBranch(c *gin.Context) {
	branchno := c.Query("branchno")
	branches, err := models.GetBranchAll(branchno)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, branches)
	}
}

// CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user models.User
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
	err = models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type TokenStatus struct {
	Code    int
	Message string
}

// GetUserByUsername ... Get the user by id
func GetUserByID(c *gin.Context) {

	// เรียกใช้งาน JWTMiddleware() และเก็บค่าที่ return ไว้ในตัวแปร middlewareErr
	middleware := M.JWTMiddleware() // เรียกใช้งาน JWTMiddleware() เพื่อรับ middleware
	middleware(c)                   // เรียกใช้งาน middleware โดยส่งค่า context เข้าไป
	if c.IsAborted() {
		// หยุดการดำเนินการเมื่อเกิดข้อผิดพลาด
		return
	}

	// id := c.Params.ByName("id")
	Username := c.Query("Username")
	var user models.User
	err := models.GetUserByID(&user, Username)
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
	var user models.User
	// id := c.Params.ByName("id")
	id := c.Query("id")
	err := models.GetUserByID(&user, id)
	if err != nil {
		errorResponse := gin.H{"status": "error", "message": err.Error()}
		// ส่ง JSON กลับไปยังผู้ใช้พร้อมกับ HTTP status code 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
		// c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(c, &user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	// else {
	// 	c.JSON(http.StatusOK, user)
	// }
}

// DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user models.User
	// id := c.Params.ByName("id")
	id := c.Query("id")

	err := models.GetUserByID(&user, id)
	if err != nil {
		// ส่ง JSON กลับไปยังผู้ใช้พร้อมกับ HTTP status code 500 Internal Server Error
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User with id " + id + " not found"})
		return
		// c.JSON(http.StatusNotFound, user)
	}
	err = models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	// หากข้อมูลถูกลบสำเร็จ
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User with id " + id + " has been deleted"})
}
