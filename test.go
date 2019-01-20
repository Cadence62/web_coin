package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//time := int32(time.Now().Unix())
	//time := int(time.Now().Unix())
	//print(time,"\n")
	//randint := rand.Intn(100)
	//print(randint,"\n")
	//tag := strconv.Itoa(time + randint)
	//print(tag,"\n")
	//tag = string([]rune(tag)[1:9])
	rand.Seed(time.Now().Unix())
	fmt.Print(rand.Intn(100))
}
