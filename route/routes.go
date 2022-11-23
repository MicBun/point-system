package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"pointSystem/controller"
	"pointSystem/middleware"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/login", controller.LoginUser)
	r.POST("/register", controller.CreateUser)

	pointMiddlewareRouter := r.Group("/point")
	pointMiddlewareRouter.Use(middleware.JwtAuthMiddleware())
	pointMiddlewareRouter.GET("", controller.GetPoint)
	pointMiddlewareRouter.POST("", controller.UpdatePoint)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
