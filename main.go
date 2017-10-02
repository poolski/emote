package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var err error
var db *gorm.DB

func main() {

	db, err = gorm.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Status{})

	// Setup Gin router
	r := gin.Default()

	// Setup routes
	r.GET("/health/", healthCheck)
	r.GET("/statuses/", getStatuses)
	r.POST("/statuses/", addStatus)

	// Run server
	r.Run(":3000")
}

func healthCheck(c *gin.Context) {
	if db_err := db.DB().Ping(); db_err != nil {
		c.JSON(500, gin.H{"status": "Error connecting to DB"})
		fmt.Println("Healthcheck: Error connecting to DB")
	} else {
		c.JSON(200, gin.H{"status": "OK"})
	}
}

func getStatuses(c *gin.Context) {
	var statuses []Status
	if err := db.Find(&statuses).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, statuses)

	}
}

func addStatus(c *gin.Context) {
	var status Status
	c.BindJSON(&status)
	db.Create(&status)
	c.JSON(200, status)
}
