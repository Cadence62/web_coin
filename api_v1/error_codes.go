package api_v1

const (
	OK             = 2000 //请求成功
	ArgError       = 3001 //请求参数错误
	NotSupportCoin = 4001 //不支持的虚拟币
)

var statusText = map[int]string{
	OK:             "OK",
	ArgError:       "Arg Error or Not Support Coin",
	NotSupportCoin: "Not Support Coin",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func ErrotMsg(code int) string {
	return statusText[code]
}
