package main

import (
	"order-service/models"
	"order-service/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    dsn := "host=postgres-order user=postgres password=postgres dbname=orderdb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }
    db.AutoMigrate(&models.Order{})

    r := gin.Default()
    routes.SetupRoutes(r, db)
    r.Run(":8083")
}
