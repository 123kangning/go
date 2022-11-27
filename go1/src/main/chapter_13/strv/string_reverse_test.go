package strv

import (
	"fmt"
	"strings"
	"testing"
)

var test = []struct {
	in  string
	out string
}{
	{"asdf", "fdsa"},
	{"", ""},
	{"1 2", "2 1"},
}

func TestReverse(t *testing.T) {
	for _, te := range test {
		if strings.Compare(Reverse(te.in), te.out) != 0 {
			t.Log(fmt.Sprintf("%s reverse fail", te.in))
			t.Fail()
		}
	}
}
