package main

import (
	"fmt"
	"time"
)

func main(){

	//var a[100]int

	for i :=0;i<100;i++{
		go func(ii int){
			for{
				fmt.Printf("Hello from "+"goroutine %d\n",i)
				//a[ii]++
			}
		}(i) //参数
		time.Sleep(time.Minute)
		//fmt.Println(a)
	}

}
