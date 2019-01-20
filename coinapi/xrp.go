package coinapi

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
	"strings"
	"time"
	_ "web_coin/config"
	"web_coin/utils"
)

var XrpServerUrl string = viper.GetString("xrp.serverurl")
var XrpPrecision string = "1e6"

type XrpBasePost struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type AccountInfoPost struct {
	Account     string `json:"account"`
	Strict      bool   `json:"strict"`
	LedgerIndex string `json:"ledger_index"`
	Queue       bool   `json:"queue"`
}

type AccountInfoResult struct {
	Result struct {
		AccountData struct {
			Account           string `json:"Account"`
			Balance           string `json:"Balance"`
			Flags             int    `json:"Flags"`
			LedgerEntryType   string `json:"LedgerEntryType"`
			OwnerCount        int    `json:"OwnerCount"`
			PreviousTxnID     string `json:"PreviousTxnID"`
			PreviousTxnLgrSeq int    `json:"PreviousTxnLgrSeq"`
			Sequence          int    `json:"Sequence"`
			Index             string `json:"index"`
		} `json:"account_data"`
		QueueData struct {
			TxnCount int `json:"txn_count"`
		} `json:"queue_data"`
		Status    string `json:"status"`
		Validated bool   `json:"validated"`
	} `json:"result"`
}

// 生成充值地址
func NewAddress(address string) string {
	rand.Seed(time.Now().Unix())
	tag_int := rand.Intn(999999999-100000000) + 100000000
	tag := strconv.Itoa(tag_int)
	return strings.Join([]string{address, tag}, "_")
}

// 返回地址余额
func GetBalance(account string) string {
	params := new(AccountInfoPost)
	params.Account = account
	params.Strict = true
	params.LedgerIndex = "current"
	params.Queue = true

	data_params := []interface{}{params}
	data := new(XrpBasePost)
	data.Method = "account_info"
	data.Params = data_params
	result_json := utils.Http_test_Post(XrpServerUrl, data)
	var balance AccountInfoResult
	err := json.Unmarshal([]byte(result_json), &balance)
	if err != nil {
		return "error"
	} else {
		balance_data, _ := decimal.NewFromString(balance.Result.AccountData.Balance)
		balance, _ := decimal.NewFromString(XrpPrecision)
		return balance_data.Div(balance).String()
	}

}
