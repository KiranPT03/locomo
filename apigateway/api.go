package apigateway

import (
	"fmt"
)

func Init_api()  string{
	fmt.Println("Hi from api_gateway package")
	s := "Hello from api_gateway package"
	return s
}
