package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo() error {
	//Wrap 方法用来包装底层错误，增加上下文文本信息并附加调用栈。 一般用于包装对第三方代码（标准库或第三方库）的调用。
	//return errors.Wrap(sql.ErrNoRows, "foo failed")
	return errors.Wrap(nil, "foo failed")
	//如果不需要增加额外上下文信息，仅附加调用栈后返回，可以使用 WithStack 方法：
	//return errors.WithStack(sql.ErrNoRows)
}

func bar() error {
	//WithMessage 方法仅增加上下文文本信息，不附加调用栈。
	//如果确定错误已被 Wrap 过或不关心调用栈，可以使用此方法。 注意：不要反复 Wrap ，会导致调用栈重复
	return errors.WithMessage(foo(), "bar failed")
}

func main() {
	err := bar()
	//Cause方法用来判断底层错误 。
	if errors.Cause(err) == sql.ErrNoRows {
		// 使用 %v 作为格式化参数，那么错误信息会保持一行， 其中依次包含调用栈的上下文文本
		fmt.Printf("data not found, %v\n", err)
		// 使用 %+v ，则会输出完整的调用栈详情
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}
