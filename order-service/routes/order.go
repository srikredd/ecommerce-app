package routes

import (
	"order-service/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/orders", func(c *gin.Context) { controllers.CreateOrder(c, db) })
	r.GET("/orders/:id", func(c *gin.Context) { controllers.GetOrder(c, db) })
	r.PUT("/orders/:id", func(c *gin.Context) { controllers.UpdateOrder(c, db) })
	r.DELETE("/orders/:id", func(c *gin.Context) { controllers.DeleteOrder(c, db) })
}
