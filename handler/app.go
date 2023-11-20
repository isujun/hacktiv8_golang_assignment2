package handler

import (
	"assignment2/docs"
	"assignment2/infrastructure/database"
	"assignment2/repository/order_repository/order_repo"
	"assignment2/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {
	db := database.GetDatabaseInstance()

	// Dependency injection
	orderRepo := order_repo.NewOrderRepo(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := NewOrderHandler(orderService)

	route := gin.Default()

	//Swagger
	docs.SwaggerInfo.Title = "Assignment 2"
	docs.SwaggerInfo.Description = "API with DDD, GIN, and GORM"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Route
	route.GET("/orders", orderHandler.GetOrders)
	route.POST("/orders", orderHandler.CreateOrder)
	route.PATCH("/orders/:orderId", orderHandler.UpdateOrder)
	route.DELETE("orders/:orderId", orderHandler.DeleteOrder)
	route.Run(":8080")
}
