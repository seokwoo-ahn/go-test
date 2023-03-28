package types

import (
	"fmt"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
)

func TestHumanize(t *testing.T) {
	fmt.Println(humanize.Bytes(123451489))
	fmt.Println(humanize.IBytes(123451489))
	fmt.Println(humanize.Time(time.Now()))
	fmt.Println(humanize.Ordinal(1))
	fmt.Println(humanize.Ordinal(2))
	fmt.Println(humanize.Comma(123443252))
	fmt.Println(humanize.Ftoa(3.54))
	fmt.Println(humanize.SI(0.000008240, ""))
	fmt.Println(humanize.SI(0.20240, ""))
	fmt.Println(humanize.SI(0.40, ""))

	if val, err := humanize.ParseBytes("24MB"); err == nil {
		fmt.Println(val)
	}
}
