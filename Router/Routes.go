package Routes

import (
	C "go_login/Controllers"
	M "go_login/Middleware"

	// L "logins/UserLogin"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Group for endpoints without JWT middleware
	grppost := r.Group("/user-api")
	{
		grppost.GET("/", C.GetRemark)
		grppost.GET("getbranch", C.GetBranch)
		grppost.POST("creatuser", C.CreateUser)
		// grppost.POST("login", L.Login)
	}

	// Group for endpoints with JWT middleware
	grpget := r.Group("/user-api")
	grpget.Use(M.JWTMiddleware()) // Applying JWT middleware
	{
		grpget.GET("getuserall", C.GetUsersAll)
		grpget.GET("getuserbyid", C.GetUserByID)
		grpget.PUT("updateuser", C.UpdateUser)
		grpget.DELETE("deleteuser", C.DeleteUser)
	}
	return r
}
