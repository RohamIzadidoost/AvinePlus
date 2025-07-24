package main

import (
	"glasscutting/internal/db"
	"glasscutting/internal/handler/api"
	"glasscutting/internal/handler/middleware"
	"glasscutting/internal/repository/gormrepo"
	"glasscutting/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// init DB
	dbConn := db.Init()

	// repositories
	userRepo := gormrepo.NewUserRepository(dbConn)
	orderRepo := gormrepo.NewOrderRepository(dbConn)

	// services
	userSvc := service.NewUserService(userRepo)
	orderSvc := service.NewOrderService(orderRepo)

	// handlers
	userH := api.NewUserHandler(userSvc)
	orderH := api.NewOrderHandler(orderSvc)
	statusH := api.NewStatusHandler(orderSvc)

	r := gin.Default()
	// serve simple web UI
	r.Static("/", "web")

	r.POST("/register", userH.Register)
	r.POST("/login", userH.Login)

	auth := r.Group("/")
	auth.Use(middleware.Auth())

	auth.POST("/orders", orderH.Create)
	auth.GET("/orders/:id", orderH.Get)

	admin := auth.Group("/admin")
	admin.Use(middleware.RequireRole("admin"))
	admin.GET("/orders", orderH.ListAll)
	admin.POST("/orders/:id/verify", statusH.Verify)
	admin.POST("/orders/:id/assign", statusH.Assign)

	employee := auth.Group("/employee")
	employee.Use(middleware.RequireRole("employee"))
	employee.POST("/orders/:id/approve", statusH.EmployeeApprove)

	user := auth.Group("/user")
	user.Use(middleware.RequireRole("customer"))
	user.POST("/orders/:id/approve", statusH.UserApprove)

	r.Run(":8080")
}
