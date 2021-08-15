package day2_client

//编码解码接口实现后，服务器要确认客户端的编码解码方式与服务器一致。
//

const MagicNumber = 0x3b3fc

type Option struct {
	MagicNumber int
	CodeType    string
}
