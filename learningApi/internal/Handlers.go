package handlers

import (
	"fmt"
	"net/http"

	"github.com/MatasGedziunas/goLearningApi/database"
	"github.com/gin-gonic/gin"
)

func GetAccountBalance(context *gin.Context) {
	account, err := database.GetAccount(context.Query("username"))
	if err != nil {
		fmt.Printf("Error: %v", err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.IndentedJSON(http.StatusOK, account)
	}

}
