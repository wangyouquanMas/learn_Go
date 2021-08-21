package main

import (
	"fmt"
	"log"
	"sort"
)

//go的隐式继承

//go method  is a function with a receiver ,a receiver is a value or a pointer of a named or structed type .

type User struct {
	Name  string
	Email string
}

func (u User) Notify0() error {

	return nil
}

func (u *User) Notify1() error {

	return nil
}

//Interface
//INotifier: t is a convention in Go to name interfaces with an -er suffix when the interface contains only one method
type Notifier interface {
	Notify() error
}

// notify Notifier ：
func SendNotification(notify Notifier) error {
	return notify.Notify()
}

// 1 指针接收，那么SendNotification（）接口实参就必须是指针
// 2 value接收者，那么那么SendNotification（）接口实参就必须是value
// 同时该接收者被认为是 接口的隐式实现
func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

//3 定义了一个interface的变量，那么这个变量里面可以存实现这个interface的任意类型的对象
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	locan  float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {

	//
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	i = mike //"多态" 【接口可以调用实现它的类的方法】
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//4  interface够持有实现该接口的对象，所以我们可以定义一个包含Men类型元素的slice，这个slice可以被赋予实现了Men接口的任意结构的对象
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, tom

	for _, value := range x {
		value.SayHi()
	}

	// Value of type User can be used to call the method
	// with a value receiver.
	bill := User{"Bill", "bill@email.com"}
	bill.Notify0()

	// Pointer of type User can also be used to call a method
	// with a value receiver.
	jill := &User{"Jill", "jill@email.com"}
	jill.Notify0()

	// Value of type User can be used to call the method
	// with a pointer receiver.
	bill1 := User{"Bill", "bill@email.com"}
	bill1.Notify1()

	// Pointer of type User can also be used to call a method
	// with a value receiver.
	jill1 := &User{"Jill", "jill@email.com"}
	jill1.Notify1()

	//interface
	user := &User{
		Name:  "janet jones",
		Email: "janet@email.com",
	}

	SendNotification(user)

	//5 interface的变量可以持有任意实现该interface类型的对象,可以通过定义interface参数，让函数接受各种类型的参数 [类似Object]。
	/*	如fmt.Print**() 可以接受任意类型的数据,就是有 a ...interface{}作为参数
		空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。【所以空interface可以作为参数】
			func Printf(format string, a ...interface{}) (n int, err error) {
			return Fprintf(os.Stdout, format, a...)
			}*/

	Test(3, 4, 5)
}

//6 如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。
type Interface interface {
	sort.Interface      //嵌入字段sort.Interface
	Push(x interface{}) //a Push method to push elements into the heap
	Pop() interface{}   //a Pop elements that pops elements from the heap
}

func Test(a ...interface{}) {
	fmt.Printf("a is %d", a)
}
