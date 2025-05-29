package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var secretKey = []byte("psmadmin") // ควรย้ายไปอยู่ใน config หรือ env
var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(username string) (string, error) {
	// สร้าง token ใหม่
	token := jwt.New(jwt.SigningMethodHS256)

	// กำหนด claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // หมดอายุใน 24 ชั่วโมง
	// claims["exp"] = time.Now().Add(time.Minute * 5).Unix() // หมดอายุใน 5 นาที

	// ลงนาม token ด้วย secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
