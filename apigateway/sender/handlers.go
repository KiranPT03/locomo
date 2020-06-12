package sender

import (
	"fmt"
)

func Create(c *gin.Context) {
	fmt.Println("In fucntion create sender")
	c.JSON(201, gin.H{
		"status":  201,
		"message": "Rule created",
	})
}
