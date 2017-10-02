package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func MailGunTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timestamp := c.Request.FormValue("timestamp")
		token := c.Request.FormValue("token")
		signature := c.Request.FormValue("signature")

		key := os.Getenv("API_KEY")
		s := []string{timestamp, token}
		authSignature := strings.Join(s, "")

		if CheckMAC([]byte(authSignature), []byte(signature), []byte(key)) == true {
			c.Next()
		} else {
			c.JSON(401, "Signatures don't match")
			return
		}
	}
}

func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
