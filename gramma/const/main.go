package main

import "fmt"

type QueryMode int64

const (
	SearchMode QueryMode = iota
	SearchAuditMode
	SearchCheckMode
	SearchReportMode
)

func main() {
	fmt.Println(SearchMode, SearchAuditMode, SearchCheckMode, SearchReportMode)
}
