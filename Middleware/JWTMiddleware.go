package Middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("psmadmin")

// VerifyJWTToken ฟังก์ชันที่ใช้ในการตรวจสอบความถูกต้องของ JWT token
//
//	func VerifyJWTToken(tokenString string, secretKey string) (*jwt.Token, error) {
//		// ตรวจสอบคีย์ลับของ JWT token
//		return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			return []byte(secretKey), nil
//		})
//	}
//
// ฟังก์ชัน VerifyJWTToken สำหรับตรวจสอบและยืนยัน JWT token โดยใช้คีย์ลับที่กำหนด
func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	// ทำการ parse JWT token ด้วยคีย์ลับที่กำหนด
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractClaimsFromJWT ฟังก์ชันที่ใช้ในการดึงข้อมูล claims จาก JWT token
func ExtractClaimsFromJWT(token *jwt.Token) (jwt.MapClaims, bool) {
	// ตรวจสอบ claims ของ JWT token
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

// Jwtmiddleware เป็น middleware ที่ใช้ในการตรวจสอบ JWT token ใน request header
func JWTMiddleware() gin.HandlerFunc {
	// Add Line Numbers to Log Output
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return func(c *gin.Context) {
		// ตรวจสอบ JWT token ที่ส่งมากับ request
		tokenString := c.GetHeader("jwt")
		fmt.Println("tokenString:", tokenString)
		if tokenString == "" {
			// ถ้าไม่มี token ส่งมา ส่งคำตอบกลับให้ผู้ใช้ในรูปแบบข้อความข้อผิดพลาด
			errorResponse := gin.H{"status": "error", "message": "Authorization header is required"}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		// ตรวจสอบคีย์ลับของ JWT token
		// สร้างคีย์ลับที่มีความสุ่มสมบูรณ์
		// secretKey := "psmadmin"
		token, err := VerifyJWTToken(tokenString)
		if err != nil || !token.Valid {
			// ถ้ามีข้อผิดพลาดในการตรวจสอบ token หรือ token ไม่ถูกต้อง ส่งคำตอบกลับให้ผู้ใช้ในรูปแบบข้อความข้อผิดพลาด
			fmt.Println("Error verifying JWT token:", err)
			errorResponse := gin.H{"status": "error", "message": "Invalid or expired token"}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		// ตรวจสอบ claims ของ JWT token
		claims, ok := ExtractClaimsFromJWT(token)
		if !ok || !token.Valid {
			fmt.Println("Error verifying JWT token:", err)
			// ถ้าข้อมูลผู้ใช้ใน token ไม่ถูกต้อง หรือ token ไม่ถูกต้อง ส่งคำตอบกลับให้ผู้ใช้ในรูปแบบข้อความข้อผิดพลาด
			errorResponse := gin.H{"status": "error", "message": "Invalid token claims"}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}
		log.Println(claims)

		// ถ้า token ถูกต้อง ส่งต่อไปยัง handler ถัดไป
		c.Next()
	}
}
