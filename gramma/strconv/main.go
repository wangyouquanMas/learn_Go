package main

import (
	"fmt"
	"strconv"
)

func main() {
	if ClickRate, err := strconv.ParseFloat("78", 64); err == nil {
		fmt.Println("ScanCtr strconv.ParseFloat(%s) err(%+v)", ClickRate, err)
		err = nil
	}
}
