package main

import (
	"product-service/models"
	"product-service/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=postgres-product user=postgres password=postgres dbname=productdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Product{})

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8082")
}
