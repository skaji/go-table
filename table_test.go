package table

import (
	"bytes"
	"testing"
)

func TestTable(test *testing.T) {
	t := New()
	t.Add([]string{"a", "b", "cccc"})
	t.Add([]string{"aa", "b", "cc"})
	buf := bytes.NewBuffer([]byte{})
	t.Render(buf)

	var expect string
	expect += "a  b cccc\n"
	expect += "aa b cc\n"
	if buf.String() != expect {
		test.Fatal("fail")
	}
}
