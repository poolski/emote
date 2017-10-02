package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func MailGunTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timestamp := c.Request.PostFormValue("timestamp")
		token := c.Request.PostFormValue("token")
		signature := c.Request.PostFormValue("signature")

		file, _ := os.Open("config.json")
		decoder := json.NewDecoder(file)
		config := Config{}
		err := decoder.Decode(&config)
		if err != nil {
			fmt.Println("error:", err)
		}

		key := config.MailgunKey

		s := []string{timestamp, token}
		authSignature := strings.Join(s, "")

		result := CompareHmac256(authSignature, signature, key)
		fmt.Println(result)

		if result == 0 {
			c.Next()
		} else {
			c.JSON(401, "Signatures don't match")
			c.Abort()
			return
		}
	}
}

func CompareHmac256(message string, signature string, key string) int {
	hash1 := ComputeHmac256(message, key)
	return strings.Compare(hash1, signature)
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
