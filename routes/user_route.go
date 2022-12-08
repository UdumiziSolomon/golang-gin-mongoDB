package routes

import (
	"github.com/UdumiziSolomon/Gopher-Backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	// user routes

	userRouter := router.Group("/api/user")

	userRouter.POST("/register", controllers.CreateUser())
	userRouter.GET("/account/:id", controllers.GetUserByID())
	userRouter.GET("/accounts", controllers.GetAllUsers())
	userRouter.DELETE("/account/delete/:id", controllers.DeleteUserByID())
	userRouter.PUT("/account/update/:id", controllers.UpdateUserByID())
}