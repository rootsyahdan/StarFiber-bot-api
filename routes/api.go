package routes

import (
	"miniproject/constants"
	"miniproject/controllers"
	"miniproject/middleware"

	mid "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	middleware.LogMiddlewares(e)

	e.GET("/admins", controllers.GetAdminsController)
	e.GET("/admins/:id", controllers.GetAdminController)
	e.POST("/admins", controllers.CreateAdminController)
	e.PUT("/admins/:id", controllers.UpdateAdminController)
	e.DELETE("/admins/:id", controllers.DeleteAdminController)
	e.POST("/admins/login", controllers.LoginAdminController)

	e.POST("/new", controllers.CreateTransactionAutomaticallyController)

	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users", controllers.GetUsersController)
	eJwt.GET("/users", controllers.GetUsersController)
	eJwt.GET("/users/:id", controllers.GetUserController)
	eJwt.POST("/users", controllers.CreateUserController)
	eJwt.PUT("/users/:id", controllers.UpdateUserController)
	eJwt.DELETE("/users/:id", controllers.DeleteUserController)

	eJwt.GET("/memberships", controllers.GetMembershipsController)
	eJwt.GET("/memberships/:id", controllers.GetMembershipController)
	eJwt.POST("/memberships", controllers.CreateMembershipController)
	eJwt.PUT("/memberships/:id", controllers.UpdateMembershipController)
	eJwt.DELETE("/memberships/:id", controllers.DeleteMembershipController)

	eJwt.GET("/transactions", controllers.GetTransactionsController)
	eJwt.GET("/transactions/:id", controllers.GetTransactionController)
	eJwt.POST("/transactions", controllers.CreateTransactionController)
	eJwt.PUT("/transactions/:id", controllers.UpdateTransactionController)
	eJwt.DELETE("/transactions/:id", controllers.DeleteTransactionController)

	return e
}
