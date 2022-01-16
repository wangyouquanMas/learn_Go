package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("20060102"))
	fmt.Println(float64(time.Now().Unix()))
}
