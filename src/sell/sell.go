package main

import (
	L "go_login/UserLogin"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	func handler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "This is the sell service")
//	}
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Group for endpoints without JWT middleware
	grppost := r.Group("/sell")
	{
		grppost.GET("/", GetRemark)
		grppost.POST("login", L.Login)
	}
	return r
}

func main() {

	r := SetupRouter()
	//running
	r.Run()
}
func GetRemark(c *gin.Context) {
	str := `
	func: สวัสดีครับ1222
	 `
	c.String(http.StatusOK, str)
}
