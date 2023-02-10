package channel_test

import "testing"

func TestWithTimeStamp(t *testing.T) {
	// timestamp로 어긋나는 예시
	WithTimeStamp()
}

func TestWithoutTime(t *testing.T) {
	WithoutTimeStamp()
}
