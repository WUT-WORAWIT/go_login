package Routes

import (
	"logins/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("/", Controllers.GetRemark)
		grp1.POST("creatuser", Controllers.CreateUser)
		grp1.GET("getuserall", Controllers.GetUsersAll)
		grp1.GET("getuserbyid", Controllers.GetUserByID)
		// grp1.PUT("user/:id", Controllers.UpdateUser)
		// grp1.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}
