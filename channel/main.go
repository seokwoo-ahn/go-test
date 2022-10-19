package main

import (
	"go-test/channel/libs"
)

func main() {
	// timestamp로 어긋나는 예시
	libs.WithTimeStamp()
	libs.WithoutTimeStamp()
}
