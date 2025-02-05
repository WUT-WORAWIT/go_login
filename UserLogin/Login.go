package UserLogin

import (
	"logins/Models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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
		// ExpiresAt: time.Now().Add(time.Minute).Unix(), // กำหนด ExpiresAt ให้มีค่าเป็นหลังจากปัจจุบัน 1 นาที
	})

	// ลงชื่อเซ็นต์กับคีย์เรื่องรหัสลับเพื่อสร้าง token
	token, err := claims.SignedString([]byte("psmadmin"))

	// ตรวจสอบว่ามีข้อผิดพลาดในขั้นตอนการสร้าง token หรือไม่
	if err != nil {
		// ถ้าเกิดข้อผิดพลาด ส่งคำตอบกลับในรูปแบบข้อความข้อผิดพลาด
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to generate JWT token"})
		return
	}

	// ส่ง token กลับไปยังผู้ใช้ในรูปแบบ JSON พร้อมกับสถานะ "success"
	c.JSON(http.StatusOK, gin.H{"status": "success", "jwt": token})
}
