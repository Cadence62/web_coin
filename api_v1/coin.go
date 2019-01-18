package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"

	_ "web_coin/config"
	"web_coin/model"
	"web_coin/utils"
)

func SupportCoin(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)
	supportcoin := viper.Get("supportcoin")
	result.Data = supportcoin

}

func NewCoinAddress(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)
	var arg model.NewCoinAddress
	if err := c.ShouldBindJSON(&arg); err != nil {
		result.Code = ArgError
		result.Msg = ErrotMsg(ArgError)
		return
	} else {
		result.Data = arg.CoinName
		result.Code = OK
		result.Msg = ErrotMsg(OK)
	}
}
