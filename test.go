package main

import (
	"strings"
)

func main() {
	arg := []string{"BTC", "BCH", "XRP"}
	for _, x := range arg {
		if a := strings.Compare("XRP", x); a == 0 { //判断2个字符串是否相等
			print("YES")
			break
		}

	}

}
