package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/ecommerce/internal/handler"
	"github.com/thesayedirfan/ecommerce/internal/repository"

	"github.com/thesayedirfan/ecommerce/internal/service"
)

func main() {
	cartRepo := repository.NewCartRepository()
	orderRepo := repository.NewOrderRepository()
	discountRepo := repository.NewDiscountRepository()


	cartService := service.NewCartService(orderRepo,cartRepo,discountRepo)
	discountService := service.NewAdminService(orderRepo, discountRepo)

	cartHandler := handler.NewCartHandler(cartService,discountService)
	adminHandler := handler.NewAdminHandler(discountService)

	router := gin.Default()

	router.POST("/cart/add", cartHandler.AddToCart)
	router.POST("/cart/checkout", cartHandler.Checkout)

	router.POST("/admin/discount/generate", adminHandler.GenerateDiscountCode)
	router.GET("/admin/stats", adminHandler.GetAdminStats)
	
	log.Println("Starting server on :8080")
	router.Run(":8080")
}