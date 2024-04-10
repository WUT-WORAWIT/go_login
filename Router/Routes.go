package Routes

import (
	C "logins/Controllers"
	M "logins/JWTmiddleware"
	L "logins/Login"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grppost := r.Group("/user-api")
	{
		grppost.GET("/", C.GetRemark)
		grppost.POST("creatuser", C.CreateUser)
		grppost.POST("login", L.Login)
	}
	grpget := r.Group("/userget")
	grpget.Use(M.JWTMiddleware()) // เรียกใช้ middleware ที่ทำการตรวจสอบ JWT
	{
		grpget.GET("getuserall", C.GetUsersAll)
		grpget.GET("getuserbyid", C.GetUserByID)
		grpget.PUT("updateuser", C.UpdateUser)
		grpget.DELETE("deleteuser", C.DeleteUser)
	}
	return r
}
