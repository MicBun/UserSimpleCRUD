package routes

import (
	"UserSimpleCRUD/controllers"
	"UserSimpleCRUD/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/login", controllers.LoginUser)

	sellerMiddlewareRoute := r.Group("/admin")
	sellerMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	sellerMiddlewareRoute.POST("/createUser", controllers.CreateUser)
	sellerMiddlewareRoute.GET("/getAllUser", controllers.GetAllUser)
	sellerMiddlewareRoute.POST("/getUserByID", controllers.GetUserByID)
	sellerMiddlewareRoute.PATCH("/updateUserByID/:id", controllers.UpdateUserByID)
	sellerMiddlewareRoute.DELETE("/deleteUserByID", controllers.DeleteUserByID)

	userMiddlewareRoute := r.Group("/user")
	sellerMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	userMiddlewareRoute.GET("/profile", controllers.GetUserByLogin)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
