package api_v1

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	urls_api_v1 := route.Group("/v1")
	{
		urls_api_v1.GET("/supportcoin", SupportCoin)        //当前支持的币种
		urls_api_v1.POST("/newcoinaddress", NewCoinAddress) //当前支持的币种
	}

}
