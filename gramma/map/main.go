package main

import "fmt"

//map
const (
	LangIDZHHans = 1
	LangIDZHHant = 2
	LangIDEN     = 3
	LangIDTH     = 4
	LangIDJA     = 5
	LangIDVI     = 6
	LangIDIND    = 7
)

// 语言代号
const (
	LangZHHans = "zh-Hans"
	LangZHHant = "zh-Hant"
	LangEN     = "en"
	LangTH     = "th"
	LangJA     = "ja"
	LangVI     = "vi"
	LangIND    = "id"
)

var LangNames = map[int64]string{
	LangIDZHHans: LangZHHans,
	LangIDZHHant: LangZHHant,
	LangIDEN:     LangEN,
	LangIDTH:     LangTH,
	//LangIDJA:     LangJA,
	LangIDVI:  LangVI,
	LangIDIND: LangIND,
}

var (
	LangIDs   = map[string]int64{}
	RegionIDs = map[string]int64{}
)

func init() {
	// 返回的是 LangNames map中的 key -> id ， value ->name  【不是顺序输出】
	for id, name := range LangNames {
		fmt.Println(name)
		LangIDs[name] = id
	}
}

func main() {

	//1 map声明 这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int
	//numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10

	//2 map是无序的  【类似java中的数据结构】

	//3 长度不固定 【扩容？】

	//4 内置的len函数同样适用于map，返回map拥有的key的数量

	//5 key对应value可覆盖

	//6 not thread-safemap和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制

	// 7 初始化map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	// 8 删除key为C的元素
	delete(rating, "C")

	// 9 make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

	//10 map 遍历获取的是

}
