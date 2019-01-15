package api_v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllCoin(c *gin.Context) {
	All_Coin := []string{"BTC", "BCH", "XRP", "ETH"}
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": All_Coin,
		})
}

func Routes(route *gin.Engine) {
	urls_api_v1 := route.Group("/v1")
	{
		urls_api_v1.GET("/allcoin", AllCoin)
	}
}
