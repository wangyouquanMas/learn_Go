package _interface

import "log"

//go的隐式继承

//go method  is a function with a receiver ,a receiver is a value or a pointer of a named or structed type .

type User struct {
	Name string
	Email string
}

func (u User) Notify0() error{

	return nil
}

func (u *User) Notify1() error{

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
// 指针接收，那么SendNotification（）接口实参就必须是指针
// value接收者，那么那么SendNotification（）接口实参就必须是value
// 同时该接收者被认为是 接口的隐式实现
func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

func main()  {
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
}