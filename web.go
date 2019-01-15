package main

import "github.com/gin-gonic/gin"
import "web_coin/api_v1"

func main() {
	router := gin.Default()
	api_v1.Routes(router)
	router.Run(":8080")
}
