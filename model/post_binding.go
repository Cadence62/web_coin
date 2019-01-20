package model

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	_ "web_coin/config"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("if_support_coin", Binding_coin_support_test)
	}
}

//申请新地址
type NewCoinAddressPost struct {
	LabelName string `json:"label_name" binding:"required"`
	CoinName  string `json:"coin_name" binding:"required,if_support_coin"`
}
type NewCoinAddressResult struct {
	CoinName string `json:"coin_name"`
	Address  string `json:"address"`
}

func Binding_coin_support_test(fl validator.FieldLevel) bool {
	supportcoin := viper.GetStringMap("supportcoin")
	if _, ok := supportcoin[fl.Field().String()]; ok {
		return true
	} else {
		return false
	}
}
