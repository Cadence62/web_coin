package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"web_coin/api_v1"
	_ "web_coin/config"
	_ "web_coin/model"
)

func main() {
	router := gin.Default()
	api_v1.Routes(router)

	_ = router.Run(viper.GetString("system.port"))
}
