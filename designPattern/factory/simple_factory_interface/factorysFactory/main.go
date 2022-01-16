package main

import (
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"io/ioutil"
	"os"
)

/*
	实现目标：
		用封装，实现工厂的工厂 【解决了简单工厂方法1的问题】
	在该场景下，当需要添加新的规则配置解析器时，只需要创建新的parser类和 parser factory类，并且在工厂的工厂中，将
新的 工厂对象添加到 缓存即可， 代码改动很少，基本符合开闭原则。
*/

type Factory interface {
	creatorParser(ruleConfigFileExtensiong *os.File, code string) runtime.Decoder
}

type FactoryJson struct {
}

type FactoryXMl struct {
}

type FactoryGob struct {
}

type FactorySFactory struct {
}

type RuleConfigSource struct {
}

func (factorySFactory FactorySFactory) RuleConfigParserFactory(ruleConfigFileExtensiong string) Factory {
	cachedFactories := map[string]Factory{
		"json": new(FactoryJson),
		"gob":  new(FactoryGob),
		"xml":  new(FactoryXMl),
	}
	return cachedFactories[ruleConfigFileExtensiong]
}

func (ruleConfigSource RuleConfigSource) load(ruleConfigFilePath string) {
	ruleConfigFileExtensiong, code := ruleConfigSource.getFileExtension(ruleConfigFilePath)
	//创建工厂的工厂
	factorySFactory := &FactorySFactory{}
	//获取工厂
	factory := factorySFactory.RuleConfigParserFactory(code)
	//获取工厂的解析器
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

func (factory FactoryJson) creatorParser(ruleConfigFileExtensiong *os.File, code string) runtime.Decoder {
	cachedParsers := map[string]runtime.Decoder{
		"json": json.NewDecoder(ruleConfigFileExtensiong),
	}
	return cachedParsers[code]
}

func (factory FactoryXMl) creatorParser(ruleConfigFileExtensiong *os.File, code string) runtime.Decoder {
	cachedParsers := map[string]runtime.Decoder{
		"xml": xml.NewDecoder(ruleConfigFileExtensiong),
	}
	return cachedParsers[code]
}

func (factory FactoryGob) creatorParser(ruleConfigFileExtensiong *os.File, code string) runtime.Decoder {
	cachedParsers := map[string]runtime.Decoder{
		"gob": gob.NewDecoder(ruleConfigFileExtensiong),
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
