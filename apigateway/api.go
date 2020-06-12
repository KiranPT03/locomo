package apigateway

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/locomo/apigateway/sender"
)

var router *gin.Engine

func InitApi() {
	fmt.Println("In function InitApi")
	router = gin.Default()
	router.POST("/locomo/api/v1.0/sender/", sender.Create)
	router.Run(":5001")
}
