package main

import (
	"context"
	"fmt"
)

//1 context应用场景 ：携带关键信息，为全链路提供线索，比如接入elk等系统，需要来一个trace_id，那WithValue就非常适合做这个事。

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", "88888888")
	// 携带session到后面的程序中去
	ctx = context.WithValue(ctx, "session", 1)

	process(ctx)
}

func process(ctx context.Context) {
	session, ok := ctx.Value("session").(int)
	fmt.Println(ok)
	if !ok {
		fmt.Println("something wrong")
		return
	}

	if session != 1 {
		fmt.Println("session 未通过")
		return
	}

	traceID := ctx.Value("trace_id").(string)
	fmt.Println("traceID:", traceID, "-session:", session)
}
