package main

import (
	"miniproject/controllers"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func main() {
	e := echo.New()
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	e.GET("/memberships", controllers.GetMembershipsController)
	e.GET("/memberships/:id", controllers.GetMembershipController)
	e.POST("/memberships", controllers.CreateMembershipController)
	e.PUT("/memberships/:id", controllers.UpdateMembershipController)
	e.DELETE("/memberships/:id", controllers.DeleteMembershipController)

	e.GET("/transactions", controllers.GetTransactionsController)
	e.GET("/transactions/:id", controllers.GetTransactionController)
	e.POST("/transactions", controllers.CreateTransactionController)
	e.PUT("/transactions/:id", controllers.UpdateTransactionController)
	e.DELETE("/transactions/:id", controllers.DeleteTransactionController)

	e.POST("/admins", controllers.CreateAdminController)
	e.POST("/admins/login", controllers.LoginAdminController)

	e.POST("/new", controllers.CreateTransactionAutomaticallyController)
	e.GET("/revanue-monthly", controllers.GetMonthlyRevenueController)

	e.Logger.Fatal(e.Start(":1312"))
}
