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
		之前每次createParser时，都需要创建一个新的parser,实际上，如果parser可以复用，为了
节省内存和对象创建的时间，我们可以将parser事先创建好缓存起来。

        现在调用craterParser函数时，从缓存中取出parser对象直接使用。
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
		"json": json.NewDecoder(ruleConfigFileExtensiong),
		"gob":  gob.NewDecoder(ruleConfigFileExtensiong),
		"xmo":  xml.NewDecoder(ruleConfigFileExtensiong),
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
