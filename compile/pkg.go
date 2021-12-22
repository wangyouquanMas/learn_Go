package main

func fk(a, b []int) {
	a = make([]int, 0)
	b = make([]int, 0)
}

func main() {
	var a []int
	var b []int
	fk(a, b)
}
