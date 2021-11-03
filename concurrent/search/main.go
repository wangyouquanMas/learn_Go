package main

import (
	"fmt"
	"time"
)

type ext struct {
	name   string
	age    int
	istrue bool
}

func istrue(ext2 *ext) {
	if ext2.istrue {
		ext2.age = 20
	} else {
		ext2.age = 30
	}
}

func testConcurrency(ext2 *ext, age int, name string) {

	ext2.istrue = name == "wang"

	istrue(ext2)

	fmt.Println("before", ext2.istrue)
	time.Sleep(2 * time.Microsecond)
	fmt.Println("after", ext2.istrue)
	// write decide
	if ext2.istrue {
		fmt.Println("1", name, ext2.age)
	} else {
		fmt.Println("2", name, ext2.age)
	}
}

func main() {

	ext := &ext{}

	for i := 0; i < 10; i++ {
		go testConcurrency(ext, 12, "wang")
		go testConcurrency(ext, 10, "zeng")
		//fmt.Println("2", ext.age, ext.name)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("3", ext.name, ext.age)
}
