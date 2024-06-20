package main

import (
	handlers "github.com/MatasGedziunas/goLearningApi/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/account/balance", handlers.GetAccountBalance)

	router.Run("localhost:8080")
}
