package main

import (
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"io/ioutil"
	"learn_Go/designPattern/factory/simple_factory_interface/cowbody"
	"os"
)

/*
   现状： 在第一种类型中，如果需要新增一种parser时，必然要改动creatorParser的代码。
      那么是不是影响开闭原则呢？ 实际上，如果不是需要频繁地添加新的parser,只是偶尔修改一下creatorParser代码，稍微不符合开闭
原则也是可以的。

	我们现在采用接口设计，当新增一种parser时，只需要新增一个实现了 runtime.encode接口就可以了，

    所以工厂方法模式比起简单工厂模式更加符合开闭原则
*/

type Factory struct {
}

type RuleConfigSource struct {
}

func (ruleConfigSource RuleConfigSource) load(ruleConfigFilePath string) {
	ruleConfigFileExtensiong, code := ruleConfigSource.getFileExtension(ruleConfigFilePath)
	factory := &Factory{}
	parser := factory.creatorParser(ruleConfigFileExtensiong, code)
	//parser, _ := json.Marshal(ruleConfigFileExtensiong)
	//var parser runtime.Decoder
	//parser = nil
	//if "json" == code {
	//	parser = json.NewDecoder(ruleConfigFileExtensiong)
	//} else if "gob" == code {
	//	parser = gob.NewDecoder(ruleConfigFileExtensiong)
	//} else if "xml" == code {
	//	parser = xml.NewDecoder(ruleConfigFileExtensiong)
	//}
	//else if "toml" == code {
	//	parser = toml.NewDecoder(ruleConfigFileExtensiong)
	//}
	//str := struct {
	//	name int
	//}{}
	str := &Student{}

	parser.Decode(str)
	fmt.Printf("config:%+v", str)
}

func (ruleConfigSource RuleConfigSource) getFileExtension(filePath string) (*os.File, string) {
	//f, _ := ioutil.ReadFile(filePath)
	f, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	//fmt.Printf("%s", string(f))
	//parser, _ := json.Marshal(f)
	return f, "json"
}

func (factory Factory) creatorParser(ruleConfigFileExtensiong *os.File, code string) runtime.Decoder {
	cachedParsers := map[string]runtime.Decoder{
		"json":   json.NewDecoder(ruleConfigFileExtensiong),
		"gob":    gob.NewDecoder(ruleConfigFileExtensiong),
		"xmo":    xml.NewDecoder(ruleConfigFileExtensiong),
		"cowboy": cowbody.NewCowbody(),
	}
	return cachedParsers[code]
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

func main() {
	//
	reRuleConfigSource := &RuleConfigSource{}

	reRuleConfigSource.load("./conclusion.json")

	stu := Student{
		Name:  "HaiCoder",
		Age:   109,
		Score: 100.5,
	}
	//str := struct {
	//	name string `json:"name"`
	//	age  int    `json:"age"`
	//}{
	//	name: "wang",
	//	age:  26,
	//}
	//f, _ := os.OpenFile("./access.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	byte1, _ := json.Marshal(stu)
	//f.Write(byte1)
	//f.Close()
	ioutil.WriteFile("./conclusion.json", byte1, 0666)

	// If the file doesn't exist, create it, or append to the file
	//f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if _, err := f.Write([]byte("appended some data\n")); err != nil {
	//	f.Close() // ignore error; Write error takes precedence
	//	log.Fatal(err)
	//}
	//if err := f.Close(); err != nil {
	//	log.Fatal(err)
	//}
}
