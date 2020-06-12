package main

import (
	"fmt"

	"github.com/KiranPT03/locomo/apigateway"
)

func main() {
	fmt.Println("Hi from main package")
	fmt.Println(apigateway.Init_api())
}
