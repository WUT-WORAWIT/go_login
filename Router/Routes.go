package Routes

import (
	C "logins/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("/", C.GetRemark)
		grp1.POST("creatuser", C.CreateUser)
		grp1.GET("getuserall", C.GetUsersAll)
		grp1.GET("getuserbyid", C.GetUserByID)
		grp1.PUT("updateuser", C.UpdateUser)
		grp1.DELETE("deleteuser", C.DeleteUser)
	}
	return r
}
