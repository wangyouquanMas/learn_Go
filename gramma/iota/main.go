package main

import "fmt"

type Business int64

var (
	Businesses = []Business{
		BusinessUGCReply,
		BusinessOGVReply,
		BusinessUGCDM,
		BusinessOGVDM,
		BusinessArchive,
		BusinessAccount,
		BusinessArchiveTransport,
	}
)

const (
	BusinessUnspecified      Business = iota
	BusinessUGCReply                  // ugc 评论审核
	BusinessOGVReply                  // ogv 评论审核
	BusinessUGCDM                     // ugc 弹幕审核
	BusinessOGVDM                     // ogv 弹幕审核
	BusinessArchive                   // ugc 稿件审核
	BusinessAccount                   // 账号审核
	BusinessArchiveTransport          // ugc 搬运稿件审核
)

func main() {
	fmt.Println(BusinessUnspecified, BusinessArchiveTransport)
}
