package main

import "fmt"

// if 拓展用法
//func (s *Service) getSynonymKeyword(keyword string) (synonymKw []string, hit bool, err error) {
// if s.searchDicts == nil || s.searchDicts.SynonymDicts == nil {
//  return
// }
// if hitItem, hit := s.searchDicts.SynonymDicts[keyword]; hit {
//  if len(hitItem.OutWord) > 0 {
//   synonymKw = hitItem.OutWord
//   log.Info("search rewrite keyword(%s) to synonym keyword(%+v)", keyword, synonymKw)
//   return synonymKw, hit, nil
//  } else {
//   hit = false
//  }
// }


// 接口
type Phone interface {
 call()

// 被包含的接口的所有方法都会被包含到新的接口中。
//只有实现接口中所有的方法，包括被包含的接口的方法，才算是实现了接口。

 Battery
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
 fmt.Println("I am Nokia, I can call you!")
}

func (nokiaPhone NokiaPhone) charge() {
 fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
 fmt.Println("I am iPhone, I can call you!")
}
func (iPhone IPhone) charge() {
 fmt.Println("I am iPhone, I can call you!")
}

type Battery interface{
 charge()
}

type NanFuBattery struct{
}

func (nanFuBattery NanFuBattery) charge(){
  fmt.Println("I'm in charge of battery")
}

func main() {

 //接口使用
 var phone Phone

 // new 返回的是传入类型的指针
 phone = new(NokiaPhone)
 phone.call()


 phone = new(IPhone)
 phone.call()

 var battery Battery

 battery = new(NanFuBattery)
 battery.charge()

 var message = "abc"

 var pointer *string
 // &也可以定义变量
 pointer =&message

 fmt.Println(*pointer)

 //强制类型转换  String 《-》 Byte

  var str1 string = "test"
  var data1 []byte = []byte(str1)
  println(data1,len(data1))


  var data[10] byte
  data[0] = 'T'
  data[1] = 'E'
  var str string = string(data[:])
  fmt.Printf(str)


  fmt.Println("************分隔符************")


 //数组
 var (
  field = []string{"name"}
  )
  fmt.Println(field)

 stringarr := []string{"要给值"}
 var stringarr1 = []string{"1"}
 var a []string =  []string{"1"}
 fmt.Println(stringarr1,stringarr,a)





}
