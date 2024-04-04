package Controllers

import (
	"fmt"
	"logins/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers ... Get all users
func GetRemark(c *gin.Context) {
	str := `
	func: CreateUser
	POST(http://localhost:8080/user-api/creatuser)
	Body -> JSON
	 {
	 	"ID": 3,
	 	"Name": "test2",
	 	"Age": 19
	 }
	 
	-----------------------------------------------
	func: GetUsers
	GET(http://localhost:8080/user-api/getuserall)
	paramiter:

	-----------------------------------------------
	func: GetUserByID
	GET(http://localhost:8080/user-api/getuserbyid?id=1)
	paramiter:
	id

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
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	// id := c.Params.ByName("id")
	id := c.Query("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
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
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
