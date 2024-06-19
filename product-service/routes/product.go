package routes

import (
	"product-service/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    r.POST("/products", func(c *gin.Context) { controllers.CreateProduct(c, db) })
    r.GET("/products/:id", func(c *gin.Context) { controllers.GetProduct(c, db) })
    r.PUT("/products/:id", func(c *gin.Context) { controllers.UpdateProduct(c, db) })
    r.DELETE("/products/:id", func(c *gin.Context) { controllers.DeleteProduct(c, db) })
}
