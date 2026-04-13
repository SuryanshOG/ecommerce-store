package main

import (
	"log"
	"product-service/internal/config"
	"product-service/internal/handler"
	"product-service/internal/repository"
	"product-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	db := config.ConnectDB()
	repo := &repository.ProductRepository{DB: db}
	svc := &service.ProductService{Repo: repo}
	h := &handler.ProductHandler{Service: svc}

	r := gin.Default()

	r.POST("/products", h.CreateProduct)
	r.GET("/products", h.GetAllProducts)

	log.Println("Product Service running on port 8082")
	r.Run(":8082")
}
