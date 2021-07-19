package main

import (
	"fmt"

	"time"

)

func main()  {

	//当前时间
	 currentTime := time.Now()
	 fmt.Println(currentTime)


	//Time的add方法 ： 就是在当前时间上添加新的时间
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)

	yesterDays := start.Add(-time.Hour*(24+12)).Format("20060102")

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)

	fmt.Printf("yesterDays = %v\n", yesterDays)

//时间拼接

	date := time.Now()
	sLocale := "167772160-4"
	_scanStart :=1
	//_scanEnd :=99

	startKey := fmt.Sprintf("{%s %s}-%d", date, sLocale, _scanStart)
	//endKey := fmt.Sprintf("{%s %s}-%d", date, sLocale, _scanEnd)
    //fmt.Println(startKey)
	fmt.Printf("+",[]byte(startKey))
	//fmt.Printf( []byte(endKey)))
	//fmt.Println(endKey)


	//
	fmt.Println("______我是间隔符_________")

	startKey1 := "{20210707}-1"
	fmt.Println(time.Now().Hour())
	fmt.Println([]byte(startKey1))
	//
	fmt.Println("______我是间隔符_________")
	//格式化输出时间
	fmt.Println(date)
	// 必须使用这个时间才能返回正确的格式化后的时间，其他的都不行
	fmt.Println(date.Format("20060102150405"))

	fmt.Println("______我是间隔符_________")

	fmt.Println(date.UTC().Hour()+8)


}
