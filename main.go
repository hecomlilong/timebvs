package main

import (
	"fmt"
	"timebvs/libs"
)

func main() {
	obj := libs.Constructor()
	obj.Set("love", "high", 10)
	obj.Set("love", "low", 20)
	fmt.Println(obj.Get("love", 15))
}

