package model

import (
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	_ "web_coin/config"
)

func init() {
	var validate *validator.Validate
	_ = validate.RegisterValidation("if_support_coin", Binding_coin_support_test)

}

//申请新地址
type NewCoinAddress struct {
	LabelName string `json:"label_name" binding:"required"`
	CoinName  string `json:"coin_name" binding:"required,if_support_coin"`
}

// 检查币种是否支持
//func Binding_coin_support(coin string) bool{
//	supportcoin := viper.GetStringMap("supportcoin")
//	if _,ok :=supportcoin[coin]; ok{
//		return true
//	}else {
//		return false
//	}
//}
func Binding_coin_support_test(fl validator.FieldLevel) bool {
	supportcoin := viper.GetStringMap("supportcoin")
	if _, ok := supportcoin[fl.Field().String()]; ok {
		return true
	} else {
		return false
	}
}
