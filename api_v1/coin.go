package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"web_coin/coinapi"

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
	var arg model.NewCoinAddressPost
	var data model.NewCoinAddressResult
	if err := c.ShouldBindJSON(&arg); err != nil {
		result.Code = ArgError
		result.Msg = Msg(ArgError)
		return
	} else {
		switch arg.CoinName {
		case "xrp":
			data.CoinName = arg.CoinName
			data.Address = coinapi.NewAddress("rKu9QEsBax7h4fv2qJnfXPVyqL98rMqBvG")
		}
		result.Data = data
		result.Code = OK
		result.Msg = Msg(OK)
	}
}
